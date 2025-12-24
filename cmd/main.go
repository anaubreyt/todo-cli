package main

import (
	"fmt"
	// _ "log"
	// _ "net/http"
	// _ "os"
	"time"
	// _ "todo-cli/internal/redisCache"
	// _ "todo-cli/pkg/config"
	// _ "todo-cli/pkg/routes"

	// _ "github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"todo-cli/pkg/rabbit"
)

func main() {
	conn, err := rabbit.GetConn("amqp://guest:guest@localhost")
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			time.Sleep(time.Second)
			conn.Publish("test-key", []byte(`{"message":"test"}`))
		}
	}()

	err = conn.StartConsumer("test-queue", "test-key", handler, 2)

	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	<-forever

	// route := mux.NewRouter()
	// routes.RegisterUserRoutes(route)
	// http.Handle("/", route)

	// err = redisCache.SetKey("todo1", "{task: 1}", 0)
	// if err != nil {
	// 	log.Printf("ERROR %v", err)
	// }

	// val, err := redisCache.GetKey("todo1")
	// if err != nil {
	// 	log.Printf("ERROR %v", err)
	// }

	// fmt.Println("[RESULT REDIS]: ", val)

	// log.Fatal(http.ListenAndServe(os.Getenv("HANDLER_URL"), route))
}

func handler(d amqp.Delivery) bool {
	if d.Body == nil {
		fmt.Println("Error, no message body!")
		return false
	}
	fmt.Println(string(d.Body))
	return true
}

// есть конфигурационный файл
// в консоли выводим меню, там список действий:
// создать / посмотреть / изменить статус / редактировать / удалить
// выход из апп
