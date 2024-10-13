pipeline { 
    agent any

    tools {
        go 'go-1.20'
    }

    environment {
        GO111MODULE = 'on'
        DOCKER_IMAGE = "lukmanadeokun31/postgres:latest"
        KUBECONFIG = credentials('kubeconfig-kind') 
    }

    stages {
        stage('Checkout postgres Code') {
            steps {
               // Clone the repository without specifying a branch or path
               git branch: 'postgres', credentialsId: 'my-github-credentials', url: 'git@github.com:AdekunleDally/voting-app.git'

            }
        }

        stage('Build the Postgres Image') {
            steps {
                dir('postgres'){
                    script {
                         bat 'docker build -t "postgres" .'                  
                    }
                }
            }
        }

        stage('Push postgres Image') {
            steps {
                script {
                    withDockerRegistry([credentialsId: 'docker-credentials', url: 'https://registry.hub.docker.com']) {
                        bat 'docker tag "postgres" "lukmanadeokun31/postgres:latest"'
                        bat 'docker push "lukmanadeokun31/postgres:latest"'
                    }
                }
            }
        }

        stage('Load image to KIND Cluster') {
            steps {
                bat 'kind load docker-image lukmanadeokun31/postgres:latest --name votingapp-microservice'
            }
        }

        stage('Deploy with Helm') {
            steps {
                bat "helm upgrade --install postgres ./postgres/postgres-chart -f ./postgres/postgres-chart/values.yaml --kubeconfig=${KUBECONFIG} --set image.repository=${DOCKER_IMAGE} --set image.tag=\"latest\""           
            }
        }

        stage('Test Deployment') {
            steps {
                bat 'kubectl get pods -n postgres-namespace'
            }
        }
    }
}
