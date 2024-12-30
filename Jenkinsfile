pipeline {
    agent any
    environment {
        GIT_CREDENTIALS_ID = 'github-ssh-key' // Replace with your Jenkins credential ID
    }
    stages {
        stage('Clone Repository') {
            steps {
                // Clone the repository
                git branch: 'main', url: 'git@github.com:<username>/<repo>.git', credentialsId: env.GIT_CREDENTIALS_ID
            }
        }
        stage('Build Docker Image') {
            steps {
                // Build the Docker image for the Go server
                sh 'docker build -t go-server .'
            }
        }
        stage('Run Server') {
            steps {
                // Run the server in a Docker container
                sh 'docker run -d -p 9090:9090 --name go-server-container go-server'
            }
        }
        stage('Test Server') {
            steps {
                // Test the server using curl
                sh 'curl http://localhost:9090'
            }
        }
        stage('Push Changes to GitHub') {
            steps {
                script {
                    // Make changes and push to GitHub
                    sh '''
                        echo "Automated change by Jenkins $(date)" > jenkins_output.txt
                        git config user.name "Jenkins Bot"
                        git config user.email "jenkins@yourdomain.com"
                        git add .
                        git commit -m "Automated commit by Jenkins"
                        git push origin main
                    '''
                }
            }
        }
        stage('Cleanup') {
            steps {
                // Stop and remove the container
                sh 'docker stop go-server-container && docker rm go-server-container'
            }
        }
    }
}

