package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Fungsi time.Sleep()
	// fmt.Println("start")
	// time.Sleep(time.Second * 4)
	// fmt.Println("afrer 4 second")

	// Scheduler Menggunakan time.Sleep()
	// for true {
	// 	fmt.Println("Assalamualaikum")
	// 	time.Sleep(1 * time.Second)
	// }

	// Fungsi time.NewTimer()
	// var timer = time.NewTimer(4 * time.Second)
	// fmt.Println("mulai")
	// <-timer.C
	// fmt.Println("selesai")

	// Fungsi time.AfterFunc()
	// var ch = make(chan bool)
	// time.AfterFunc(4*time.Second, func() {
	// 	fmt.Println("Expired")
	// 	// ch <- true
	// })

	// fmt.Println("start")
	// <-ch
	// fmt.Println("finish")

	// Scheduler Menggunakan Ticker

	// done := make(chan bool)
	// ticker := time.NewTicker(time.Second)

	// go func() {
	// 	time.Sleep(10 * time.Second)
	// 	done <- true
	// }()

	// for {
	// 	select {
	// 	case <-done:
	// 		ticker.Stop()
	// 		return
	// 	case t := <-ticker.C:
	// 		fmt.Println("Hello !!", t)
	// 	}
	// }

	// Kombinasi Timer & Goroutine

	var timeout = 5
	var ch = make(chan bool)

	go timer(timeout, ch)
	go watcher(timeout, ch)

	var input string
	fmt.Print("what is 725/25 ? ")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the answer was right!")
	} else {
		fmt.Println("the answer was wrong!")
	}

}

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\n time out ! no answer more than ", timeout, "seccond")
	os.Exit(0)
}
