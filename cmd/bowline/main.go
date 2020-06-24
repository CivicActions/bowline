package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/CivicActions/bowline/pkg/exposedcmd"
)

func parseComposeFilesEnv() (string, []string) {
	envBCF := os.Getenv("BOWLINE_COMPOSE_FILE")
	envCF := os.Getenv("COMPOSE_FILE")

	// Default if no env vars
	var composeFileString string

	switch {
	case envBCF != "":
		// BOWLINE_COMPOSE_FILE is the preferred env var
		composeFileString = envBCF
	case envCF != "":
		// Use COMPOSE_FILE as second choice
		composeFileString = envCF
	default:
		// Mimic the docker-compose default
		if _, err := os.Stat("docker-compose.override.yml"); os.IsNotExist(err) {
			// Override file doesn't exist, only use docker-compose.yml
			composeFileString = "docker-compose.yml"
		} else {
			// Override file exists, use both
			composeFileString = "docker-compose.yml:docker-compose.override.yml"
		}
	}

	composeFiles := filepath.SplitList(composeFileString)
	return composeFileString, composeFiles
}

func main() {
	composeFileString, composeFiles := parseComposeFilesEnv()
	fmt.Println("# composeFiles: ", composeFiles)

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
		fmt.Printf("alias %s='COMPOSE_FILE=\"%s\" %s'\n", alias, composeFileString, command)
	}
	fmt.Println("export BOWLINE_ACTIVATED=1")
	// Print some info to user.
	keys := make([]string, 0, len(commands))
	for key := range commands {
		keys = append(keys, key)
	}
	fmt.Printf("echo -e 'Bowline activated.\nCommands added to shell: %q'\n", strings.Join(keys, ", "))
}
