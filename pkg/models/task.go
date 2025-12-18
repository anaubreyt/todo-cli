package models

import (
	"fmt"
	_ "fmt"
	"log"
	"strconv"
	"time"
	"todo-cli/internal/redisCache"
	"todo-cli/pkg/config"
	_ "todo-cli/pkg/config"
)

type Task struct {
	Status      bool      `json:"status"`
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Tm_stamp    time.Time `json:"tmStamp"`
	Author      string    `json:"author"`
}

func (Task) TableName() string {
	return "task"
}

type TaskList struct {
	List       []*Task
	NameOfList string
}

func GetTaskRedis(key string) any {
	task, err := redisCache.GetKey(key)
	if err != nil {
		log.Fatal(err)
	}
	return task
}

func GetTaskPostgressBYID(key int) any {
	db := config.GetDB()
	var task []Task
	result := db.Find(&task)
	if result.Error != nil {
		fmt.Printf("Some error: %v", result.Error.Error())
	}
	return task
}
func GetTask() []Task {

	var task []Task

	return task
}
func GetTask_byID(key string) any {
	var task any
	task = GetTaskRedis(key)
	if task != nil {
		return task
	}
	id, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println("Error:", err)
	}
	task = GetTaskPostgressBYID(id)

	if task == nil {
		panic(err)
	}

	redisCache.SetKey(key, task, 0)

	return task
}

func AddTask(task *Task) *Task {
	db := config.GetDB()
	db.Create(&task)
	return task
}

func NewTask(name string, descr string) *Task {
	return &Task{
		Status:      false,
		Id:          0,
		Name:        name,
		Description: descr,
		Tm_stamp:    time.Now(),
	}
}

func NewTaskList(name string) *TaskList {
	return &TaskList{NameOfList: name}
}

func (t *TaskList) AddTaskToList(elem *Task) {
	position := len(t.List)
	elem.Id = position
	t.List = append(t.List, elem)
}

func (t *TaskList) PrintList() {
	for _, v := range t.List {
		fmt.Printf("Position: %v, name %v\n", v.Id, v.Name)
	}
}

func (t *TaskList) DeleteTaskFromList(position int) {
	if position >= len(t.List) {
		err := fmt.Errorf("position %v outside of list range", position)
		fmt.Println("error when deleting:", err.Error())
		return
	}
	before := t.List[:position]
	for _, elem := range t.List {
		if elem.Id > position {
			elem.Id--
		}
	}
	after := t.List[position+1:]
	t.List = append(before, after...)
}
