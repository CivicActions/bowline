#!/bin/bash -e

export PATH=$HOME/bin:$PATH

# Create required directories just in case.
mkdir -p /var/www/logs /var/www/files-private
echo "*" > /var/www/logs/.gitignore
echo "*" > /var/www/files-private/.gitignore

# Set the apache user and group to match the host user.
OWNER=$(stat -c '%u' /var/www)
GROUP=$(stat -c '%g' /var/www)
usermod -o -u $OWNER www-data
groupmod -o -g $GROUP www-data
chown -R --silent www-data:www-data /var/www
echo Using the following user for running apache:
id www-data

/etc/apache2/foreground.sh
