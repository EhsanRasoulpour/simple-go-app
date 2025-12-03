pipeline{
    agent any
    stages{
        stage('checkout'){
            steps{
                checkout scm
            }
        }

        stage('Setup') {
            steps {
            sh 'go version || true'
            sh 'go mod download'
        }
    }
    }
}