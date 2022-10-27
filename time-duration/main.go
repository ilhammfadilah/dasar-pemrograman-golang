package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	time.Sleep(5 * time.Second)

	duration := time.Since(start)

	fmt.Println("Time elapsed in seconds ", duration.Seconds())
	fmt.Println("Time elapsed in minutes ", duration.Minutes())
	fmt.Println("Time elapsed in hours ", duration.Hours())
}
