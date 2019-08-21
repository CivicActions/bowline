// LoadFile loads a compose file into KomposeObject
package compose

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	//nolint:goimports
	"gopkg.in/yaml.v2"

	"github.com/docker/cli/cli/compose/loader"
	"github.com/docker/cli/cli/compose/types"
	libcomposeconfig "github.com/docker/libcompose/config"
	"github.com/docker/libcompose/lookup"
	"github.com/docker/libcompose/project"
	"github.com/pkg/errors"
)

// We parse a very limited version of the compose file, based on just what we need from the v3 structs.
type Config struct {
	Version  string
	Services Services
}

// Services is a list of ServiceConfig
type Services []ServiceConfig

// ServiceConfig is the configuration of one service
type ServiceConfig struct {
	Name          string
	ContainerName string
	Image         string
	Labels        Labels
}

// Labels is a mapping type for labels
type Labels map[string]string

func LoadFile(files []string) (Config, error) {
	var config Config

	// Load the json / yaml file in order to get the version value
	var version string

	for _, file := range files {
		composeVersion, err := getVersionFromFile(file)
		if err != nil {
			return config, errors.Wrap(err, "Unable to load yaml/json file for version parsing")
		}

		// Check that the previous file loaded matches.
		if len(files) > 0 && version != "" && version != composeVersion {
			return config, errors.New("All Docker Compose files must be of the same version")
		}
		version = composeVersion
	}

	// Convert based on version
	switch version {
	// Use libcompose for 1 or 2
	// If blank, it's assumed it's 1 or 2
	case "", "1", "1.0", "2", "2.0":
		config, err := parseV1V2(files)
		config.Version = version
		if err != nil {
			return config, err
		}
		return config, nil
		// Use docker/cli for 3
	case "3", "3.0", "3.1", "3.2", "3.3":
		config, err := parseV3(files)
		config.Version = version
		if err != nil {
			return config, err
		}
		return config, nil
	default:
		return config, fmt.Errorf("version %s of Docker Compose is not supported. Please use version 1, 2 or 3", version)
	}
}

func getVersionFromFile(file string) (string, error) {
	type ComposeVersion struct {
		Version string `json:"version"` // This affects YAML as well
	}
	var version ComposeVersion
	loadedFile, err := ReadFile(file)

	if err != nil {
		return "", err
	}

	err = yaml.Unmarshal(loadedFile, &version)
	if err != nil {
		return "", err
	}

	return version.Version, nil
}

// Parse Docker Compose with libcompose (only supports v1 and v2). Eventually we will
// switch to using only libcompose once v3 is supported.
func parseV1V2(files []string) (Config, error) {
	var config Config
	var composeObject *project.Project

	// Gather the appropriate context for parsing
	context := &project.Context{}
	context.ComposeFiles = files

	if context.ResourceLookup == nil {
		context.ResourceLookup = &lookup.FileResourceLookup{}
	}

	if context.EnvironmentLookup == nil {
		cwd, err := os.Getwd()
		if err != nil {
			return config, nil
		}
		context.EnvironmentLookup = &lookup.ComposableEnvLookup{
			Lookups: []libcomposeconfig.EnvironmentLookup{
				&lookup.EnvfileLookup{
					Path: filepath.Join(cwd, ".env"),
				},
				&lookup.OsEnvLookup{},
			},
		}
	}

	// Load the context and let's start parsing
	composeObject = project.NewProject(context, nil, nil)
	err := composeObject.Parse()
	if err != nil {
		return config, errors.Wrap(err, "composeObject.Parse() failed, Failed to load compose file")
	}

	// Map the parsed config to a struct we understand.
	for name, service := range composeObject.ServiceConfigs.All() {
		config.Services = append(config.Services, ServiceConfig{
			Name:          name,
			ContainerName: service.ContainerName,
			Image:         service.Image,
			Labels:        Labels(service.Labels),
		})
	}

	return config, nil
}

// The purpose of this is not to deploy, but to be able to parse
// v3 of Docker Compose into a suitable format. In this case, whatever is returned
// by docker/cli's ServiceConfig
func parseV3(files []string) (Config, error) {

	// In order to get V3 parsing to work, we have to go through some preliminary steps
	// for us to hack up github.com/docker/cli in order to correctly parse.
	var composeConfig *types.Config
	var config Config

	// Gather the working directory
	workingDir, err := getComposeFileDir(files)
	if err != nil {
		return config, err
	}

	// get environment variables
	env, err := buildEnvironment()
	if err != nil {
		return config, errors.Wrap(err, "cannot build environment variables")
	}

	configFiles := []types.ConfigFile{}
	for _, file := range files {
		// Load and then parse the YAML first!
		loadedFile, err := ReadFile(file)
		if err != nil {
			return config, err
		}

		// Parse the Compose File
		parsedComposeFile, err := loader.ParseYAML(loadedFile)
		if err != nil {
			return config, err
		}

		// Config file
		configFiles = append(configFiles, types.ConfigFile{
			Filename: file,
			Config:   parsedComposeFile,
		})

	}

	// Config details
	configDetails := types.ConfigDetails{
		WorkingDir:  workingDir,
		ConfigFiles: configFiles,
		Environment: env,
	}

	// Actual config
	// We load it in order to retrieve the parsed output configuration!
	// This will output a github.com/docker/cli ServiceConfig
	// Which is similar to our version of ServiceConfig
	composeConfig, err = loader.Load(configDetails)
	if err != nil {
		return config, err
	}

	for _, service := range composeConfig.Services {
		config.Services = append(config.Services, ServiceConfig{
			Name:          service.Name,
			ContainerName: service.ContainerName,
			Image:         service.Image,
			Labels:        Labels(service.Labels),
		})
	}

	return config, nil
}

// converts os.Environ() ([]string) to map[string]string
// nolint:lll
// based on https://github.com/docker/cli/blob/5dd30732a23bbf14db1c64d084ae4a375f592cfa/cli/command/stack/deploy_composefile.go#L143
func buildEnvironment() (map[string]string, error) {
	env := os.Environ()
	result := make(map[string]string, len(env))
	for _, s := range env {
		// if value is empty, s is like "K=", not "K".
		if !strings.Contains(s, "=") {
			return result, errors.Errorf("unexpected environment %q", s)
		}
		kv := strings.SplitN(s, "=", 2)
		result[kv[0]] = kv[1]
	}
	return result, nil
}

// getComposeFileDir returns compose file directory
// Assume all the docker-compose files are in the same directory
func getComposeFileDir(inputFiles []string) (string, error) {
	inputFile := inputFiles[0]
	if strings.Index(inputFile, "/") != 0 {
		workDir, err := os.Getwd()
		if err != nil {
			return "", errors.Wrap(err, "Unable to retrieve compose file directory")
		}
		inputFile = filepath.Join(workDir, inputFile)
	}
	return filepath.Dir(inputFile), nil
}

// ReadFile read data from file or stdin
func ReadFile(fileName string) ([]byte, error) {
	// StdinData is used for reading compose config
	if fileName == "-" {
		StdinData, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return StdinData, errors.Wrap(err, "Unable to read stdin")
		}
		return StdinData, nil
	}
	return ioutil.ReadFile(fileName)
}
