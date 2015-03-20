FROM node:0.12-wheezy
MAINTAINER David Numan <david.numan@civicactions.com>

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && \
  apt-get install -y procps pv && \
  apt-get autoclean && apt-get autoremove

# https://rvm.io/rvm/install
RUN \
  gpg --keyserver hkp://keys.gnupg.net --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3 && \
  curl -sSL https://get.rvm.io | bash -s stable --ruby && \
  echo 'source /usr/local/rvm/scripts/rvm' >> /etc/bash.bashrc

# Install Sass and Compass.
RUN bash -lc 'gem update --system' && bash -lc 'gem install compass'

WORKDIR /data
ADD devtools.js /data/devtools.js
CMD ["/usr/local/bin/node","devtools.js"]

EXPOSE 80
