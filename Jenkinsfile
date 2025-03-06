pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo "Building the Project..........."
                dir('Myfirstapi') {  // Change 'Myfirstapi' to your actual folder name
                    bat "mvn clean"
                }
            }
        }
        stage('Test') {
            steps {
                echo "Testing the Project..............."
                dir('Myfirstapi') {
                    bat "mvn test"
                }
            }
        }
        stage('Compile') {
            steps {
                echo "Compiling the Project..............."
                dir('Myfirstapi') {
                    bat "mvn compile"
                }
            }
        }
        stage('Deploy') {
            steps {
                echo "Deploying the Project..............."
            }
        }
    }
}
