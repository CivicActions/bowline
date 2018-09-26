#!/usr/bin/env bash

export PROJECT=$(gcloud info --format='value(config.project)')
if [ -z "$PROJECT" ]; then
  echo "Run 'gcloud config set project project-name' with a project that has been configured with Jenkins according to the instructions at https://cloud.google.com/solutions/using-jenkins-for-distributed-builds-on-compute-engine"
  exit 1
fi

echo "This will build all Jenkins test agent GCE images and push them to your current gcloud project: '$PROJECT'."
read -n 1 -s -r -p "Press any key to continue"
echo

packer version &> /dev/null
if [ $? -ne 0 ]; then
  echo "Download and install Packer and retry: https://www.packer.io/intro/getting-started/install.html."
  exit 1
fi

packer build jenkins-agent.json
