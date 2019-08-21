package compose

import (
	"testing"
)

func AssertConfig(services Services, t *testing.T) {
	for _, s := range services {
		//fmt.Printf("\n%s labels: %q\n", s.Name, s.Labels)
		switch s.Name {
		case "test":
			if s.Image != "alpine:latest" {
				t.Error(s.Image)
			}
		case "singleimg":
			if s.Labels["exposed.command.single"] != "testcommand" {
				t.Error(s.Labels)
			}
		case "multiple":
			if s.Labels["exposed.command.multiple.test"] != "testcommand" {
				t.Error(s.Labels)
			}
		}
	}
}

func TestV1Format(t *testing.T) {
	files := []string{"../../fixtures/docker-compose.v1.yml"}
	config, err := LoadFile(files)
	if err != nil {
		t.Errorf("Could not parse compose files: %q, %s", files, err)
	}
	AssertConfig(config.Services, t)
}

func TestV2Format(t *testing.T) {
	files := []string{"../../fixtures/docker-compose.v2.yml"}
	config, err := LoadFile(files)
	if err != nil {
		t.Errorf("Could not parse compose files: %q, %s", files, err)
	}
	AssertConfig(config.Services, t)
}

func TestV3Format(t *testing.T) {
	files := []string{"../../fixtures/docker-compose.yml"}
	config, err := LoadFile(files)
	if err != nil {
		t.Errorf("Could not parse compose files: %q, %s", files, err)
	}
	AssertConfig(config.Services, t)

	// Check with multiple files
	files = append(files, "../../fixtures/docker-compose.v3-override.yml")
	config, err = LoadFile(files)
	if err != nil {
		t.Errorf("Could not parse compose files: %q, %s", files, err)
	}
	AssertConfig(config.Services, t)
}
