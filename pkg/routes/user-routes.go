package routes

import (
	"todo-cli/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterTaskRoutes = func(router *mux.Router) {
	//Get tasks
	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	//Post tasks
	router.HandleFunc("/tasks", controllers.AddTask).Methods("POST")
	//Delete tasks
	router.HandleFunc("/delete", controllers.DeleteTask).Methods("DELETE")
	//Update tasks
	router.HandleFunc("/tasks", controllers.AddTask).Methods("PATCH")
}

// 
// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
// в качестве handler func(ResponseWriter, *Request) как раз и выступает controllers.AddTask

// хендлер имеет два параметра:
// ResponseWriter - поток ответа 
// *Request - информация о запросе