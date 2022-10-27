package main

import (
	"fmt"
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
	var ch = make(chan bool)
	time.AfterFunc(4*time.Second, func() {
		fmt.Println("Expired")
		// ch <- true
	})

	fmt.Println("start")
	<-ch
	fmt.Println("finish")
}
