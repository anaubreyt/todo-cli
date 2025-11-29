package tasks

import (
	_ "fmt"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Status    bool
	Uuid      uuid.UUID
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
		Uuid:      uuid.New(),
		Name:      name,
		Descr:     descr,
		TimeStamp: time.Now(),
	}
}
