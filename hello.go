package main

import "fmt"
import "github.com/go-redis/redis/v7"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import _ "github.com/lib/pq"

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	//var r = client.HGetAll("company:1")
	//fmt.Println(r)
	db1, err := sql.Open("mysql", "app:app@/app")
	if err != nil {
		return
	}
	defer db1.Close()
	err = db1.Ping()
	fmt.Println(err)

	db2, err := sql.Open("postgres",
		"host=localhost port=5432 user=app password=app dbname=app sslmode=disable")
	if err != nil {
		return
	}
	defer db2.Close()
	err = db2.Ping()
	fmt.Println(err)

}
