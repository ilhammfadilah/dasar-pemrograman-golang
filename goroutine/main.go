/*
	goroutine mirip seperti thread, tapi sebenarnya bukan.
	sebuah native trhead bisa memiliki banyak goroutine.
	lebih cocok dipanggil sebagai mini thread.
	hanya membutuhkan 2kB memori untuk satu buah goroutine.
	bersifat asynchronous.
*/ 

package main
import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2) // jumlah core untuk eksekusi program

	go print(5, "hallo") // di ekseskusi sebagai goroutine baru
	print(5, "assalamualaikum")

	var input string
	fmt.Scanln(&input) // menghentikan proses
}
/*
untuk menerapkan goroutine, proses yang akan di eksekusi harus dibungkus kedalam fungsi
pada saat eksekusi fungsi ditambah keyword go didepannya
*/