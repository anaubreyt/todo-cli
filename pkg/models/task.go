package models

import (
	"fmt"
	"time"
	"todo-cli/pkg/config"
	_ "todo-cli/pkg/config"

	"gorm.io/gorm"
)

type Task struct {
	Status      bool      `json:"status"`
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Tm_stamp    time.Time `json:"-" gorm:"column:tm_stamp;type:timestamp"`
	Author      string    `json:"author"`

	// Для приема JSON
	TmStampJSON int64 `json:"tmStamp" gorm:"-"`
}

// Хук GORM BeforeCreate для установки текущего времени при создании task и правильного формата сохранения в БД
func (t *Task) BeforeCreate(tx *gorm.DB) error {
	// Если в JSON указан tmStamp
	if t.TmStampJSON > 0 {
		t.Tm_stamp = time.UnixMilli(t.TmStampJSON)
	}

	// Если в JSON tmStamp пуст ставится текущее время
	if t.Tm_stamp.IsZero() {
		t.Tm_stamp = time.Now()
	}
	return nil
}

// Хук GORM для корректной обработки JSON запроса GET из БД
// Для корректной работы time.Time в БД с GORM изменили тип колонки
// 1. ALTER TABLE task DROP COLUMN tm_stamp;
// 2. ALTER TABLE task ADD COLUMN tm_stamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP;
func (t *Task) AfterFind(tx *gorm.DB) error {
	fmt.Printf("DEBUG AfterFind: Task ID=%d, Tm_stamp=%v, IsZero=%v\n",
		t.Id, t.Tm_stamp, t.Tm_stamp.IsZero())
	if !t.Tm_stamp.IsZero() {
		t.TmStampJSON = t.Tm_stamp.UnixMilli()
		fmt.Printf("DEBUG AfterFind: Task ID=%v, TmStampJSON=%d\n",
			t.Id, t.TmStampJSON)
	} else {
		fmt.Printf("DEBUG AfterFInd: Task ID=%v, tm_stamp is zero!\n",
			t.Id)
	}
	return nil
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

func AddTask(task *Task) *Task {
	db := config.GetDB()
	db.Create(&task)
	return task
}

// localhost:8182/tasks/1 - for test in Insomnia
func DeleteTask(id int) error {
	db := config.GetDB()
	var task []Task
	result := db.First(&task, id)
	if result.Error != nil {
		return result.Error
	}
	result = db.Delete(&task)
	return result.Error
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
