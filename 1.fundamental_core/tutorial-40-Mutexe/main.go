package main

import (
	"fmt"
	"sync"
)

// Container adalah struktur untuk menyimpan counter dengan sinkronisasi Mutex
type Container struct {
	mu       sync.Mutex   // Mutex untuk sinkronisasi akses ke counters
	counters map[string]int // Map untuk menyimpan counter untuk setiap nama
}

// inc adalah metode untuk menambahkan nilai counter untuk nama yang diberikan dalam Container
func (c *Container) inc(name string) {
	c.mu.Lock()         // Mengunci Mutex sebelum mengakses counters
	defer c.mu.Unlock() // Menggunakan defer untuk memastikan Mutex akan terlepas setelah fungsi selesai
	c.counters[name]++  // Menambahkan nilai counter untuk nama yang diberikan
}

func main() {
	// Inisialisasi Container dengan counter awal untuk setiap nama
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup // WaitGroup untuk menunggu semua goroutine selesai

	// doIncrement adalah fungsi untuk menambahkan nilai counter sebanyak n kali
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name) // Memanggil metode inc untuk menambah nilai counter sebanyak n kali
		}
		wg.Done() // Memberi tahu WaitGroup bahwa goroutine telah selesai
	}

	wg.Add(3)                   // Menambahkan 3 goroutine ke WaitGroup
	go doIncrement("a", 10000)  // Goroutine untuk menambah nilai counter "a" sebanyak 10000 kali
	go doIncrement("a", 10000)  // Goroutine untuk menambah nilai counter "a" sebanyak 10000 kali
	go doIncrement("b", 10000)  // Goroutine untuk menambah nilai counter "b" sebanyak 10000 kali

	wg.Wait()                  // Menunggu semua goroutine selesai
	fmt.Println(c.counters)    // Mencetak nilai counter setelah semua goroutine selesai
}
