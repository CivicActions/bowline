#!/usr/bin/env bash

echo "Installing all stable releases of Docker Compose"
TAGS=$(git ls-remote https://github.com/docker/compose | grep refs/tags | grep -oP '[0-9]+\.[0-9][0-9]+\.[0-9]+$' | sort -n)
for COMPOSE_VERSION in $TAGS; do
  echo "Fetching Docker Compose version ${COMPOSE_VERSION}"
  sudo curl -LsS -C - https://github.com/docker/compose/releases/download/${COMPOSE_VERSION}/docker-compose-Linux-x86_64 -o /usr/local/bin/docker-compose-${COMPOSE_VERSION}
done
sudo chmod a+x /usr/local/bin/docker-compose-*
echo "Symlinking most recent stable version"
sudo ln -s /usr/local/bin/docker-compose-${COMPOSE_VERSION} /usr/local/bin/docker-compose

echo "Installing all stable releases of Habitus"
TAGS=$(git ls-remote https://github.com/cloud66-oss/habitus | grep refs/tags | grep -oP '[0-9]+\.[0-9]+\.[0-9]+$' | sort -n)
for HABITUS_VERSION in $TAGS; do
  echo "Fetching Habitus version ${HABITUS_VERSION}"
  sudo curl -LsS -C - https://github.com/cloud66-oss/habitus/releases/download/${HABITUS_VERSION}/habitus_linux_amd64 -o /usr/local/bin/habitus-${HABITUS_VERSION}
done
sudo chmod a+x /usr/local/bin/habitus-*
echo "Symlinking most recent stable version"
sudo ln -s /usr/local/bin/habitus-${HABITUS_VERSION} /usr/local/bin/habitus
