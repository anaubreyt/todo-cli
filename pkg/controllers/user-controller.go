package controllers

import (
	"encoding/json"
	"net/http"
	"todo-cli/pkg/models"
	"todo-cli/pkg/utils"
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
	user := models.AddTask(addTask)
	res, _ := json.Marshal(user)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}
