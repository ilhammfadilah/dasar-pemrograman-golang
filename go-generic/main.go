package main

import "fmt"

func main() {
	total1 := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println("total1: ", total1)

	total2 := Sum([]float32{2.5, 7.2})
	fmt.Println("total2: ", total2)

	total3 := Sum([]float64{1.23, 6.33, 12.6})
	fmt.Println("total3: ", total3)
}

func Sum[V int | float32 | float64](numbers []V) V {
	var total V
	for _, e := range numbers {
		total += e
	}
	return total
}
