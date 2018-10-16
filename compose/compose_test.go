package compose

import (
	"os"
	"testing"
)

func TestV1Format(t *testing.T) {
	os.Chdir("../fixtures")
	config, err := LoadFile(files)
	t.Fail("Image incorrect")
}

func TestV2Format(t *testing.T) {
	os.Chdir("../fixtures")
	config, err := LoadFile(files)
}

func TestV3Format(t *testing.T) {
	os.Chdir("../fixtures")
	config, err := LoadFile(files)
}
