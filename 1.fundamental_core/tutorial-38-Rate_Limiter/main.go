package main

import (
	"fmt"
	"time"
)

func main() {
	// Membuat channel untuk menerima permintaan
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i // Mengirimkan permintaan ke channel
	}
	close(requests)

	// Membuat rate limiter yang membatasi laju 1 permintaan per 200 ms
	limiter := time.Tick(200 * time.Millisecond)

	// Memproses permintaan
	for req := range requests {
		<-limiter // Menunggu izin dari rate limiter
		fmt.Println("request", req, time.Now())
	}
}
