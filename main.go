package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/docker/docker/client"
	"github.com/kubernetes/kompose/pkg/kobject"
	"github.com/kubernetes/kompose/pkg/loader"
)

func getComposeExposedCommands(composeFiles []string) (map[string]string, error) {
	commands := make(map[string]string)

	// Loader parses input from file into komposeObject.
	l, err := loader.GetLoader("compose")
	if err != nil {
		return commands, fmt.Errorf("Could not load compose parser: %s", err)
	}

	c := kobject.KomposeObject{
		ServiceConfigs: make(map[string]kobject.ServiceConfig),
	}
	c, err = l.LoadFile(composeFiles)
	if err != nil {
		return commands, fmt.Errorf("Could not load compose file: %s", err)
	}
	fmt.Println(c)

	docker, err := client.NewEnvClient()
	if err != nil {
		return commands, fmt.Errorf("Could not initialize Docker client: %s", err)
	}

	for svcName, s := range c.ServiceConfigs {
		if s.Image != "" {
			fmt.Printf("%s, %s\n", svcName, s.Image)
			image, _, err := docker.ImageInspectWithRaw(context.Background(), s.Image)
			if err != nil {
				return commands, fmt.Errorf("Could not inspect image %s for service %s: %s", s.Image, svcName, err)
			}
			fmt.Println(image.Config.Labels)
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
	build := append(args, "build", "--pull")
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
