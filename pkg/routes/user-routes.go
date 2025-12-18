package routes

import (
	"todo-cli/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router) {
	//Get tasks
	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/task/{taskId:[0-9]+}", controllers.GetTask).Methods("GET")
	router.HandleFunc("/tasks", controllers.AddTask).Methods("POST")
}
