#!/usr/bin/env bash

# This script installs bowline in new or existing Drupal projects.

read_continue () {
  continue_prompt=${1-'Would you like to continue? [Y/n]'};
  echo $continue_prompt
  read response
  response="$(echo ${response} | tr 'A-Z' 'a-z')"
  if [ "${response}" = "n" ]; then
    echo "Exiting"
    exit 1
  fi
}

if [ ! -e ".git" ]; then
  echo "This directory does not contain a git repository."
  echo "This install script must be run from your project directory."

  read_continue "Would you like to create a new project here in ${PWD}? [Y/n]"

  echo "Creating git repository in ${PWD}"
  git init
fi

# Add bowline remote.
git remote add bowline git@github.com:davenuman/bowline.git
git remote update

# Add the Bowline files, except for the readme and install files.
if [ ! -e "lib/bowline" ]; then
  git checkout bowline/master .
  git add . && git status
  git rm --cached readme.md bowline-install.sh
  git commit -m 'Adding bowline code'
fi

if [ ! -e "docroot" ] && [ -e "html" ]; then
  echo "Bowline is configured for Drupal to be in the docroot directory but it looks like you have an html directory."
  echo "This script will move the html directory to docroot and create a symlink at html to it."
  read_continue
  mv html docroot
  ln -sv docroot html
  ls -ld docroot html
fi

echo -e "Bowline is now installed into your project repository.\n"

read_continue "Would you like to build the docker containers now? [Y/n]"

. bin/activate
build
echo -e "Build complete\n"

read_continue "Would you like to initialize the Drupal settings file? [Y/n]"
settings_init
echo -e "Settings initialized\n"
if [ -e "docroot/core/CHANGELOG.txt" ];then
  # Drupal 8 needs newer drush
  composer require drush/drush:8.*
  echo -e "Drush updated for drupal 8.\n"
fi

echo "If you are importing an existing database, save the database export file as .snapshot.sql.gz in this project root directory, then run the 'import' command. If not, you can continue with installing Drupal."
read_continue "Would you like to perform a fresh Drupal site install? [Y/n]"
drush si --sites-subdir=default --site-name=Bowline-Site-Install

echo -e "\nDone. Getting one time login link...\n"
drush uli

echo -e "\nActivate Bowline with the following (including the .)\n. bin/activate\n"
