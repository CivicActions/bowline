#!/usr/bin/env bash

# This script is a functional test of bowline that should work in any POSIX
# compatible shell. This script should also pass shellcheck -x linting.
# It should be invoked from the git root as an argument of the shell you want
# to test (this will override the hashbang above) - e.g.: zsh ./tests/test.sh

# Report some succinct version information for orientation when reading logs.
python -mplatform # Easy cross-platform way to output kernel and distro
echo "Shell $(ps -p $$ --no-headers -o comm=) ${BASH_VERSION}${ZSH_VERSION}${KSH_VERSION}"
docker version --format 'Docker client {{.Client.Version}}'
docker version --format 'Docker server {{.Server.Version}}'
docker-compose version | head -n1

echo Starting test
cd fixtures || exit 1
. ../activate
if [ -z ${BOWLINE_ACTIVATED+x} ]; then
  echo "ERROR: Failed to activate"
  exit 2
fi

# Check functionality of an alias that will not exist elsewhere.
if ! echo bowline | bowlinesum | grep -Fq '749335d2c792bd109c40eba7a7fccd75ede4da532c892cbf892d603665feea3f'; then
  echo "ERROR: Aliases not present"
  exit 3
fi
echo Success
