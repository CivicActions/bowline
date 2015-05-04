# Docker for Drupal development and testing

## Requirements
1. [Docker](https://www.docker.com/)
  - Make sure you can successfully run docker commands without sudo. See [Ubuntu example](https://docs.docker.com/installation/ubuntulinux/#giving-non-root-access).
1. [Docker Compose](http://docs.docker.com/compose/)

See also the [wiki](https://github.com/davenuman/bowline/wiki) for [platform-specific instructions](https://github.com/davenuman/bowline/wiki/Platform-specific-instructions).

## Pre-install Preparations (Optional)
If you prefer to do the heavy downloading ahead of time, run the following docker commands:
```bash
docker pull davenuman/bowline-web-php
docker pull mysql:5.5
```

Also pull these if you plan to use the proxy for nice urls (recommended):
```bash
docker pull nginx
docker pull devries/dnsmasq
```

## Install Instructions

### New Drupal Project
Start a new git repo then add bowline as a remote. Change myproject to the name of your project:
```bash
mkdir myproject
cd myproject
git init
git remote add bowline git@github.com:davenuman/bowline.git
git remote update
```

Check out the bowline code. This will stage the files in your current branch
```
git checkout bowline/master .
```

Add and commit your new code. It is recommended to review the files staged before committing. For example:
```bash
git add . && git status
git rm --cached readme.md	# You probably don't want this readme in your project.
git commit -m 'Starting with bowline code'
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
git rm --cached readme.md	# You probably don't want this readme in your project.
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
```

Get a snapshot of your current site's database and rename it to `.snapshot.sql.gz` then import it:
```
import
```

When that is complete you can proceed with appropriate drush commands such as this:

```
drush st
drush uli # Get a login url.
```

And finally if it all looks good, commit the code to your repo:
```
git commit -m 'Adding bowline code'
```

## Docker subnet proxy for domain names as well as IP addresses
You may have noticed when running the `bowline` command that there is a section called "Proxy". As it suggests, run the following command to activate the proxy:
``` bash
invoke_proxy
```
This will create and start an nginx container linked to a dnsmasq container. After they start it will add the web IP address and the project name to the dns container. The nginx proxy then uses that dns for finding your site and servers it as {projectname}.localtest.me (for example https://myproject.localtest.me/). Once the proxy is active, bowline is aware of it when using the `bowine` or `drush` commands:
``` bash
bowline
drush st
```

**Important:** the nginx proxy will not start if you have something else using port 80. You must either stop your other service (recommended) or edit the port in `lib/proxy/fig.yml` file to something other than 80.

## Post-Install: Test and document your development sandbox
1. Review [sandbox.md](sandbox.md ) which is indented to become your instructions for your development team. It will need to be modified to the specifics of your project.
1. Replace the content of this readme.md file with appropriate description of your project.

## Updating Bowline
The intention for typical usage of Bowling is to set up this repository as a secondary git remote. If you followed these instructions, the remote is called bowline. (You can add whatever is appropriate for your project as the origin remote.)

If you would like to pull in the latest Bowline code you basically just need to update the bin and lib directories.This can be done with the following command:
```bash
git checkout bowline/master -- bin lib
```

Then review the change and commit the update and push to your project.

## Hoist Riggings

Bowline aims to be a general purpose tool while also allowing for any project specific needs. The [hoist](bin/hoist) command helps to allow for many features yet keep simplicity for less involved projects. The word "hoist" is playing on the sailing term, suggesting that you would use the appropriate rigging depending on the direction you are heading.

### Using Hoist
After activating, running the `hoist` command will output something similar to this:
```bash
Usage: hoist [rigging]
Available riggings:
behat  devtools  drupal-core-dev
```

So to "hoist" the behat rigging for your project you would do this:
```bash
hoist behat
```

Thus, you can easily install the pieces you need for your project. Each hoisting is logged in logs/rigging

### Creating a Rigging
Riggings are simply a directory with a rigging_name.hoist bash script. A good example is the [behat rigging](lib/rigging/behat). The [behat.hoist](lib/rigging/behat/behat.hoist) script runs a composer command and copies a file to the bin directory, which may be typical processes of other possible riggings.


## Contributing

Pull requests welcome in typical GitHub fashion. If you would like to add a new feature to the project, consider using the hoist method described above.
