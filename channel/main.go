/*
channel digunakan untuk menghubungkan goroutine satu dengan yang lain.
mekanisme yang dilakukan adalah serah - terima data dari channel tersebut.
mengiriman / penerimaan bersifat blocking (synchronous)
*/

/*
cahnnel merupakan sebuah variabel.
dibuat dengan menggunakan kombinasi keyword make dan chain

*/

package main
import "fmt"
import "runtime"
func main() {
	runtime.GOMAXPROCS(2)
	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello, %s ", who)
		messages <- data
	}

	go sayHelloTo("ilham muhamad fadilah")
	go sayHelloTo("agung nugraha")
	go sayHelloTo("putri mutiara")

	var message1 = <- messages
	fmt.Println(message1)
	var message2 = <- messages
	fmt.Println(message2)
	var message3 = <- messages
	fmt.Println(message3)
}