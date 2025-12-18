package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo-cli/internal/redisCache"
	_ "todo-cli/pkg/config"
	"todo-cli/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	routes.RegisterUserRoutes(route)
	http.Handle("/", route)
	
	err := redisCache.SetKey("todo1", "{task: 1}", 0)
	if err != nil {
		log.Printf("ERROR %v", err)
	}
	
	val, err := redisCache.GetKey("todo1")
	if err != nil {
		log.Printf("ERROR %v", err)
	}
	
	fmt.Println("[RESULT REDIS]: ", val)
	
	log.Fatal(http.ListenAndServe(os.Getenv("HANDLER_URL"), route))
}

// есть конфигурационный файл
// в консоли выводим меню, там список действий:
// создать / посмотреть / изменить статус / редактировать / удалить
// выход из апп
