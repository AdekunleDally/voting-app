package main

import (
	  
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var voteTmpl = template.Must(template.ParseFiles("frontend/static/index.html"))
var redisClient *redis.Client

// var dbConn *pgx.Conn
// var ctx = context.Background()

func init() {

	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	initDB()
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Println("No .env file found")
	// }

	// POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	// POSTGRES_PORT := os.Getenv("POSTGRES_PORT")
	// POSTGRES_USER := os.Getenv("POSTGRES_USER")
	// POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")

	// connStr := "postgres://" + POSTGRES_USER + ":" + POSTGRES_PASSWORD + "@" + POSTGRES_HOST + ":" + POSTGRES_PORT + "/votingdb"

	// dbConn, err = pgx.Connect(ctx, connStr)
	// if err != nil {
	// 	log.Fatal("Unable to connect to database:", err)
	// }
}

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handle form submission
		option := r.FormValue("option")
		log.Printf("Vote received: %s", option)
		redisClient.LPush(ctx, "votes", option)
		// Redirect to the vote page after handling the form submission
		http.Redirect(w, r, "/vote", http.StatusSeeOther)
		return
	}

	// Render the template
	err := voteTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
  
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	SERVER_HOST := os.Getenv("SERVER_HOST")
	SERVER_PORT_ONE := os.Getenv("SERVER_PORT_ONE")
	SERVER_PORT_TWO := os.Getenv("SERVER_PORT_TWO")

	// Run the vote processor as a goroutine
	go ProcessVotes()

	// Set up HTTP handlers
	http.HandleFunc("/vote", VoteHandler)
	http.HandleFunc("/result", ResultHandler)

	// Serve static files
	http.Handle("/frontend/static/", http.StripPrefix("/frontend/static", http.FileServer(http.Dir("frontend/static"))))
	http.Handle("/results/static/", http.StripPrefix("/results/static", http.FileServer(http.Dir("results/static"))))

	// Start the HTTP server for handling votes
	log.Printf("Starting vote server on %s:%s", SERVER_HOST, SERVER_PORT_ONE)
	go func() {
		log.Fatal(http.ListenAndServe(SERVER_HOST+":"+SERVER_PORT_ONE, nil))
	}()

	// Start the HTTP server for displaying results
	log.Printf("Starting result server on %s:%s", SERVER_HOST, SERVER_PORT_TWO)
	log.Fatal(http.ListenAndServe(SERVER_HOST+":"+SERVER_PORT_TWO, nil))
}
