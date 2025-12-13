package controllers

import (
	"encoding/json"
	"net/http"
	"todo-cli/pkg/models"
)

func GetTasks(writer http.ResponseWriter, request *http.Request) {
	task := models.GetTask()

	res, _ := json.Marshal(task)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}
