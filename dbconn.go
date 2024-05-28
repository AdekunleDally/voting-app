// dbconn.go
package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var ctx = context.Background()
var dbConn *pgx.Conn

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")

	connStr := "postgres://" + POSTGRES_USER + ":" + POSTGRES_PASSWORD + "@" + POSTGRES_HOST + ":" + POSTGRES_PORT + "/votingdb"

	dbConn, err = pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
}

// // package main
// // import (
// // 	"context"
// // 	"github.com/jackc/pgx/v4"
// // )
// // var dbConn *pgx.Conn
// // var ctx = context.Background()

// // func init() {
// // 	redisClient = redis.NewClient(&redis.Options{
// // 		Addr: "localhost:6379",
// // 	})

// // 	err := godotenv.Load()
// // 	if err != nil {
// // 		log.Println("No .env file found")
// // 	}

// // 	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
// // 	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")
// // 	POSTGRES_USER := os.Getenv("POSTGRES_USER")
// // 	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")

// // 	connStr := "postgres://" + POSTGRES_USER + ":" + POSTGRES_PASSWORD + "@" + POSTGRES_HOST + ":" + POSTGRES_PORT + "/votingdb"

// // 	dbConn, err = pgx.Connect(ctx, connStr)
// // 	if err != nil {
// // 		log.Fatal("Unable to connect to database:", err)
// // 	}
// // }
