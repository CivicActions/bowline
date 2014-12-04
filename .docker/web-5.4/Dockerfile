FROM debian:wheezy
MAINTAINER David Numan david.numan_at_civicactions.com

# Keep upstart from complaining
RUN dpkg-divert --local --rename --add /sbin/initctl
RUN ln -sf /bin/true /sbin/initctl

# Packages.
RUN DEBIAN_FRONTEND=noninteractive apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get -y install sqlite3 mysql-client postgresql-client apache2 curl git libapache2-mod-php5 openssh-server php5-curl php5-gd php5-mysql php5-sqlite php5-pgsql php-apc php-pear php-codesniffer pwgen sudo python-setuptools nano apt-utils tidy vim
RUN DEBIAN_FRONTEND=noninteractive apt-get autoclean
RUN DEBIAN_FRONTEND=noninteractive apt-get autoremove

# whitelist = phar
RUN mkdir -p /etc/php5/conf.d/
RUN echo "suhosin.executor.include.whitelist = phar" >> /etc/php5/conf.d/suhosin.ini

# Composer.
RUN curl -sS https://getcomposer.org/installer | sudo php -- --install-dir=/usr/local/bin -d detect_unicode=0
RUN sudo ln -s /usr/local/bin/composer.phar /usr/local/bin/composer

# Enable a2enmod rewrite
RUN a2enmod rewrite php5

# Sudo.
RUN echo %sudo	ALL=NOPASSWD: ALL >> /etc/sudoers
RUN mkdir /var/run/sshd

# Scripts.
ADD ./conf/scripts/start.sh /start.sh
ADD ./conf/scripts/foreground.sh /etc/apache2/foreground.sh

# Apache.
RUN rm -rf /var/www/*
RUN chmod 755 /etc/apache2/foreground.sh

# Apache / PHP.
ADD ./conf/apache2/vhost.conf /etc/apache2/sites-available/default
ADD ./conf/php5/apache2.ini /etc/php5/apache2/php.ini
ADD ./conf/php5/cli.ini /etc/php5/cli/php.ini
ADD ./conf/php5/apc.ini /etc/php5/conf.d/apc.ini

# Drupal.
ADD ./conf/drupal/settings.docker.php /etc/settings.docker.php

EXPOSE 80

# Run test script.
RUN chmod 755 /start.sh
