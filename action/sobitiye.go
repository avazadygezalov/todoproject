package action

import "time"

type action struct {
	err  string
	data string
}

var actions = make(map[string]action)

func ActT(a string) {
	actions[a] = action{
		data: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func ActF(a string, b string) {
	actions[a] = action{
		err:  b,
		data: time.Now().Format("2006-01-02 15:04:05"),
	}
}
