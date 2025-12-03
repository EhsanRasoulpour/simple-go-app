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
                    image 'golang:1.20-alpine'
                    args '-v /go/pkg/mod:/go/pkg/mod' // cache go modules
                    reuseNode true
                }
            }
            steps {
                script {
                    sh "docker build -t mySimpleApp ."
                }
            }
        }

    }
}