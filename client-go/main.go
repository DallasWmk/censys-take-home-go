package main

import (
    "github.com/go-redis/redis"
    "fmt"
)

func main() {
    // Create a new Redis client
    client := redis.NewClient(&redis.Options{
        Addr:     "10.1.0.2:6379",
        Password: "",
        DB:       0,
    })

    pong, err := client.Ping().Result()
    if err != nil {
        panic(err)
    }

    fmt.Println("Connected to Redis:", pong)

    // Close redis connection when finished
    defer client.Close()
}
