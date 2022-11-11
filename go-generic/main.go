package main

import "fmt"

func main() {
	total1 := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println("total1: ", total1)
}

func Sum(numbers []int) int {
	var total int
	for _, e := range numbers {
		total += e
	}
	return total
}
