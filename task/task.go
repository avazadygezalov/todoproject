package task

import (
	"bufio"
	"github.com/k0kubun/pp/v3"
	"os"
	"strings"
	"time"
	"todoproject/action"
)

type toDoList struct {
	description string
	dataCreate  string
	dataDone    string
	status      bool
}

var list = make(map[string]toDoList)

func AddTask() {
	scan := bufio.NewScanner(os.Stdin)
	pp.Println("Введите название задачи которую хотите добавить (1 слово) а дальше описание")
	scan.Scan()
	action.ActT(scan.Text())
	text := strings.Fields(scan.Text())
	var descr string
	for i := 1; i < len(text); i++ {
		if i != len(text)-1 {
			descr = descr + text[i] + " "
		} else {
			descr = descr + text[i]
		}
	}
	list[text[0]] = toDoList{
		description: descr,
		dataCreate:  time.Now().Format("2006-01-02 15:04:05"),
	}
	return
}

func DeleteTask() {
	scan := bufio.NewScanner(os.Stdin)
	pp.Println("Введите название задачи которую хотите удалить")
	scan.Scan()
	action.ActT(scan.Text())
	delete(list, scan.Text())
}

func GetTask() {
	pp.Println(list)
}

func ChangeStatus() {
	scan := bufio.NewScanner(os.Stdin)
	pp.Println("Введите название задачи у которой хотите изменить статус")
	scan.Scan()
	action.ActT(scan.Text())
	taskName := scan.Text()
	pp.Println("Задача выполнена? да нет")
	scan.Scan()
	action.ActT(scan.Text())
	switch scan.Text() {
	case "да":
		list[taskName] = toDoList{
			status: true,
		}
	case "нет":
		list[taskName] = toDoList{
			status: false,
		}
	default:
		pp.Println("Вы ввели не то")
		action.ActF(taskName, "Вы ввели не то в качестве ответа на вопрос Задача выполнена?")
	}
}
