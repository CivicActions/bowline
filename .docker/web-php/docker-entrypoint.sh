#!/bin/bash -e
set -e

[ -e "/var/www/.docker/etc/bashrc" ] && source /var/www/.docker/etc/bashrc

# Create required directories just in case.
mkdir -p /var/www/logs/php-fpm /var/www/files-private
echo "*" > /var/www/logs/.gitignore

# Set the apache user and group to match the host user.
OWNER=$(stat -c '%u' /var/www)
GROUP=$(stat -c '%g' /var/www)
USERNAME=www-data
[ -e "/etc/debian_version" ] || USERNAME=apache
if [ "$OWNER" != "0" ]; then
  usermod -o -u $OWNER $USERNAME
  usermod -s /bin/bash $USERNAME
  groupmod -o -g $GROUP $USERNAME
  usermod -d /var/www $USERNAME
  chown -R --silent $USERNAME:$USERNAME /var/www
fi
echo The apache user and group has been set to the following:
id $USERNAME

chmod ug+rwx /var/www/logs

# Set Docker For Mac xdebug settings since docker0 bridge doesn't exist.
if [ "${HOST_OS}" = "Darwin" ]; then
  echo xdebug.remote_host=docker.for.mac.localhost >> /usr/local/etc/php/conf.d/xdebug.ini
  echo xdebug.remote_connect_back=0 >> /usr/local/etc/php/conf.d/xdebug.ini;
else
  echo xdebug.remote_connect_back=1 >> /usr/local/etc/php/conf.d/xdebug.ini;
fi

exec "$@"
