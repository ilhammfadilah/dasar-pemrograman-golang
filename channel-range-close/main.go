package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)
	var pesan = make(chan string)

	go sendMessage(messages)
	printMessage(messages)

	go mengirimPesan(pesan)
	menerimaPesan(pesan)
}

func sendMessage(messages chan<- string) {
	for i := 0; i < 20; i++ {
		messages <- fmt.Sprintf("data %d", i)
	}
	close(messages)
}

func printMessage(messages <-chan string) {
	for message := range messages {
		fmt.Println(message)
	}
}

func mengirimPesan(p chan<- string) {
	p <- fmt.Sprintln("i love you")
	close(p)
}

func menerimaPesan(p <-chan string) {
	fmt.Println(p)
}
