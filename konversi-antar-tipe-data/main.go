package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string -> int
	var number = "29"

	var num, err = strconv.Atoi(number)

	if err == nil {
		fmt.Println(num)
	}
}
