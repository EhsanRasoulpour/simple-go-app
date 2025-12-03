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

        stage('Start DinD & build') {
            steps {
                sh '''
                # start dinD if not running
                if [ -z "$(docker ps -q -f name=ci-dind)" ]; then
                    docker run -d --name ci-dind --privileged -e DOCKER_TLS_CERTDIR="" docker:24-dind
                    # wait for docker daemon to be ready
                    for i in $(seq 1 20); do
                    docker -H tcp://127.0.0.1:2375 info && break || sleep 1
                    done
                fi
                # use CLI to build against DinD daemon
                docker -H tcp://127.0.0.1:2375 build -t myuser/my-simple-app:${BUILD_NUMBER} .
                '''
            }
}


    }
}