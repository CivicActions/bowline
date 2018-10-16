pipeline {
    agent none
    stages {
        stage('Build and push image') {
            agent any
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
                        def latestImage = docker.build("civicactions/bowline-ci", ".")
                        latestImage.push("${env.GIT_COMMIT}")
                    }
                }
            }
        }
        stage('Run tests') {
            parallel {
                stage('Ubuntu 16.04') {
                    agent { 
                        label 'ubuntu-1604'
                    }
                    environment {
                        BOWLINE_IMAGE_SUFFIX = "-ci:${env.GIT_COMMIT}"
                    }
                    steps {
                        checkout scm
                        sh 'ls -la'
                        sh 'docker-compose version'
                        sh 'docker info'
                        dir 'fixtures'
                        sh '. ../activate && if [ -z ${BOWLINE_ACTIVATED+x} ]; then echo ERROR: Failed to activate; exit 1; fi'
                    }
                }
                stage('Ubuntu 18.04') {
                    agent { 
                        label 'ubuntu-1804'
                    }
                    environment {
                        BOWLINE_IMAGE_SUFFIX = "-ci:${env.GIT_COMMIT}"
                    }
                    steps {
                        checkout scm
                        sh 'ls -la'
                        sh 'docker-compose version'
                        sh 'docker info'
                        dir 'fixtures'
                        sh '. ../activate && if [ -z ${BOWLINE_ACTIVATED+x} ]; then echo ERROR: Failed to activate; exit 1; fi'
                    }
                }
                stage('Ubuntu 16.04 with latest Docker') {
                    agent { 
                        label 'ubuntu-1604-latest-docker'
                    }
                    environment {
                        BOWLINE_IMAGE_SUFFIX = "-ci:${env.GIT_COMMIT}"
                    }
                    steps {
                        checkout scm
                        sh 'ls -la'
                        sh 'docker-compose version'
                        sh 'docker info'
                        dir 'fixtures'
                        sh '. ../activate && if [ -z ${BOWLINE_ACTIVATED+x} ]; then echo ERROR: Failed to activate; exit 1; fi'
                    }
                }
                stage('Ubuntu 18.04 with latest Docker') {
                    agent { 
                        label 'ubuntu-1804-latest-docker'
                    }
                    environment {
                        BOWLINE_IMAGE_SUFFIX = "-ci:${env.GIT_COMMIT}"
                    }
                    steps {
                        checkout scm
                        sh 'ls -la'
                        sh 'docker-compose version'
                        sh 'docker info'
                        dir 'fixtures'
                        sh '. ../activate && if [ -z ${BOWLINE_ACTIVATED+x} ]; then echo ERROR: Failed to activate; exit 1; fi'
                    }
                }
                stage('CentOS 7') {
                    agent { 
                        label 'centos-7'
                    }
                    environment {
                        BOWLINE_IMAGE_SUFFIX = "-ci:${env.GIT_COMMIT}"
                    }
                    steps {
                        checkout scm
                        sh 'ls -la'
                        sh 'docker-compose version'
                        sh 'docker info'
                        dir 'fixtures'
                        sh '. ../activate && if [ -z ${BOWLINE_ACTIVATED+x} ]; then echo ERROR: Failed to activate; exit 1; fi'
                    }
                }
                stage('CentOS 7 with latest Docker') {
                    agent { 
                        label 'centos-7-latest-docker'
                    }
                    environment {
                        BOWLINE_IMAGE_SUFFIX = "-ci:${env.GIT_COMMIT}"
                    }
                    steps {
                        checkout scm
                        sh 'ls -la'
                        sh 'docker-compose version'
                        sh 'docker info'
                        dir 'fixtures'
                        sh '. ../activate && if [ -z ${BOWLINE_ACTIVATED+x} ]; then echo ERROR: Failed to activate; exit 1; fi'
                    }
                }
                stage('Test on OS X 10') {
                    agent { 
                        label 'osx-10-docker'
                    }
                    environment {
                        BOWLINE_IMAGE_SUFFIX = "-ci:${env.GIT_COMMIT}"
                    }
                    steps {
                        checkout scm
                        sh 'ls -la'
                        sh 'docker-compose version'
                        sh 'docker info'
                        dir 'fixtures'
                        sh '. ../activate && if [ -z ${BOWLINE_ACTIVATED+x} ]; then echo ERROR: Failed to activate; exit 1; fi'
                    }
                }
                stage('Windows Server 2016 with Docker') {
                    agent { 
                        label 'windows-server-2016-docker'
                    }
                    environment {
                        PATH = "${env.PATH};${env.ALLUSERSPROFILE}\\chocolatey\\bin"
                        BOWLINE_IMAGE_SUFFIX = "-ci:${env.GIT_COMMIT}"
                    }
                    steps {
                        checkout scm
                        bat '''@"%SystemRoot%\\System32\\WindowsPowerShell\\v1.0\\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))"'''
                        bat 'choco install -y docker-compose'
                        bat 'dir'
                        bat 'docker-compose version'
                        bat 'docker info'
                        dir 'fixtures'
                    }
                }
            }
        }
    }
}
