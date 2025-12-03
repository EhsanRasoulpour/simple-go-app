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
                }
            }
            steps {
                sh 'go version || true'
                sh 'go mod download'
            }
        }
    }
}