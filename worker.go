package main

import (
	
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)
  

func ProcessVotes() {  
	for {
		vote, err := redisClient.LPop(ctx, "votes").Result()
		if err != nil {
			if err == redis.Nil {
				time.Sleep(time.Second)
				continue
			}
			log.Fatal(err)
		}

		if vote == "cat" {
			dbConn.Exec(ctx, "UPDATE votes SET count = count + 1 WHERE option = 'cat'")
		} else if vote == "dog" {
			dbConn.Exec(ctx, "UPDATE votes SET count = count + 1 WHERE option = 'dog'")
		}
	}
}
