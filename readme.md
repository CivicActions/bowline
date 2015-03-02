# Docker for Drupal development and testing

## Requirements
1. [Docker](https://www.docker.com/)
  - Make sure you can successfully run docker commands without sudo. See [Ubuntu example](https://docs.docker.com/installation/ubuntulinux/#giving-non-root-access).
1. [Fig](http://www.fig.sh/)

## Install Instructions

### New Drupal Project
1. Clone this repo. `git clone git@github.com:davenuman/bowline.git myproject` Then `cd myproject`. (Change myproject to the name of your project.)
1. Activate bowline, adding the bowline environment to your bash session:
  - `. bin/activate`  # The "dot space" is intentional, not a typo.
  - This should add your project name to your bash prompt to indicate that your session has extra features. For example, **~/myproject (myproject) $**.
  - Enter `bowline` to see a list commands available and the status of the containers. Note that the commands listed (such as drush) override any commands that were previously in your $PATH.
1. Build containers:
  - `build`
  - This will build the containers and can take a long time.
1. Install Drupal and login:
  - `settings_init`
  - `drush si --sites-subdir=default`
  - `drush uli` Get a login url.

### Existing Drupal Project
1. Go to your project workspace. Make sure your git working directory is clean with `git status` and you might want to try this in a new branch for testing first with `git checkout -b dockerize`.
1. Add this repository as a remote:
  - `git remote add bowline git@github.com:davenuman/bowline.git`
  - `git remote update`
1. Check out the bowline code. This will stage the files in your current branch
  - `git checkout bowline/master .`
  - It is possible though unlikely that this step modified some of your files. Check this with `git status` to see what is staged. Or more specifically, you can `git status -s|grep ^M` to list modified files. Feel free to correct these now if you like but you should be able to continue either way.
1. Activate bowline, adding the bowline environment to your bash session:
  - `. bin/activate`  # The "dot space" is intentional, not a typo.
  - This should add your project name to your bash prompt to indicate that your session has extra features. For example, **~/myproject (myproject) $**.
  - Enter `bowline` to see a list commands available and the status of the containers. Note that the commands listed (such as drush) override any commands that were previously in your $PATH.
1. Add docker setting to Drupal's settings.php file:
  - `settings_init`
  - `drush uli` Get a login url.

## Post-Install: Test and document your development sandbox
1. Review [sandbox.md](sandbox.md ) which is indented to become your instructions for your development team. It will need to be modified to the specifics of your project.
1. Replace the content of this readme.md file with appropriate description of your project.
