package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Chdir("fixtures")
	main()
}
