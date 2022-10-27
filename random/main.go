package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(10)
	fmt.Println("random ke-1 : ", rand.Int())
	fmt.Println("random ke-2 : ", rand.Int())
	fmt.Println("random ke-3 : ", rand.Int())

	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
}
