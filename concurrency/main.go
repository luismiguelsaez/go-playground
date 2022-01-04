package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Concurrency test")

	messages := make(chan int, 1)

	go putMessages(messages, 10)
	go getMessages(messages)

	time.Sleep(time.Second * 60)
}

func putMessages(c chan<- int, n int) {
	var i int
	for i = 1; i <= n; i++ {
		c <- i
		time.Sleep(time.Millisecond * 500)
	}
}

func getMessages(c <-chan int) {
	for {
		msg := <-c
		fmt.Println("Got message:", msg)
	}
}
