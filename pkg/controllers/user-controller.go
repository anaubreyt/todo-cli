package controllers

import (
	"encoding/json" // для работы с JSON
	"fmt"
	"net/http"            // HTTP-сервер и клиент
	"todo-cli/pkg/models" // модели данных
	"todo-cli/pkg/utils"  // вспомогательные функции
)

func GetTasks(writer http.ResponseWriter, request *http.Request) {
	task := models.GetTask()
	res, _ := json.Marshal(task)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func AddTask(writer http.ResponseWriter, request *http.Request) {
	addTask := &models.Task{}
	utils.ParseBody(request, addTask)
	task := models.AddTask(addTask)
	res, _ := json.Marshal(task)
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func DeleteTask(writer http.ResponseWriter, request *http.Request) {
	deleteTask := &models.Task{}
	utils.ParseBody(request, &deleteTask)
	res := models.DeleteTask(deleteTask)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	fmt.Println(res)
}