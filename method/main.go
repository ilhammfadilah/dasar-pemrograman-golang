/*
	method merupakan fungsi yang ada pada type.
	method bisa di akses melalui variabel objek.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var student1 = student{
		name: "ilham",
		grade: 9,
	}

	student1.sayHello()

	var nickname = student1.getNameAt(1)
	fmt.Println(nickname)
}

type student struct {
	name string
	grade int
}

func (s student) sayHello() {
	fmt.Println("Hallo ", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}