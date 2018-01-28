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

	/*
		for i := 0; i < 7; i = i + 1 {
			s := rand.Intn(6)
			//println(s)
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
			fmt.Println(i, "日目", msg)
		}
	*/
	i := 0
	for {
		s := rand.Intn(6)
		//println(s)
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
		fmt.Println(i, "日目", msg)
		i = i + 1
		if x == 6 {
			break
		}
	}
}
