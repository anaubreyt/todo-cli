package main

import (
	"fmt"
	"log"
	"net/http"
	_ "todo-cli/pkg/config"
	"todo-cli/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	routes.RegisterUserRoutes(route)
	http.Handle("/", route)
	fmt.Println("Listening on 8182...")
	log.Fatal(http.ListenAndServe("localhost:8182", route))
}

// есть конфигурационный файл
// в консоли выводим меню, там список действий:
// создать / посмотреть / изменить статус / редактировать / удалить
// выход из апп
