package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr: "192.168.1.24:6380",
})

func main() {
	subscriber := redisClient.Subscribe(ctx, "app:notifications")

	user := User{}

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		fmt.Println("Received message from " + msg.Channel + " channel.")
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal([]byte(msg.Payload), &user); err != nil {
			panic(err)
		}

		fmt.Println("Received message from " + msg.Channel + " channel.")
		fmt.Printf("%+v\n", user)
	}
}
