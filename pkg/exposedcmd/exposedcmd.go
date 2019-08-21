package exposedcmd

import (
	"bufio"
	"context"
	"fmt"
	"path"
	"strings"

	"github.com/CivicActions/bowline/pkg/compose"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	shellquote "github.com/kballard/go-shellquote"
)

func getContainerOutput(ctx context.Context, docker *client.Client, image string, command string) ([]string, error) {
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

	err = docker.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{})
	if err != nil {
		return lines, err
	}

	if err := scanner.Err(); err != nil {
		return lines, err
	}
	return lines, nil
}

func GetComposeExposedCommands(composeFiles []string, composeProjectName string) (map[string]string, error) {
	commands := make(map[string]string)

	// Loader parses input from file.
	c, err := compose.LoadFile(composeFiles)
	if err != nil {
		return commands, fmt.Errorf("could not load compose file: %s", err)
	}

	ctx := context.Background()
	docker, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return commands, fmt.Errorf("could not initialize Docker client: %s", err)
	}

	for _, s := range c.Services {
		var imgName string
		if s.Image != "" {
			_, _, err := docker.ImageInspectWithRaw(context.Background(), s.Image)
			if err != nil {
				return commands, fmt.Errorf("could not inspect image %s for service %s: %s", s.Image, s.Name, err)
			}
			imgName = s.Image
		} else {
			imgName = composeProjectName + "_" + s.Name
		}
		image, _, err := docker.ImageInspectWithRaw(context.Background(), imgName)
		if err != nil {
			return commands, fmt.Errorf("could not inspect image %s for service %s: %s", imgName, s.Name, err)
		}
		// TODO: Merge in compose and image labels here.
		mergedLabels := mergeLabelMaps(image.Config.Labels, s.Labels)
		for label, value := range mergedLabels {
			if strings.HasPrefix(label, "exposed.command.multiplecommand") {
				label = strings.TrimPrefix(strings.TrimPrefix(label, "exposed.command.multiplecommand"), ".")
				lines, err := getContainerOutput(ctx, docker, imgName, value)
				if err != nil {
					return commands, fmt.Errorf("could not run multiplecommand %s (%s) on image %s: %s", label, value, imgName, err)
				}
				for _, line := range lines {
					cmdParts := strings.SplitN(line, " ", 1)
					_, cmd := path.Split(cmdParts[0])
					commands[cmd] = fmt.Sprintf("docker-compose run --rm %s %s", s.Name, line)
				}
			}
			if strings.HasPrefix(label, "exposed.command.multiple.") {
				label = strings.TrimPrefix(label, "exposed.command.multiple.")
				commands[label] = fmt.Sprintf("docker-compose run --rm %s %s", s.Name, value)
			}
			if strings.HasPrefix(label, "exposed.command.single") {
				commands[value] = fmt.Sprintf("docker-compose run --rm %s", s.Name)
			}
		}
	}
	return commands, nil
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
