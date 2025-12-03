pipeline{
    agent any
    stages{
        stage('checkout'){
            steps{
                checkout scm
            }
        }

        stage('setup') {
            agent docker{
                image 'golang:1.20'
                // args '-v /go/pkg/mod:/go/pkg/mod'
            }
            steps {
                sh 'go version || true'
                sh 'go mod download'
            }
        }
    }
}