package controllers

import (
	"encoding/json"
	"log"
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
	task := models.AddTask(addTask)
	res, err := json.Marshal(task)
	// Добавили проверку на прохождение сериализации
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		log.Printf("JSON marshal error: %v", err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	// Залогируем ошибку от .Write(res) на случай разрыва соединения
	// Но статус код выше не поменяется!
	_, err = writer.Write(res)
	if err != nil {
		log.Printf("Writer response error: %v", err)
	}
}
