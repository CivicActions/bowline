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
                        sh 'bash -O expand_aliases ./tests/test.sh'
                        sh 'bash -O expand_aliases --posix ./tests/test.sh'
                        sh 'dash ./tests/test.sh'
                        sh 'zsh ./tests/test.sh'
                        sh 'mksh ./tests/test.sh'
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
                        sh 'bash -O expand_aliases ./tests/test.sh'
                        sh 'bash -O expand_aliases --posix ./tests/test.sh'
                        sh 'dash ./tests/test.sh'
                        sh 'zsh ./tests/test.sh'
                        sh 'mksh ./tests/test.sh'
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
                        sh 'bash -O expand_aliases ./tests/test.sh'
                        sh 'bash -O expand_aliases --posix ./tests/test.sh'
                        sh 'dash ./tests/test.sh'
                        sh 'zsh ./tests/test.sh'
                        sh 'mksh ./tests/test.sh'
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
                        sh 'bash -O expand_aliases ./tests/test.sh'
                        sh 'bash -O expand_aliases --posix ./tests/test.sh'
                        sh 'dash ./tests/test.sh'
                        sh 'zsh ./tests/test.sh'
                        sh 'mksh ./tests/test.sh'
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
                        sh 'bash -O expand_aliases ./tests/test.sh'
                        sh 'bash -O expand_aliases --posix ./tests/test.sh'
                        sh 'dash ./tests/test.sh'
                        sh 'zsh ./tests/test.sh'
                        sh 'mksh ./tests/test.sh'
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
                        sh 'bash -O expand_aliases ./tests/test.sh'
                        sh 'bash -O expand_aliases --posix ./tests/test.sh'
                        sh 'dash ./tests/test.sh'
                        sh 'zsh ./tests/test.sh'
                        sh 'mksh ./tests/test.sh'
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
                        sh 'bash -O expand_aliases ./tests/test.sh'
                        sh 'bash -O expand_aliases --posix ./tests/test.sh'
                        // TODO: Add more test.sh once dash/zsh/mksh are installed
                    }
                }
                stage('Windows Server 2016 with Docker') {
                    agent { 
                        label 'windows-server-2016-docker'
                    }
                    environment {
                        PATH = "${env.PATH};${env.ALLUSERSPROFILE}\\chocolatey\\bin"
                        BOWLINE_IMAGE_SUFFIX = "-ci:${env.GIT_COMMIT}"
                        COMPOSE_CONVERT_WINDOWS_PATHS = "1"
                    }
                    steps {
                        // TODO: Switch to test.sh
                        checkout scm
                        bat '''@"%SystemRoot%\\System32\\WindowsPowerShell\\v1.0\\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))"'''
                        bat 'choco install -y docker-compose cygwin'
                        bat 'dir'
                        bat 'docker-compose version'
                        bat 'docker info'
                        //dir('fixtures') {
                        // Windows shell notes
                        //  With COMPOSE_CONVERT_WINDOWS_PATHS above, docker command works
                        //  $PWD does not exist (cd gets correct path) and alias command is different
                        //
                        // Git bash notes TERM="cygwin" plus EXEPATH=c:\Program Files\Git\bin 
                        //    bat '''"%PROGRAMFILES%\\Git\\bin\\bash.exe"'''
                        //  May need COMPOSE_CONVERT_WINDOWS_PATHS prefix for docker command
                        //  $PWD is incorrect - uses /c/path - can use `cmd //c cd` to get right path
                        //
                        // Cygwin notes TERM="cygwin"
                        //  Takes you out of existing directory (need to re-cd to bowline directory)
                        //  for pwd you need to do cygpath -w "$(pwd)"
                        // 
                        // Babun notes
                        // Wants to use it's own shell - too hard to test automatically.
                        //}
                    }
                }
            }
        }
    }
}
