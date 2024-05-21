package main

import "fmt"

func main() {
    // Membuat channel "jobs" dengan buffer 5
    jobs := make(chan int, 5)
    // Membuat channel "done" untuk memberi tahu bahwa semua pekerjaan telah selesai
    done := make(chan bool)

    // Goroutine untuk menerima dan memproses pekerjaan dari channel "jobs"
    go func() {
        // Loop tak terbatas untuk menerima data dari channel "jobs"
        for {
            // Menerima data pekerjaan dari channel "jobs"
            j, more := <-jobs
            // Memeriksa apakah masih ada pekerjaan yang tersedia di channel "jobs"
            if more {
                // Jika masih ada pekerjaan, cetak informasi penerimaan pekerjaan
                fmt.Println("received job", j)
            } else {
                // Jika tidak ada pekerjaan lagi, cetak informasi bahwa semua pekerjaan telah diterima
                fmt.Println("received all jobs")
                // Mengirimkan sinyal bahwa semua pekerjaan telah selesai melalui channel "done"
                done <- true
                // Keluar dari goroutine
                return
            }
        }
    }()

    // Loop untuk mengirimkan pekerjaan ke channel "jobs"
    for j := 1; j <= 3; j++ {
        // Mengirimkan nomor pekerjaan ke channel "jobs"
        jobs <- j
        // Cetak informasi pengiriman pekerjaan
        fmt.Println("sent job", j)
    }
    // Menutup channel "jobs" setelah semua pekerjaan telah dikirim
    close(jobs)
    // Cetak informasi bahwa semua pekerjaan telah dikirim
    fmt.Println("sent all jobs")

    // Menunggu sinyal bahwa semua pekerjaan telah selesai dari channel "done"
    <-done

    // Menerima informasi tentang apakah masih ada pekerjaan yang tersisa di channel "jobs"
    _, ok := <-jobs
    // Cetak informasi tentang apakah masih ada pekerjaan yang tersisa di channel "jobs"
    fmt.Println("received more jobs:", ok)
}
