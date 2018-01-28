package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	seed := time.Now().UnixNano()
	var message = omikuji(seed)
	fmt.Fprint(w, message)
}

func omikuji(seed int64) string {
	rand.Seed(time.Now().UnixNano())

	var saikoro int = rand.Intn(6) + 1
	var unsei string

	switch {
	case saikoro == 1:
		unsei = "凶"
	case 2 <= saikoro && saikoro <= 3:
		unsei = "吉"
	case 4 <= saikoro && saikoro <= 5:
		unsei = "中吉"
	case saikoro == 6:
		unsei = "大吉"
	}

	return unsei
}
