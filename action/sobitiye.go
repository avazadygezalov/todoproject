package action

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/k0kubun/pp/v3"
	"time"
)

type action struct {
	err  string
	data string
}

var actions = make(map[string]action)

func ActT(a string) {
	hash := md5.Sum([]byte(time.Now().Format("2006-01-02 15:04:05"))) // Сохраняем массив
	b := hex.EncodeToString(hash[:])
	actions[a+"|"+b] = action{
		data: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func ActF(a string, b string) {
	hash := md5.Sum([]byte(time.Now().Format("2006-01-02 15:04:05"))) // Сохраняем массив
	c := hex.EncodeToString(hash[:])
	actions[a+"|"+c] = action{
		err:  b,
		data: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func GetAction() {
	pp.Println(actions)
}
