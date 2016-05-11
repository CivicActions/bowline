# Drupal 7 Installation with Docker and Docker Compose

## Requirements
- [Git 2.0+](http://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [Docker 1.3+](https://docs.docker.com/installation/)
  - Make sure you can successfully run docker commands without sudo. See [Ubuntu example](https://docs.docker.com/installation/ubuntulinux/#giving-non-root-access).
- [Docker Compose](http://docs.docker.com/compose/) (Linux users need to install Docker Compose; OSX users get Docker Composed installed automatically on the virtual machines)
- [Vagrant](https://www.vagrantup.com) (OSX and Windows users only)
- [VirtualBox - A GUI for your virtual machines](https://www.virtualbox.org/wiki/Downloads) (Optional, OSX users only)

## Step 1 (Linux). Getting Started with Docker

Install the following software:

- [Git 2.0+](http://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [Docker 1.8+](https://docs.docker.com/installation/)
  - Make sure you can successfully run docker commands without sudo. See [Ubuntu example](https://docs.docker.com/installation/ubuntulinux/#giving-non-root-access).
- [Docker Compose](http://docs.docker.com/compose/)


## Step 1 (OSX). Getting Started with Docker

### 1.0.  Install Git

[Git 2.0+](http://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

### 1.1. Remove boot2docker from OSX

We've had problems reliably running Docker on OSX. These instructions describe running `boot2docker` on [Blinkreaction's Vagrant VM](https://github.com/blinkreaction/boot2docker-vagrant). Some Linux users may also opt run Docker this way from a Vagrant vm.

**REMOVE** boot2docker or any boot2docker variant as it will conflict with the boot2docker variant to be installed. If you have previously installed boot2docker follow [these instructions to remove](http://therealmarv.com/blog/how-to-fully-uninstall-the-offical-docker-os-x-installation/).

To make sure you removed boot2docker, run `vagrant global-status` and confirm that it does not mention boot2docker in the list. Also, running `boot2docker` in the command line should say that the command is not found.

When you run `vagrant global-status` it is possible to see stale results (machines say they're running but they're not). To prune the invalid entries, run `vagrant global-status --prune`.

### 1.2. Install blinkreaction's boot2docker Vagrant VM

Run the following commands to install docker and other requirements:

``` bash
bash <(curl -s https://raw.githubusercontent.com/blinkreaction/boot2docker-vagrant/master/scripts/presetup-mac.sh)
cd ~/workspace
bash <(curl -s https://raw.githubusercontent.com/blinkreaction/boot2docker-vagrant/master/scripts/setup.sh)
```

## Step 2. Clone this repo

``` bash
git clone PROJECT_GIT_URL
cd PROJECT
```

### Step 3. Place a recent, sanitized copy of the database in the base directory:

``` bash
rsync -P user@example.com:latest/sanitized-daily_drupal.sql.gz .snapshot.sql.gz
```

### Step 4. Build and run your "PROJECT" container

Docker creates virtual instances of Linux on your local machine, and this next step installs those machines, then sets the ~/workspace/PROJECT directory to work with the Linux environments installed here.

``` bash
import
```

- Your first build can take some time as it downloads virtual Linux environments but subsequent builds will be fast.
- When the build is completed a status report will be displayed that includes a link to the new site. For example: `web address: http://172.17.0.5/` Click the link or copy it to your browser.
- If the web address in the status report is not populated, something has prevented your docker instance from being built successfully. Please see the **Troubleshooting** section below.
- If the address is displayed on the status report, but the website does not come up, you may need to manually include a subnet route to your docker instance. The command for adding the route is also at the end of the status report. It will look something like this: `sudo route -n add 172.0.0.0/8 192.168.59.103` (you can also try `sudo route -n add 172.17.0.0/16 192.168.10.10`)
- If you like, you can stop the container when done:

```bash
docker-compose stop
```

- Restart with:

```bash
build
```
**NOTE: Each time you restart the container, the IP address of the site will change. See the bottom of the status message for the new IP address for your site.**

### Step 5. Run tests

``` bash
run
```

### Updating your build

To update your code and restart your container:

``` bash
git pull
build
```

### Daily usage

You will need to restart a few functions each time rebuild your container or restart your computer.

```bash
. bin/activate
build
```

*OSX users:

```bash
vagrant up
. bin/activate
build
```

### Troubleshooting

#### Issue 1: No web address in status report

If you aren't seeing a web address after you build your container, it could be cause by not having docker correctly install and or configured. If you suspect this is the case, try rebuilding your docker installation.

**OSX USERS:** Try rebuilding your boot2docker installation.

```bash
vagrant halt
vagrant up
```

Then, rerun the following command:

```bash
build
```

#### Issue 2: Container doesn't appear to be working

If your container isn't working it may be from an outdated docker image. Try removing containers and images and start over:

``` bash
destroy  # Read the prompts carefully.
build
import   # Start over with a fresh build.
```

## Behat Testing
- Behat yml and feature files are checked in under `tests/behat/`
- _more tbd..._

## Selenium Testing
_tbd..._
