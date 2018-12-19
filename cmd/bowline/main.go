package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/CivicActions/bowline/pkg/bowline"
)

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
	commands, err := bowline.GetComposeExposedCommands(composeFiles)
	if err != nil {
		fmt.Printf("echo -e 'Error generating aliases: %q'", err)
		os.Exit(1)
	}
	for alias, command := range commands {
		fmt.Printf("alias %s='%s'\n", alias, command)
	}
	fmt.Println("export BOWLINE_ACTIVATED=1")
	// Print some info to user.
	keys := make([]string, 0, len(commands))
	for key := range commands {
		keys = append(keys, key)
	}
	fmt.Printf("echo -e 'Bowline activated.\nCommands added to shell: %q'\n", strings.Join(keys, ", "))
}
