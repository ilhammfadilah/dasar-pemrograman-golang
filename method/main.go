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

	var murid1 = murid{"ilham", 9}
	fmt.Println("murid1 before change", murid1.name)

	murid1.changeName1("fadilah")
	fmt.Println("murid1 after change", murid1.name)

	murid1.changeName2("muhamad")
	fmt.Println("murid1 after change 2", murid1.name)
}

type student struct {
	name string
	grade int
}

type murid struct {
	name string
	grade int
}

func (s student) sayHello() {
	fmt.Println("Hallo ", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func (s murid) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
}

func (s *murid) changeName2(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}