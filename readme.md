# Docker for Drupal development and testing

## Install Instructions

1. Go to your project workspace. Make sure your git working directory is clean with `git status` and you might want to try this in a new branch for testing first with `git checkout -b dockerize`.
1. Add this repository as a remote:
  - `git remote add xfigr git@git.civicactions.net:david.numan/transfigure.git`
  - `git remote update`
1. Check out the transfigure code. This will stage the files in your current branch
  - `git checkout xfigr/master .`
  - It is possible though unlikely that this step modified some of your files. Check this with `git status` to see what is staged. Or more specifically, you can `git status -s|grep ^M` to list modified files. Feel free to correct these now if you like but you should be able to continue either way.
1. Review sandbox.md which is indented to become your instructions for your development team. It will need to be modified to the specifics of your project.

## Setting up your development sandbox

See the [Sandbox readme file](sandbox.md)
