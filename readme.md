# Docker for Drupal development and testing

## Requirements
1. [Docker](https://www.docker.com/)
  - Make sure you can successfully run docker commands without sudo. See [Ubuntu example](https://docs.docker.com/installation/ubuntulinux/#giving-non-root-access).
1. [Fig](http://www.fig.sh/)

See also the [wiki](https://github.com/davenuman/bowline/wiki) for [platform-specific instructions](https://github.com/davenuman/bowline/wiki/Platform-specific-instructions).

## Install Instructions

### New Drupal Project
Clone this repo. Change myproject to the name of your project:
```bash
git clone git@github.com:davenuman/bowline.git myproject
cd myproject
```

Activate bowline, adding the bowline environment to your bash session (The "dot space" is intentional, not a typo):
``` bash
. bin/activate
```
This should add your project name to your bash prompt to indicate that your session has extra features. For example, **~/myproject (myproject) $**.
Enter `bowline` to see a list commands available and the status of the containers. Note that the commands listed (such as drush) override any commands that were previously in your $PATH.
Build the default containers:
```
build
```
This will build the containers and can take a long time.

Install Drupal and login:
```
settings_init
drush si --sites-subdir=default
drush uli  # Get a login url.
```

### Existing Drupal Project
1. Go to your project workspace. Make sure your git working directory is clean with `git status` and you might want to try this in a new branch for testing first with `git checkout -b dockerize`.
Add this repository as a remote:
```
git remote add bowline git@github.com:davenuman/bowline.git
git remote update
```

Check out the bowline code. This will stage the files in your current branch
```
git checkout bowline/master .
```
It is possible though unlikely that this step modified some of your files. Check this with `git status` to see what is staged. Or more specifically, you can `git status -s|grep ^M` to list modified files. Feel free to correct these now if you like but you should be able to continue either way.

Activate bowline, adding the bowline environment to your bash session (The "dot space" is intentional, not a typo):
``` bash
. bin/activate
```
This should add your project name to your bash prompt to indicate that your session has extra features. For example, **~/myproject (myproject) $**.
Enter `bowline` to see a list commands available and the status of the containers. Note that the commands listed (such as drush) override any commands that were previously in your $PATH.
Build the default containers:
```
build
```
This will build the containers and can take a long time.

Add docker setting to Drupal's settings.php file:
```
settings_init
drush uli # Get a login url.
```

## Post-Install: Test and document your development sandbox
1. Review [sandbox.md](sandbox.md ) which is indented to become your instructions for your development team. It will need to be modified to the specifics of your project.
1. Replace the content of this readme.md file with appropriate description of your project.
