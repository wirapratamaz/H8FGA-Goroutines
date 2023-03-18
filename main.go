package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("main execution started")

	// dua fungsi secara concurrent atau parallel
	go firstProcess(8)
	secondProcess(8)

	// menampilkan jumlah goroutine yang sedang berjalan
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	// pemanggilan fungsi jeda
	time.Sleep(time.Second * 2)

	// setelah semua goroutine selesai dieksekusi
	fmt.Println("Main execution ended")
}

func firstProcess(index int) {
	fmt.Println("First process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("First Process func ended")
}
func secondProcess(index int) {
	fmt.Println("second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("second Process func ended")
}
