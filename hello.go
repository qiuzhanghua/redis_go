package main

import "fmt"
import "github.com/go-redis/redis/v7"

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	//var r = client.HGetAll("company:1")
	//fmt.Println(r)
}
