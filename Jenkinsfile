pipeline {
    agent none
    stages {
        stage('Build') {
            agent any
            steps {
                checkout scm
                sh 'ls -la'
                sh 'docker-compose version'
                sh 'docker info'
            }
        }
        stage('Test on Ubuntu 16.04') {
            agent { 
                label 'ubuntu-1604'
            }
            steps {
                sh 'docker run hello-world'
            }
        }
    }
}
