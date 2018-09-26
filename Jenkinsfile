pipeline {
    agent none
    stages {
        stage('Run Linux tests') {
            parallel {
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
                stage('Test on Ubuntu 18.04') {
                    agent { 
                        label 'ubuntu-1804'
                    }
                    steps {
                        checkout scm
                        sh 'ls -la'
                        sh 'docker-compose version'
                        sh 'docker info'
                        sh 'docker run hello-world'
                    }
                }
                stage('Test on Ubuntu 16.04 with latest Docker') {
                    agent { 
                        label 'ubuntu-1604-latest-docker'
                    }
                    steps {
                        checkout scm
                        sh 'ls -la'
                        sh 'docker-compose version'
                        sh 'docker info'
                        sh 'docker run hello-world'
                    }
                }
                stage('Test on Ubuntu 18.04 with latest Docker') {
                    agent { 
                        label 'ubuntu-1804-latest-docker'
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
    }
}
