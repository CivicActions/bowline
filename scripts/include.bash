#!/usr/bin/env bash

GIT_ROOT=$(git rev-parse --show-toplevel)
GIT_BRANCH=$(git rev-parse --abbrev-ref @)
SLUG=${GIT_ROOT##*/}
SLUG=${SLUG//-/}  # Fig doesn't allow the - char.
SLUG=${SLUG//_/}  # Fig doesn't allow the _ char.
FIG="fig"  # Add command options here.

# Quick and dirty drush command: container + drush = crush.
if [ $(pwd) != '/var/www/scripts' ]; then
  container="${SLUG}_web_1"
  IP=$(docker inspect --format='{{.NetworkSettings.IPAddress}}' ${container})
  alias crush="docker exec -it $container /var/www/vendor/drush/drush/drush --root=/var/www/docroot --uri=http://${IP}"
fi
