package main

import (
	"fmt"
	"time"
)

func main() {
	var time1 = time.Now()
	fmt.Printf("time 1 %v \n", time1)
	fmt.Println("tahun :", time1.Year())

	var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
	fmt.Printf("time2 %v\n", time2)

	// Parsing dari string ke time.Time
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2015-09-02 08:04:00"

	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())

	layoutFormat = "02/01/2006 MST"
	value = "02/09/2015 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())

	// Predefined Layout Format Untuk Keperluan Parsing Time
	var tanggal, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
	fmt.Println(tanggal.String())

	// Format dari time.Time ke string
	var tgl, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
	var dateS1 = tgl.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dateS1", dateS1)

	var dateS2 = tgl.Format(time.RFC3339)
	fmt.Println("dateS2", dateS2)
}
