package main

import (
	"bufio"
	"fmt"
	"github.com/k0kubun/pp/v3"
	"os"
	"todoproject/action"
	"todoproject/task"
)

func main() {
	comd()
	for {
		fmt.Println("Введите команду")
		cmd := bufio.NewScanner(os.Stdin)
		cmd.Scan()
		switch cmd.Text() {
		case "добавить задачу":
			task.AddTask()
			action.ActT(cmd.Text())
		case "удалить задачу":
			task.DeleteTask()
			action.ActT(cmd.Text())
		case "стоп":
			return
		case "сменить статус":
			task.ChangeStatus()
			action.ActT(cmd.Text())
		case "список":
			task.GetTask()
			action.ActT(cmd.Text())
		case "помощь":
			comd()
			action.ActT(cmd.Text())
		case "события":
			action.GetAction()
			action.ActT(cmd.Text())
		default:
			comd()
			action.ActF(cmd.Text(), "Такой команды нет")
			pp.Println("такой команды нет")
		}
	}
}

func comd() {
	pp.Println("список команд:")
	pp.Println("добавить задачу")
	pp.Println("удалить задачу")
	pp.Println("сменить статус")
	pp.Println("список")
	pp.Println("события")
	pp.Println("помощь")
}
