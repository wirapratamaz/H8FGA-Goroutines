package main

import (
	"fmt"
	"sync"
)

func main() {
	fruits := []string{"Apple", "Mango", "Strawberry", "Durian"}

	// WaitGroup adalah sebuah struktur data yang digunakan untuk menunggu (wait) satu atau lebih
	// goroutine untuk menyelesaikan tugas mereka sebelum proses selanjutnya dilakukan.
	var wg sync.WaitGroup //variable wg yang bertipe sync.WaitGroup

	//perulangan pada slice fruits
	for index, fruit := range fruits {
		//Setiap goroutine menambahkan 1 ke WaitGroup
		wg.Add(1)
		go printFruit(index, fruit, &wg)
	}
	// last -  menunggu sampai semua goroutine selesai
	wg.Wait()
}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("Index => %d, fruit => %s\n", index, fruit)
	//Ini mengurangi hitungan tunggu WaitGroup sebanyak 1
	wg.Done()
}
