package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Concurrency test")

	messages := make(chan int, 1)

	go putMessages(messages, 5)
	getMessages(messages)
}

func putMessages(c chan<- int, n int) {
	var i int
	for i = 1; i <= n; i++ {
		c <- i
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

func getMessages(c <-chan int) {
	var closed int = 0
	for {
		msg, state := <-c
		switch {
		case state:
			fmt.Println("Got message:", msg)
		case !state:
			closed = 1
		}
		if closed == 1 {
			break
		}
	}
}
