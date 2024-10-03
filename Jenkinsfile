pipeline { 
    agent any

    tools {
        go 'go-1.20' // Ensure go-1.20 is installed in Jenkins global tools
    }

    environment {
        GO111MODULE = 'on'
        DOCKER_IMAGE = "lukmanadeokun31/worker-service:${env.BUILD_NUMBER}"
        KUBECONFIG = credentials('kubeconfig-kind') 
    }

    stages {
        stage('Checkout Code') {
            steps {
               // Corrected syntax for git
            //   git branch: 'main', credentialsId: 'github-credentials', url: 'https://github.com/AdekunleDally/voting-app.git'
                 git branch: 'main', credentialsId: 'my-github-credentials', url: 'git@github.com:AdekunleDally/voting-app.git'

            }
        }

        stage('Test') {
            steps {
                dir('worker-service') {
                    bat 'go test ./...' // Running Go tests in voting-service directory on Windows
                }
            }
        }

        stage('Build the Docker Image') {
            steps {
                dir('worker-service') {
                    script {
                        docker.build(DOCKER_IMAGE)
                    }
                }
            }
        }

        stage('Push the Docker Image') {
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'docker-credentials') {
                    docker.image("${DOCKER_IMAGE}:${env.BUILD_NUMBER}").push()   // Push with build number tag
                    docker.image("${DOCKER_IMAGE}:latest").push()               // Push with 'latest' tag
                    }
                }
            }
        }

        // stage('Deploy the worker-service Kubernetes with Helm') {
        //     steps {
        //         script {
        //             sh """
        //             helm upgrade --install worker ./worker-service/worker-chart \
        //                 --set image.repository=${DOCKER_IMAGE} \
        //                 --namespace worker \
        //                 --kubeconfig $KUBECONFIG
        //             """
        //         }
        //     }
        // }
    }
    // post {
    //     failure {
    //         script {
    //             // Rollback logic for failed deployment
    //             sh """
    //             helm rollback voting-app
    //             """
    //         }
    //     }
    //     always {
    //         cleanWs() // Clean workspace after build
    //     }
    // }
}