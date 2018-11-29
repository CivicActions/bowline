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
                        sh 'bash --login -i ./tests/test.sh'
                        sh 'bash --login --posix -i ./tests/test.sh'
                        sh 'dash ./tests/test.sh'
                        sh 'zsh --login --interactive ./tests/test.sh'
                        sh 'mksh -li ./tests/test.sh'
                    }
                }
                stage('Ubuntu 18.04') {
                    agent { 
                        label 'ubuntu-1804'
                    }
                    environment {
                        BOWLINE_IMAGE_SUFFIX = "-ci:${env.GIT_COMMIT}"
                        // Force older API version, since the Docker server here is too old.
                        DOCKER_API_VERSION = '1.35'
                    }
                    steps {
                        checkout scm
                        sh 'bash --login -i ./tests/test.sh'
                        sh 'bash --login --posix -i ./tests/test.sh'
                        sh 'dash ./tests/test.sh'
                        sh 'zsh --login --interactive ./tests/test.sh'
                        sh 'mksh -li ./tests/test.sh'
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
                        sh 'bash --login -i ./tests/test.sh'
                        sh 'bash --login --posix -i ./tests/test.sh'
                        sh 'dash ./tests/test.sh'
                        sh 'zsh --login --interactive ./tests/test.sh'
                        sh 'mksh -li ./tests/test.sh'
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
                        sh 'bash --login -i ./tests/test.sh'
                        sh 'bash --login --posix -i ./tests/test.sh'
                        sh 'dash ./tests/test.sh'
                        sh 'zsh --login --interactive ./tests/test.sh'
                        sh 'mksh -li ./tests/test.sh'
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
                        sh 'bash --login -i ./tests/test.sh'
                        sh 'bash --login --posix -i ./tests/test.sh'
                        sh 'zsh --login --interactive ./tests/test.sh'
                        sh 'mksh -li ./tests/test.sh'
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
                        sh 'bash --login -i ./tests/test.sh'
                        sh 'bash --login --posix -i ./tests/test.sh'
                        sh 'zsh --login --interactive ./tests/test.sh'
                        sh 'mksh -li ./tests/test.sh'
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
                        // OS X appears to have an old bash and needs expand_aliases to test.
                        sh 'bash -O expand_aliases ./tests/test.sh'
                        sh 'bash --posix -O expand_aliases ./tests/test.sh'
                        // TODO: Add more test.sh once dash/zsh/mksh are installed
                    }
                }
                stage('Windows 10 with Docker') {
                    agent { 
                        label 'windows-10-docker'
                    }
                    environment {
                        PATH = "${env.PATH};${env.PROGRAMFILES}\\Docker\\Docker"
                        BOWLINE_IMAGE_SUFFIX = "-ci:${env.GIT_COMMIT}"
                    }
                    steps {
                        checkout scm
                        bat '"%PROGRAMFILES%\\Git\\bin\\bash.exe" -O expand_aliases ./tests/test.sh'
                        bat 'set PATH=C:\\tools\\cygwin\\bin;%PATH% && c:\\tools\\cygwin\\bin\\bash.exe -O expand_aliases ./tests/test.sh'
                    }
                }
            }
        }
    }
}
