// Map adalah tipe data asosiatif yang ada di Go, berbentuk key-value pair.

package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("februari", chicken["februari"])
}