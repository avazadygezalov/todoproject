package action

import (
	"github.com/k0kubun/pp/v3"
	"time"
)

type action struct {
	name string
	err  string
	data string
}

var actions = make([]action, 0)

func ActT(a string) {
	actions = append(actions, action{
		name: a,
		data: time.Now().Format("2006-01-02 15:04:05"),
	})
}

func ActF(a string, b string) {
	actions = append(actions, action{
		name: a,
		err:  b,
		data: time.Now().Format("2006-01-02 15:04:05"),
	})
}

func GetAction() {
	pp.Println(actions)
}
