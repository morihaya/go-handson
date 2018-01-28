package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 現在時刻を数値で取得する
	t := time.Now().UnixNano()
	// 乱数のたねを設定
	rand.Seed(t)
	// xは0-6の間の値になる

	s := rand.Intn(6)
	println(s)
	x := s + 1
	//x := 6

	var msg string
	if x == 6 {
		msg = "大吉"
	} else if x == 5 || x == 4 {
		msg = "中吉"
	} else if x == 2 || x == 3 {
		msg = "小吉"
	} else if x == 1 {
		msg = "凶"
	}
	fmt.Println(msg)
}
