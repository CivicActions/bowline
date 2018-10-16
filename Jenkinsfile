pipeline {
    agent none
    stages {
        stage('Build and push image') {
            when { branch 'master' }
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'civicactionsjenkins') {
                        def latestImage = docker.build("civicactions-internal-it/home", "--build-arg GATSBY_JAZZ_URL=${GATSBY_JAZZ_URL} .")
                        latestImage.push("latest")
                        latestImage.push("${env.GIT_COMMIT}-${env.BUILD_NUMBER}")
                        slackSend channel: 'marketing-home', message: "Master branch built and image pushed successfully to Docker registry"
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
                    steps {
                        checkout scm
                        sh 'ls -la'
                        sh 'docker-compose version'
                        sh 'docker info'
                        sh 'docker run hello-world'
                        sh 'docker build -t civicactions/bowline .'
                        sh '. ./activate && if [ -z ${BOWLINE_ACTIVATED+x} ]; then echo ERROR: Failed to activate; exit 1; fi'
                    }
                }
                stage('Ubuntu 18.04') {
                    agent { 
                        label 'ubuntu-1804'
                    }
                    steps {
                        checkout scm
                        sh 'ls -la'
                        sh 'docker-compose version'
                        sh 'docker info'
                        sh 'docker run hello-world'
                        sh 'docker build -t civicactions/bowline .'
                        sh '. ./activate && if [ -z ${BOWLINE_ACTIVATED+x} ]; then echo ERROR: Failed to activate; exit 1; fi'
                    }
                }
                stage('Ubuntu 16.04 with latest Docker') {
                    agent { 
                        label 'ubuntu-1604-latest-docker'
                    }
                    steps {
                        checkout scm
                        sh 'ls -la'
                        sh 'docker-compose version'
                        sh 'docker info'
                        sh 'docker run hello-world'
                        sh 'docker build -t civicactions/bowline .'
                        sh '. ./activate && if [ -z ${BOWLINE_ACTIVATED+x} ]; then echo ERROR: Failed to activate; exit 1; fi'
                    }
                }
                stage('Ubuntu 18.04 with latest Docker') {
                    agent { 
                        label 'ubuntu-1804-latest-docker'
                    }
                    steps {
                        checkout scm
                        sh 'ls -la'
                        sh 'docker-compose version'
                        sh 'docker info'
                        sh 'docker run hello-world'
                        sh 'docker build -t civicactions/bowline .'
                        sh '. ./activate && if [ -z ${BOWLINE_ACTIVATED+x} ]; then echo ERROR: Failed to activate; exit 1; fi'
                    }
                }
                stage('CentOS 7') {
                    agent { 
                        label 'centos-7'
                    }
                    steps {
                        checkout scm
                        sh 'ls -la'
                        sh 'docker-compose version'
                        sh 'docker info'
                        sh 'docker run hello-world'
                        sh 'docker build -t civicactions/bowline .'
                        sh '. ./activate && if [ -z ${BOWLINE_ACTIVATED+x} ]; then echo ERROR: Failed to activate; exit 1; fi'
                    }
                }
                stage('CentOS 7 with latest Docker') {
                    agent { 
                        label 'centos-7-latest-docker'
                    }
                    steps {
                        checkout scm
                        sh 'ls -la'
                        sh 'docker-compose version'
                        sh 'docker info'
                        sh 'docker run hello-world'
                        sh 'docker build -t civicactions/bowline .'
                        sh '. ./activate && if [ -z ${BOWLINE_ACTIVATED+x} ]; then echo ERROR: Failed to activate; exit 1; fi'
                    }
                }
                stage('Test on OS X 10') {
                    agent { 
                        label 'osx-10-docker'
                    }
                    steps {
                        checkout scm
                        sh 'ls -la'
                        sh 'docker-compose version'
                        sh 'docker info'
                        sh 'docker run hello-world'
                    }
                }
                stage('Windows Server 2016 with Docker') {
                    agent { 
                        label 'windows-server-2016-docker'
                    }
                    environment {
                        PATH = "${env.PATH};${env.ALLUSERSPROFILE}\\chocolatey\\bin"
                    }
                    steps {
                        checkout scm
                        echo "PATH : ${env.PATH}"
                        bat '''@"%SystemRoot%\\System32\\WindowsPowerShell\\v1.0\\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))"'''
                        bat 'choco install -y docker-compose'
                        bat 'dir'
                        bat 'docker info'
                        bat 'docker run hello-world'
                        bat 'docker-compose version'
                    }
                }
            }
        }
    }
}
