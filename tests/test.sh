#!/usr/bin/env bash

# This script is a functional test of bowline that should work in any POSIX
# compatible shell. This script should also pass shellcheck -x linting.
# It should be invoked from the git root as an argument of the shell you want
# to test (this will override the hashbang above) - e.g.: zsh ./tests/test.sh

# Report some succinct version information for orientation when reading logs.
uname -a
echo "Shell $(ps -p $$ | tail -n1 | awk '{print $NF}') ${BASH_VERSION}${ZSH_VERSION}${KSH_VERSION}"
docker version --format 'Docker server {{.Server.Version}}'
docker version --format 'Docker client {{.Client.Version}}'
docker-compose version | head -n1

echo "Starting direct run test"
if ! ./activate | grep -Fq 'Instead, source this file'; then
  echo "ERROR: Activate did not report error when run directly"
  exit 1
fi

echo "Starting main test"
cd fixtures || exit 2
docker-compose pull
docker-compose build --pull
. ../activate
if [ -z ${BOWLINE_ACTIVATED+x} ]; then
  echo "ERROR: Failed to activate"
  exit 3
fi

# Check functionality of an alias that will not exist elsewhere.
if ! echo bowline | bowlinesum | grep -Fq '749335d2c792bd109c40eba7a7fccd75ede4da532c892cbf892d603665feea3f'; then
  echo "ERROR: Aliases not present"
  exit 4
fi
echo Success
