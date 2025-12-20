package main

import (
	"log"
	"net/http"
	_ "todo-cli/pkg/config"
	"todo-cli/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	routes.RegisterTaskRoutes(route)
	http.Handle("/", route)
	log.Fatal(http.ListenAndServe("localhost:8182", route))
}

// есть конфигурационный файл
// в консоли выводим меню, там список действий:
// создать / посмотреть / изменить статус / редактировать / удалить
// выход из апп


// route.HandleFunc("/test", func (writer http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("test test testik")
// 		writer.Header().Set("Content-Type", "application/json")
// 		writer.WriteHeader(http.StatusOK)
// 		s := "Привет, мир!"
// 		b := []byte(s)
// 		writer.Write(b)
// })