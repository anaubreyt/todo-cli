package models

import (
	"context"
	"encoding/json"
	"fmt"
	_ "fmt"
	"time"
	"todo-cli/pkg/config"
	_ "todo-cli/pkg/config"
)

type Task struct {
	Status      bool   `json:"status"`
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tm_stamp    string `json:"tm_stamp"`
	Author      string `json:"author"`
}

func (Task) TableName() string {
	return "task"
}

type TaskList struct {
	List       []*Task
	NameOfList string
}

func GetTask() []Task {
	db := config.GetDB()
	var task []Task
	result := db.Find(&task)
	if result.Error != nil {
		fmt.Printf("Some error: %v", result.Error.Error())
	}
	return task
}

func GetSingleTask(id string) Task {
	ctx := context.Background()
	cl := config.GetRedis()
	var t Task
	val, err := cl.Get(ctx, id).Result()
	if err == nil {
		json.Unmarshal([]byte(val), &t)
		fmt.Printf("Value %v for id %v returned from cache\n", val, id)
		return t
	} else {
		db := config.GetDB()
		res := db.Where("id = ?", id)
		res = res.First(&t)
		if res.Error != nil {
			return Task{}
		} else {
			fmt.Printf("Value %v for key %v returned from database and stored in cache\n", t, id)
			jsonData, _ := json.Marshal(t)
			cl.Set(ctx, id, jsonData, time.Minute*10)
			return t
		}
	}
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
		Tm_stamp:    time.Now().String(),
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
