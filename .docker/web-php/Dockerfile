#
# PHP 5.4
# - also includes Apache, Composer, Drupal configuration
#

FROM debian:wheezy
MAINTAINER David Numan <david.numan@civicactions.com>

ENV DEBIAN_FRONTEND noninteractive

# Keep upstart from complaining
RUN dpkg-divert --local --rename --add /sbin/initctl
RUN ln -sf /bin/true /sbin/initctl

# Packages.
RUN apt-get update && \
  apt-get -y install git mysql-client apache2 libapache2-mod-php5 wget curl \
    php5-curl php5-gd php5-mysql php5-sqlite php5-pgsql php-apc php-pear php5-xdebug \
    pwgen sudo python-setuptools nano apt-utils tidy vim pv && \
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

# Sudo.
RUN echo %sudo	ALL=NOPASSWD: ALL >> /etc/sudoers

# Scripts.
ADD ./conf/scripts/start.sh /start.sh
ADD ./conf/scripts/foreground.sh /etc/apache2/foreground.sh

# Apache.
RUN rm -rf /var/www/*
RUN chmod 755 /etc/apache2/foreground.sh
RUN chmod 755 /start.sh

# Apache / PHP.
ADD ./conf/apache2/vhost.conf /etc/apache2/sites-available/default
ADD ./conf/php5/apache2.ini /etc/php5/apache2/php.ini
ADD ./conf/php5/cli.ini /etc/php5/cli/php.ini
ADD ./conf/php5/apc.ini /etc/php5/conf.d/apc.ini

# Xdebug settings.
RUN \
  echo xdebug.remote_enable=1 >> /etc/php5/mods-available/xdebug.ini && \
  echo xdebug.remote_connect_back=1 >> /etc/php5/mods-available/xdebug.ini && \
  echo xdebug.remote_autostart=0 >> /etc/php5/mods-available/xdebug.ini && \
  echo xdebug.max_nesting_level=256 >> /etc/php5/mods-available/xdebug.ini && \
  echo xdebug.remote_log=/var/www/logs/xdebug.log >> /etc/php5/mods-available/xdebug.ini

EXPOSE 80

# Set a custom entrypoint.
COPY ./docker-entrypoint.sh /
ENTRYPOINT ["/docker-entrypoint.sh"]
