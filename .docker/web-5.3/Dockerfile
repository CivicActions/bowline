#
# PHP 5.3
# - also includes Apache, Composer, Drupal configuration
#

FROM debian:squeeze
MAINTAINER David Numan <david.numan@civicactions.com>

ENV DEBIAN_FRONTEND noninteractive

# Install utils
RUN apt-get update && apt-get -y install wget curl

# Configure Dotdeb sources
RUN \
  wget -O - http://www.dotdeb.org/dotdeb.gpg | apt-key add - && \
  echo "deb http://packages.dotdeb.org squeeze all" > /etc/apt/sources.list.d/dotdeb.list && \
  echo "\ndeb-src http://packages.dotdeb.org squeeze all" >> /etc/apt/sources.list.d/dotdeb.list && \
  apt-get update

# Install PHP
RUN apt-get -y install php5

# Keep upstart from complaining
RUN dpkg-divert --local --rename --add /sbin/initctl
RUN ln -sf /bin/true /sbin/initctl

# Packages.
RUN apt-get update && \
  apt-get -y install git mysql-client apache2 libapache2-mod-php5 \
    php5-curl php5-gd php5-mysql php5-sqlite php5-pgsql php-apc php-pear php-codesniffer \
    pwgen sudo python-setuptools nano apt-utils tidy vim && \
  apt-get autoclean && apt-get autoremove

# whitelist = phar
RUN \
  mkdir -p /etc/php5/conf.d/ && \
  echo "suhosin.executor.include.whitelist = phar" >> /etc/php5/conf.d/suhosin.ini

# Composer.
RUN \
  curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin -d detect_unicode=0 && \
  ln -s /usr/local/bin/composer.phar /usr/local/bin/composer

# Enable a2enmod rewrite
RUN a2enmod rewrite php5

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

# Start apache.
RUN chmod 755 /start.sh
