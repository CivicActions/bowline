# Exposed Commands API

The concept of exposed commands derives from the use of exposed ports in Docker. Similar to a port being exposed for accessing a container's service, a command is made available to the user shell in order to perform some task. This API makes use of Docker's label system to identify commands. Labels can be added to images in a Dockerfile using the [LABEL](https://docs.docker.com/engine/reference/builder/#label) instruction. Also, labels can be added to containers when they are run and can thus be defined in a docker-compose.yml file. Labels in both images and containers are considered in this API.

Docker label keys used in this API are defined as follows:

## LABEL expose.command.single

Expose a single command in the image. The value of this label is the command exposed to the user interface. When executed, this will execute the Docker image default command.

### Example

```
FROM hello-world
LABEL expose.command.single=hello
```
The above label instruction will expose a command with the name `hello` which will run the docker image (with no arguments) resulting in a hello-world message.

## LABEL expose.command.multiple.*

Label keys starting with `expose.command.multiple.` are inspected and the part of the key after that portion is used as the command name. The value is used as the command run in the container.

### Example
```
~

LABEL expose.command.multiple.cksum="/usr/bin/cksum"
LABEL expose.command.multiple.md5sum="/usr/bin/md5sum"
```
An image with the above labels should make the `cksum` and the `md5sum` available, which should run the commands in the container as defined by the label value.

## LABEL expose.command.multiplecommand

The value of this label should be a command that is run inside the image used to get a key-value list of exposed commands. The command needs to be run independently without any mountpoints or docker-compose related dependencies.

### Example

```
FROM alpine:3.8
COPY exposedcommands /usr/local/bin
LABEL expose.command.multiplecommand="/usr/local/bin/exposedcommands"
```
