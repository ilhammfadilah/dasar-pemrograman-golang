package main
import (
	. "level-akses/library"
	"fmt"
)

func main() {
	// library.SayHello("ilham")
	// library.introduce("ilham")

	// var student1 = library.Student{"ilham", 9}
	// fmt.Println("student 1 name: ", student1.Name)
	// fmt.Println("student 1 grade: ", student1.Grade)

	var student2 = Student{"fadilah", 8}
	fmt.Println("student 2 name: ", student2.Name)
	fmt.Println("student 2 grade: ", student2.Grade)
}