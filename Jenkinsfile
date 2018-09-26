pipeline {
    agent none
    stages {
        stage('Test on Ubuntu 16.04') {
            agent { 
                label 'ubuntu-1604'
            }
            steps {
                checkout scm
                sh 'ls -la'
                sh 'docker-compose version'
                sh 'docker info'
                sh 'docker run hello-world'
            }
        }
    }
}
