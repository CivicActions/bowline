package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/CivicActions/bowline/pkg/exposedcmd"
)

func main() {
	composeFiles := []string{"docker-compose.yml"}
	composeProjectName := os.Getenv("COMPOSE_PROJECT_NAME")
	if composeProjectName == "" {
		fmt.Printf("echo -e 'Error getting composer project name: ensure COMPOSE_PROJECT_NAME is set'")
		os.Exit(1)
	}
	commands, err := exposedcmd.GetComposeExposedCommands(composeFiles, composeProjectName)
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
