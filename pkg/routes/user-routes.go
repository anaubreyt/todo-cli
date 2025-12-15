package routes

import (
	"todo-cli/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", controllers.AddTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")
}
