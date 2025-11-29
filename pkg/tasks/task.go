package tasks

import (
	"fmt"
	_ "fmt"
	"time"
)

type Task struct {
	Status    bool
	Id        int
	Name      string
	Descr     string
	TimeStamp time.Time
	Author    string
}

type TaskList struct {
	List       []*Task
	NameOfList string
}

func NewTask(name string, descr string) *Task {
	return &Task{
		Status:    false,
		Id:        0,
		Name:      name,
		Descr:     descr,
		TimeStamp: time.Now(),
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

func (t *TaskList) DeleteTaskFromList(position int) {
	fmt.Println(t.List)
	for _, elem := range t.List {
		if position == elem.Id {
			before := t.List[:position]
			after := t.List[position+1:]
			t.List = append(before, after...)
		}
	}
	fmt.Println(t.List)
}
