package main

import (
	"math/rand"
	"time"
)

func randomDefer() {
	rand.Seed(time.Now().UnixNano())
	for rand.Intn(100) > 42 {
		defer func() {
			println("changkun/go-under-the-hood")
		}()
	}
}

func main() {
	randomDefer()
}
