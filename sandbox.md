# Drupal 7 Installation with Docker and Fig

### Requirements
- [Git 2.0+](http://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [Docker 1.3+](https://docs.docker.com/installation/)
  - Docker service running.
  - Able to run docker command without sudo. (Need to fix this.)
- [Docker Compose](http://docs.docker.com/compose/)
- [Composer](https://getcomposer.org/) (optional)
- A `workspace` directory in your home directory

### Step 1. Clone D7 into ~/workspace/PROJECT
Get the latest "integration" branch code for PROJECT. You will be able to work with this code base normally, as you would in any non-dockerized sandbox.

``` bash
cd ~/workspace
git clone REPOSITORY_URL PROJECT
cd PROJECT
```
NOTE: The rest of these commands run from the `~/workspace/PROJECT` directory.

### Step 2. Install [githooks](https://git.civicactions.net/civicactions/githooks/tree/master)

Githooks enables pre-commit tests that are performed before a commit is successful (generally code quality tests)

``` bash
rm -rf .git/hooks
git clone git@git.civicactions.net:civicactions/githooks.git .git/hooks
```

### Step 3. Place a recent, sanitized copy of the database in the base directory:

``` bash
rsync -P user@example.com:latest/sanitized-daily_drupal.sql.gz .snapshot.sql.gz
```

### Step 4. Build and run your "PROJECT" container

Docker creates virtual instances of Linux on your local machine, and this next step installs those machines, then sets the ~/workspace/PROJECT directory to work with the Linux environments installed here.

``` bash
./scripts/build import
```

- Your first build can take some time as it downloads virtual Linux environments but subsequent builds will be fast.
- When the build is completed a status report will be displayed that includes a link to the new site. For example: `web address: http://172.17.0.5/` Click the link or copy it to your browser.
- If the web address in the status report is not populated, something has prevented your docker instance from being built successfully. Please see the **Troubleshooting** section below.
- If the address is displayed on the status report, but the website does not come up, you may need to manually include a subnet route to your docker instance. The command for adding the route is also at the end of the status report. It will look something like this: `sudo route -n add 172.0.0.0/8 192.168.59.103`
- If you like, you can stop the container when done:

```bash
docker-compose stop
```

- Restart with:

```bash
./scripts/build
```
**NOTE: Each time you restart the container, the IP address of the site will change. See the bottom of the status message for the new IP address for your site.**

### Step 5. Run tests

``` bash
./scripts/run
```

### Using drush and docker-compose

Drush won't work as configured with your containers. We've created a workaround command for using drush commands inside docker containers called "crush". NOTE: this step will need to be repeated each time you start a new bash session
If you would like to use `drush` and `docker-compose` commands on the newly built containers, source the bash include file:

``` bash
source scripts/include.bash
```

- This will add a bash alias for drush called `crush` (container+drush=crush)
  - Examples: `crush uli`, `crush st`, `crush sqlc`, `crush cc all`

### Updating your build

To update your code and restart your container:

``` bash
git pull
./scripts/build
```

### Daily usage

You will need to restart a few functions each time rebuild your container or restart your computer.

**Start crush**

```bash
source scripts/include.bash
```

**OSX users: Restart docker and start crush**

```bash
boot2docker start
$(boot2docker shellinit)
./scripts/build
source scripts/include.bash
```

### Troubleshooting

#### Issue 1: No web address in status report

If you aren't seeing a web address after you build your container, it could be cause by not having docker correctly install and or configured. If you suspect this is the case, try rebuilding your docker installation.

**OSX USERS:** Try rebuilding your boot2docker installation.

``````bash
boot2docker delete
boot2docker init
$(boot2docker shellinit)
```

Then, rerun the following command:

```bash
./scripts/build import
```

#### Issue 2: Container doesn't appear to be working

If your container isn't working it may be from an outdated docker image. Try removing containers and images and start over:

``` bash
./scripts/build destroy  # Read the prompts carefully.
./scripts/build import   # Start over with a fresh build.
```

#### Issue 3: Website working, but theme not displaying

If your website comes up at the expected address, but it's not displaying a theme, it could be a permissions issue in your ~/workspace/PROJECT directory. To fix:

```bash
./scripts/build
```

## Behat Testing
- Behat yml and feature files are checked in under `tests/behat/`
- _more tbd..._

## Selenium Testing
_tbd..._
