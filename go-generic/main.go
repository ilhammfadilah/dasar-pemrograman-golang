package main

import "fmt"

func main() {
	// total1 := Sum([]int{1, 2, 3, 4, 5})
	// fmt.Println("total1: ", total1)

	// total2 := Sum([]float32{2.5, 7.2})
	// fmt.Println("total2: ", total2)

	// total3 := Sum([]float64{1.23, 6.33, 12.6})
	// fmt.Println("total3: ", total3)

	ints := map[string]int64{"first": 34, "second": 12}
	floats := map[string]float64{"first": 35.98, "second": 26.99}

	fmt.Printf("generic sums with constraint: %v and %v\n",
		SumNumbers1(ints),
		SumNumbers2(floats))
}

func Sum[V int | float32 | float64](numbers []V) V {
	var total V
	for _, e := range numbers {
		total += e
	}
	return total
}

func SumNumbers1(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumbers2[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, V := range m {
		s += V
	}
	return s
}
