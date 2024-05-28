# A Voting  App Service

This Golang application demonstrates a sample voting system, providing users with an interface to vote between two options: cats and dogs.

Upon selection, votes are stored in Redis, functioning as an in-memory database. A worker application, also written in Golang, processes these votes and updates a persistent PostgreSQL database. The PostgreSQL database maintains a table with vote counts for each category.

Furthermore, the voting results are presented in another Golang web application. This application retrieves the results from the PostgreSQL database and displays them to the users."

## Architecture
The application consists of the following components:

1. Voting Frontend Application (Golang): Provides an interface for users to vote.
2. Redis Service: An in-memory database to temporarily store votes.
3. Worker Service (Golang): Processes votes from Redis and updates PostgreSQL.
4. PostgreSQL Service: A persistent database to store the vote counts.
5. Results Frontend Application (Golang): Displays the voting results.

## Prerequisites
1. Golang
2. Docker
3. Redis
4. PostgreSQL

## Setup

This project requires an install of Go, Redis and PostgresSQL



## Running the application

To generate the  code, run the following command:
 

go run main.go worker.go result.go
The voting app inteface runs on localhost:8080 while the results interface runs on localhost:8081/result

