package main

import (
	"fmt"
	"time"
)

func main() {
	// start := time.Now()

	// time.Sleep(5 * time.Second)

	// duration := time.Since(start)

	// fmt.Println("Time elapsed in seconds ", duration.Seconds())
	// fmt.Println("Time elapsed in minutes ", duration.Minutes())
	// fmt.Println("Time elapsed in hours ", duration.Hours())

	// Kalkulasi Durasi Antara 2 Objek Waktu

	// time1 := time.Now()
	// time.Sleep(5 * time.Second)
	// time2 := time.Now()

	// duration := time2.Sub(time1)
	// fmt.Println("Time elapsed in seconds ", duration.Seconds())
	// fmt.Println("Time elapsed in minutes ", duration.Minutes())
	// fmt.Println("Time elapsed in hours ", duration.Hours())

	// Konversi Angka ke time.Duration

	var n time.Duration = 60
	duration := n * time.Second
	fmt.Println(duration)
}
