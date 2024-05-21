package main

import (
	"fmt"
	"time"
)

func main() {
    // Membuat timer baru yang akan berhenti setelah 2 detik
    timer1 := time.NewTimer(2 * time.Second)

    // Menunggu hingga timer1 berakhir (selama 2 detik)
	// ".C" disini dimaksudkan bawaan dari time.NewTimer() yang menyediakan channel untuk mengirim waktu timer berhenti
    <-timer1.C 
    // Setelah timer1 berakhir, mencetak pesan "Timer 1 fired"
    fmt.Println("Timer 1 fired")

    // Membuat timer kedua yang akan berhenti setelah 1 detik
    timer2 := time.NewTimer(time.Second)
    // Menjalankan timer kedua di dalam goroutine
    go func() {
        // Menunggu hingga timer2 berakhir (selama 1 detik)
        <-timer2.C
        // Setelah timer2 berakhir, mencetak pesan "Timer 2 fired"
        fmt.Println("Timer 2 fired")
    }()
    // Menghentikan timer2 sebelum berakhir
    stop2 := timer2.Stop()
    // Jika timer2 berhasil dihentikan, mencetak pesan "Timer 2 stopped"
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }

    // Menunggu selama 2 detik sebelum program berakhir
    time.Sleep(2 * time.Second)
}
