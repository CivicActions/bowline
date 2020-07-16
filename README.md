# Bowline

Bowline (since Version 2) is a tool that inspects docker images and exposes commands to the end user. This is convenient for a developer sandbox environment where developers need to run various cli tools in their process. Also in a CI environment it allows for simple, easy to read tasks such as `build` and `test`.
This allows us to build all the tooling into a docker image instead of adding various scripts and tools outside of the docker ecosystem to achieve simple tasks.

## Using Bowline in a project

Using Bowline is only meaningful if you have at least one docker image that makes use of the [Exposed Commands API](exposed-commands-api.md). To "install" Bowline, simply add the activate file from this repo to your project. Run the following to activate the exposed commands in your docker images.

```
. activate
```

### Working with Docker Compose

Currently Bowline only supports docker-compose based projects. By default, using the activate script as above will mimic the default of docker-compose and parse the docker-compose.yml file and, if it exists, the docker-compose.override.yml file.

Bowline supports the [COMPOSE_FILE environment variable](https://docs.docker.com/compose/reference/envvars/#compose_file) to allow a custom set of docker-compose files to be parsed. If you already have `COMPOSE_FILE` set in your environment and want a different set of files to be parsed by bowline, you can set the `BOWLINE_COMPOSE_FILE` variable in the same format which will be used instead.

Example usage:

```bash
$ COMPOSE_FILE="docker-compose.yml:docker-compose.test.yml"
# In this case, let's say the project just wants to extend the COMPOSE_FILE list, so it is included, like with $PATH
$ BOWLINE_COMPOSE_FILE="${COMPOSE_FILE}:docker-compose.cli.*.yml"
$ ls docker-compose.cli.*.yml
docker-compose.cli.foo.yml
docker-compose.cli.bar.yml
$ . activate
Bowline activated.
Commands added to shell: "test,foo,bar"
$ alias test
alias test='COMPOSE_FILE="docker-compose.yml:docker-compose.test.yml:docker-compose.cli.foo.yml:docker-compose.cli.bar.yml" docker-compose run --rm test test'
$ alias foo
alias foo='COMPOSE_FILE="docker-compose.yml:docker-compose.test.yml:docker-compose.cli.foo.yml:docker-compose.cli.bar.yml" docker-compose run --rm foo foo'
$ alias bar
alias bar='COMPOSE_FILE="docker-compose.yml:docker-compose.test.yml:docker-compose.cli.foo.yml:docker-compose.cli.bar.yml" docker-compose run --rm bar bar'
```

## Developing Bowline

```
# Ensure you are running go 1.11
go version
# Enable Go modules
export GO111MODULE=on
# Run the code (first time will download dependencies)
cd fixtures
go run ../main.go
```

### Updating Docker client libraries

Getting the Docker client version right is a balancing act - too old and it doesn't support the newest edge Docker servers, too new and it doesn't support the oldest ones. Use https://docs.docker.com/develop/sdk/#api-version-matrix to figure out which is the newest version you can support, then look up the official tag name on https://github.com/docker/engine. On older versons you will need to cross reference the commit from the most recent commit to the component for that tag in https://github.com/docker/docker-ce and come up with a matching commit hash instead.

Docker uses an obscure package management system, and things tend to work best if we stick with libraries that match versions they are using. So we go get the main library and then reference the versions from it's vendors.conf, preferring the engine versions.

```
# Enable Go modules
export GO111MODULE=on

CLI=a30dd1b6f35541a1ae32df67e7bdd38d0535aab9
go get github.com/docker/cli@$CLI
curl -s https://raw.githubusercontent.com/docker/cli/$CLI/vendor.conf | grep -f <(cat go.mod | grep -Po "^\t\K([^ ]*)") | cut -d' ' -f1-2 | sed -e 's/ /@/' -e 's/^/go get /'
# Execute the output of the above command

ENGINE=e9b9e4ace294230c6b8eb010eda564a2541c4564
go get github.com/docker/docker@$ENGINE
curl -s https://raw.githubusercontent.com/docker/docker/$ENGINE/vendor.conf | grep -f <(cat go.mod | grep -Po "^\t\K([^ ]*)") | cut -d' ' -f1-2 | sed -e 's/ /@/' -e 's/^/go get /'
# Execute the output of the above command

go mod tidy
```

Then test and commit.
