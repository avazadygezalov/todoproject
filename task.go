package main

import (
	"bufio"
	"github.com/k0kubun/pp/v3"
	"os"
	"strings"
	"time"
)

func addTask(task map[string]toDoList) {
	scan := bufio.NewScanner(os.Stdin)
	pp.Println("Введите название задачи которую хотите добавить (1 слово) а дальше описание")
	scan.Scan()
	text := strings.Fields(scan.Text())
	var descr string
	for i := 1; i < len(text); i++ {
		if i != len(text)-1 {
			descr = descr + text[i] + " "
		} else {
			descr = descr + text[i]
		}
	}
	task[text[0]] = toDoList{
		description: descr,
		dataCreate:  time.Now().Format("2006-01-02 15:04:05"),
	}
	return
}

func deleteTask(task map[string]toDoList) {
	scan := bufio.NewScanner(os.Stdin)
	pp.Println("Введите название задачи котрую хотите удалить")
	scan.Scan()
	delete(task, scan.Text())
}
