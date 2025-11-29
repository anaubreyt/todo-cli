package main

import (
	"todo-cli/pkg/tasks"
)

func main() {
	// fmt.Println("hello")
	task := tasks.NewTask("Cup of tea", "let's go")
	task2 := tasks.NewTask("Cup of tea2", "let's go2")
	task3 := tasks.NewTask("Cup of tea3", "let's go3")
	to_do := tasks.NewTaskList("Home work")
	to_do.AddTaskToList(task)
	to_do.AddTaskToList(task2)
	to_do.AddTaskToList(task3)
	to_do.DeleteTaskFromList(0)
}

// есть конфигурационный файл
// в консоли выводим меню, там список действий:
// создать / посмотреть / изменить статус / редактировать / удалить
// выход из апп
