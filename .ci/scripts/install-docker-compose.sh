#!/usr/bin/env bash

echo "Installing all stable releases of Docker Compose"
TAGS=$(git ls-remote https://github.com/docker/compose | grep refs/tags | grep -oP '[0-9]+\.[0-9][0-9]+\.[0-9]+$')
for COMPOSE_VERSION in $TAGS; do
  echo "Fetching Docker Compose version ${COMPOSE_VERSION}"
  sudo curl -LsS -C - https://github.com/docker/compose/releases/download/${COMPOSE_VERSION}/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose-${COMPOSE_VERSION}
done
sudo chmod a+x /usr/local/bin/docker-compose-*
echo "Symlinking most recent stable version"
sudo ln -s /usr/local/bin/docker-compose-${COMPOSE_VERSION} /usr/local/bin/docker-compose