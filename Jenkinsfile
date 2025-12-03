pipeline{
    agent any
    stages{
        stage('checkout'){
            steps{
                checkout scm
            }
        }

        stage('setup') {
            agent {
                docker {
                    image 'golang:1.20-alpine'
                    args '-v /go/pkg/mod:/go/pkg/mod' // cache go modules
                    reuseNode true
                }
            }
            steps {
                sh '''
                    go mod download
                    go version || true
                '''
            }
        }

        stage('Build') {
            agent {
                docker {
                    image 'golang:1.20-alpine'
                    args '-v /go/pkg/mod:/go/pkg/mod' // cache go modules
                    reuseNode true
                }
            }
            steps {
                sh '''
                    export GOCACHE=$WORKSPACE/.cache
                    mkdir -p $GOCACHE
                    CGO_ENABLED=0 GOOS=linux go build -v -o app .
                    ls -lh build/app || true
                '''
            }
        }
        stage('Test') {
            agent {
                docker {
                    image 'golang:1.20-alpine'
                    args '-v /go/pkg/mod:/go/pkg/mod' // cache go modules
                    reuseNode true
                }
            }
            steps {
                sh '''
                    export GOCACHE=$WORKSPACE/.cache
                    mkdir -p $GOCACHE
                    go test ./... -v
                '''
            }
        }

        stage('Docker Build') {
            agent {
                docker {
                    image 'docker:29.1.1-dind-alpine3.22'
                    args '--privileged -v /var/lib/docker:/var/lib/docker'
                    reuseNode true
                }
            }
            environment {
                DOCKER_HOST = "tcp://localhost:2629" // Docker daemon in DinD
            }
            steps {
                script {
                    sh '''
                        export HOME=/tmp
                        mkdir -p $HOME/.docker
                        docker build -t my-simple-app .
                    '''
                }
            }
        }

    }
}