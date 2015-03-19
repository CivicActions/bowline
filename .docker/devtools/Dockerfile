FROM ruby:2.2
MAINTAINER David Numan <david.numan@civicactions.com>

ENV DEBIAN_FRONTEND noninteractive

# Install Sass and Compass.
RUN gem update --system && gem install compass

# Define working directory.
WORKDIR /data
