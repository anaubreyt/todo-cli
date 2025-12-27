package controllers

import (
	"encoding/json" // для работы с JSON
	"fmt"
	"log"
	"net/http" // HTTP-сервер и клиент
	"strconv"
	"todo-cli/pkg/models" // модели данных
	"todo-cli/pkg/utils"  // вспомогательные функции

	"github.com/gorilla/mux"
)

func GetTasks(writer http.ResponseWriter, request *http.Request) {
	task := models.GetTask()
	res, _ := json.Marshal(task)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func GetTaskById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	idNum, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Ошибка приведения id к числу при запросе таски по id: %v", err)
	}
	task := models.GetTaskById(idNum)
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
	// с помощью библиотеки "github.com/gorilla/mux" разбираем параметры из урла в мапу
	vars := mux.Vars(request)
	// достаём нужный нам параметр
	id := vars["id"]
	// создаём переменную со структурой таски
	deleteTask := &models.Task{}
	// вот тут важно, у нас Id - число по типу, а в урле всё строкой приходит, поэтому нужно привести к числовому типу
	idNum, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Ошибка приведения id к числу при удалении: %v", err)
	}
	// задаём айдищник в существующей структуре
	deleteTask.Id = idNum
	// дальше сумрачная магия и спокойно удаляем таску
	utils.ParseBody(request, &deleteTask)
	res := models.DeleteTask(deleteTask)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	fmt.Println(res)
}

func UpdateTask(writer http.ResponseWriter, request *http.Request) {
	updateTask := &models.Task{}
	utils.ParseBody(request, updateTask)
	vars := mux.Vars(request)
	id := vars["id"]
	idNum, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Ошибка приведения id к числу при обновлении: %v", err)
	}
	updateTask.Id = idNum
	task := models.UpdateTask(updateTask)
	res, _ := json.Marshal(task)
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}