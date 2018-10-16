package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/CivicActions/bowline/compose"
	"github.com/docker/docker/client"
)

func getComposeExposedCommands(composeFiles []string) (map[string]string, error) {
	commands := make(map[string]string)

	// Loader parses input from file.
	c, err := compose.LoadFile(composeFiles)
	if err != nil {
		return commands, fmt.Errorf("Could not load compose file: %s", err)
	}
	fmt.Println(c)

	docker, err := client.NewEnvClient()
	if err != nil {
		return commands, fmt.Errorf("Could not initialize Docker client: %s", err)
	}

	for _, s := range c.Services {
		if s.Image != "" {
			fmt.Printf("\nsvcname, image: %s, %s\n", s.Name, s.Image)
			image, _, err := docker.ImageInspectWithRaw(context.Background(), s.Image)
			if err != nil {
				return commands, fmt.Errorf("Could not inspect image %s for service %s: %s", s.Image, s.Name, err)
			}
			fmt.Println(image.Config.Labels)
		} else {
			imgName := "bowline_inspect_" + s.Name
			fmt.Printf("\nsvcname, image: %s, %s\n", s.Name, s.Image)
			image, _, err := docker.ImageInspectWithRaw(context.Background(), imgName)
			if err != nil {
				return commands, fmt.Errorf("Could not inspect image %s for service %s: %s", s.Image, s.Name, err)
			}
			fmt.Println(image.Config.Labels)
			for label, value := range image.Config.Labels {
				if label == "expose.command.multiplecommand" {
					// TODO: Use docker.ContainerExecCreate (I think) to execute the command in the image.
					fmt.Printf("docker run --rm %s %s", imgName, value)
				}
				if strings.HasPrefix(label, "expose.command.multiple.") {
					fmt.Printf("alias %s %s\n", label)
				}
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

func main() {
	composeFiles := []string{"docker-compose.yml"}
	err := initCompose(composeFiles)
	if err != nil {
		fmt.Printf("echo -e 'Error running docker-compose initialization: %q'", err)
		os.Exit(1)
	}
	_, err = getComposeExposedCommands(composeFiles)
	if err != nil {
		fmt.Println("echo -e 'Error generating aliases: %q'", err)
		os.Exit(1)
	}
	//fmt.Println(aliases)
}
