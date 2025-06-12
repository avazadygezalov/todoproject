package main

import (
	"bufio"
	"github.com/k0kubun/pp/v3"
	"os"
)

type toDoList struct {
	description string
	dataCreate  string
	dataDone    string
	status      bool
}

var list = make(map[string]toDoList)

func main() {
	for {
		pp.Println("Введите команду")
		cmd := bufio.NewScanner(os.Stdin)
		cmd.Scan()
		switch cmd.Text() {
		case "добавить задачу":
			addTask(list)
		case "удалить задачу":
			deleteTask(list)
		case "стоп":
			return
		}
		pp.Println(list)
	}
}
