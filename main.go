package main

import (
	"fmt"
	"os"
)

func main() {
	//*case 1 defer
	//block sebuah function hingga selesai
	defer fmt.Println("defer functioin starts to execute")
	fmt.Println("Hai everyone!!")
	fmt.Println("Welcome back to course Golang by Hactiv 8")

	//*case 2 defer
	callDefereFunc() //output yg dikeluarkan akan lebih dulu
	fmt.Println("Hello Everyone")

	//*case 3 exit
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before exiting")
	//menghentikan suatu program secara paksa
	os.Exit(1)
}

// * case 2 defer
func callDefereFunc() {
	defer deferFunc() //dipanggil dalam block function ini bukan main
}

func deferFunc() {
	fmt.Println("Defer function starts to execute")
}
