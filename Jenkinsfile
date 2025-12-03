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

        stage('Docker Build (use host docker socket)') {
            agent {
                docker {
                    image 'docker:29.1.1-cli' // CLI only
                    args '-v /var/run/docker.sock:/var/run/docker.sock -v $HOME/.docker:/root/.docker'
                    reuseNode true
                }
            }
            steps {
                script {
                    sh '''
                    echo "whoami / id:"
                    whoami
                    id
                    echo "socket perms:"
                    ls -l /var/run/docker.sock || true
                    stat -c "owner=%U group=%G mode=%a" /var/run/docker.sock || true
                    echo "groups for current user:"
                    groups || true

                    docker version
                    docker build --progress=plain -t my-simple-app:latest .
                    '''
                }
            }
        }

    }
}