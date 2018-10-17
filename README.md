# Bowline

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
