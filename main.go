package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/CivicActions/bowline/compose"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	shellquote "github.com/kballard/go-shellquote"
)

func getContainerOutput(docker *client.Client, ctx context.Context, image string, command string) ([]string, error) {
	var lines []string

	// Split the string and parse values according to shell quoting rules.
	cmd, err := shellquote.Split(command)
	if err != nil {
		return lines, err
	}

	resp, err := docker.ContainerCreate(ctx, &container.Config{
		Image: image,
		Cmd:   cmd,
		Tty:   true,
	}, nil, nil, "")
	if err != nil {
		return lines, err
	}

	if err = docker.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return lines, err
	}

	statusCh, errCh := docker.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err = <-errCh:
		if err != nil {
			return lines, err
		}
	case <-statusCh:
	}

	out, err := docker.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return lines, err
	}

	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return lines, err
	}
	return lines, nil
}

func getComposeExposedCommands(composeFiles []string) (map[string]string, error) {
	commands := make(map[string]string)

	// Loader parses input from file.
	c, err := compose.LoadFile(composeFiles)
	if err != nil {
		return commands, fmt.Errorf("Could not load compose file: %s", err)
	}

	ctx := context.Background()
	docker, err := client.NewEnvClient()
	if err != nil {
		return commands, fmt.Errorf("Could not initialize Docker client: %s", err)
	}

	for _, s := range c.Services {
		var imgName string
		if s.Image != "" {
			_, _, err := docker.ImageInspectWithRaw(context.Background(), s.Image)
			if err != nil {
				return commands, fmt.Errorf("Could not inspect image %s for service %s: %s", s.Image, s.Name, err)
			}
			imgName = s.Image
		} else {
			imgName = "bowline_inspect_" + s.Name
		}
		image, _, err := docker.ImageInspectWithRaw(context.Background(), imgName)
		if err != nil {
			return commands, fmt.Errorf("Could not inspect image %s for service %s: %s", imgName, s.Name, err)
		}
		// TODO: Merge in compose and image labels here.
		mergedLabels := mergeLabelMaps(image.Config.Labels, s.Labels)
		for label, value := range mergedLabels {
			if strings.HasPrefix(label, "expose.command.multiplecommand") {
				label = strings.TrimPrefix(strings.TrimPrefix(label, "expose.command.multiplecommand"), ".")
				lines, err := getContainerOutput(docker, ctx, imgName, value)
				if err != nil {
					return commands, fmt.Errorf("Could not run multiplecommand %s (%s) on image %s: %s", label, value, imgName, err)
				}
				for _, line := range lines {
					cmdParts := strings.SplitN(line, " ", 1)
					_, cmd := path.Split(cmdParts[0])
					commands[cmd] = fmt.Sprintf("docker-compose run --rm %s %s", s.Name, line)
				}
			}
			if strings.HasPrefix(label, "expose.command.multiple.") {
				label = strings.TrimPrefix(label, "expose.command.multiple.")
				commands[label] = fmt.Sprintf("docker-compose run --rm %s %s", s.Name, value)
			}
			if strings.HasPrefix(label, "expose.command.single") {
				commands[value] = fmt.Sprintf("docker-compose run --rm %s", s.Name)
			}
		}
	}
	return commands, nil
}

func initCompose(composeFiles []string) error {
	var args []string
	for _, f := range composeFiles {
		args = append(args, "-f", f)
	}
	pull := append(args, "pull")
	cmd := exec.Command("docker-compose", pull...)
	var out bytes.Buffer
	cmd.Stderr = &out
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Docker-compose pull failed:\n  %s", strings.Replace(out.String(), "\n", "\n  ", -1))
	}

	out.Reset()
	build := append(args, "--project-name=bowline_inspect", "build", "--pull")
	cmd = exec.Command("docker-compose", build...)
	cmd.Stderr = &out
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Docker-compose build failed:\n  %s", strings.Replace(out.String(), "\n", "\n  ", -1))
	}
	return nil
}

func mergeLabelMaps(maps ...map[string]string) map[string]string {
	result := make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func main() {
	composeFiles := []string{"docker-compose.yml"}
	err := initCompose(composeFiles)
	if err != nil {
		fmt.Printf("echo -e 'Error running docker-compose initialization: %q'", err)
		os.Exit(1)
	}
	commands, err := getComposeExposedCommands(composeFiles)
	if err != nil {
		fmt.Printf("echo -e 'Error generating aliases: %q'", err)
		os.Exit(1)
	}
	for alias, command := range commands {
		fmt.Printf("alias %s='%s'\n", alias, command)
	}
	fmt.Println("export BOWLINE_ACTIVATED=1")
}
