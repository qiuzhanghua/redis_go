package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)
import "github.com/go-redis/redis/v7"
import _ "github.com/go-sql-driver/mysql"
import _ "github.com/lib/pq"

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()
	//pong, err := client.Ping().Result()
	//fmt.Println(pong, err)
	////var r = client.HGetAll("company:1")
	////fmt.Println(r)
	//db1, err := sql.Open("mysql", "app:app@/app")
	//if err != nil {
	//	return
	//}
	//defer db1.Close()
	//err = db1.Ping()
	//fmt.Println(err)
	//
	//db2, err := sql.Open("postgres",
	//	"host=localhost port=5432 user=app password=app dbname=app sslmode=disable")
	//if err != nil {
	//	return
	//}
	//defer db2.Close()
	//err = db2.Ping()
	//fmt.Println(err)

	app := iris.New()
	//app.Logger().SetLevel("debug")

	app.Use(recover.New())
	//	app.Use(logger.New())

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		cmd := client.Get("key")
		if cmd.Err() != redis.Nil {
			ctx.WriteString(cmd.Val())
		} else {
			ctx.StatusCode(iris.StatusNotFound)
		}
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
