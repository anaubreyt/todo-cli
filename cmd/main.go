package main

import (
	"fmt"
	"todo-cli/pkg/tasks"
)

func main() {
	// fmt.Println("hello")
	task := tasks.NewTask("Cup of tea", "let's go")
	fmt.Println(task)
}

// есть конфигурационный файл
// в консоли выводим меню, там список действий:
// создать / посмотреть / изменить статус / редактировать / удалить
// выход из апп
