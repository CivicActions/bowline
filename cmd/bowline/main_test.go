package main

import (
	"os"
	"testing"
)

func TestParseComposeFilesEnv(t *testing.T) {
	var testSets = []struct {
		name    string
		envVars map[string]string
		want    string
	}{
		{"BOWLINE_COMPOSE_FILE one file",
			map[string]string{"BOWLINE_COMPOSE_FILE": "docker-compose.yml"}, "docker-compose.yml"},
		{"BOWLINE_COMPOSE_FILE two files",
			map[string]string{"BOWLINE_COMPOSE_FILE": "docker-compose.yml:file2"}, "docker-compose.yml:file2"},
		{"COMPOSE_FILE",
			map[string]string{"COMPOSE_FILE": "docker-compose.yml"}, "docker-compose.yml"},
		{"No var set",
			map[string]string{}, "docker-compose.yml"},
		{"Both vars set",
			map[string]string{"COMPOSE_FILE": "wrong value", "BOWLINE_COMPOSE_FILE": "correct"}, "correct"},
	}

	for _, ts := range testSets {
		ts := ts
		t.Run(ts.name, func(t *testing.T) {
			os.Unsetenv("BOWLINE_COMPOSE_FILE")
			os.Unsetenv("COMPOSE_FILE")
			for key, val := range ts.envVars {
				os.Setenv(key, val)
			}
			composeFileString, _ := parseComposeFilesEnv()
			if composeFileString != ts.want {
				t.Errorf("Unexpected results: %s does not match %s", composeFileString, ts.want)
			}
		})
	}

}
