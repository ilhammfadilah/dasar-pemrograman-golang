/*
	struct merupakan kumpulan property dan function yang dibungkus kedalam tipe data baru dengan nama tertentu.
*/

package main

import "fmt"

func main() {
	var s1 student
	s1.name = "ilham"
	s1.grade = 9

	var s2 = student{"muhamad", 10}

	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s1.grade)

	fmt.Println("name :", s2.name)
	fmt.Println("grade :", s2.grade)
}

type student struct {
	name string
	grade int
}
