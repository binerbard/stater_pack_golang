package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// readOp adalah struktur untuk operasi baca
type readOp struct {
	key  int
	resp chan int
}

// writeOp adalah struktur untuk operasi tulis
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readOps uint64 // Jumlah operasi baca
	var writeOps uint64 // Jumlah operasi tulis

	reads := make(chan readOp) // Channel untuk operasi baca
	writes := make(chan writeOp) // Channel untuk operasi tulis

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key] // Mengirim nilai ke channel respons
			case write := <-writes:
				state[write.key] = write.val // Menulis nilai ke state
				write.resp <- true // Mengirim pesan berhasil ke channel respons
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read // Mengirim operasi baca ke channel reads
				<-read.resp // Menerima respons dari channel
				atomic.AddUint64(&readOps, 1) // Menambah jumlah operasi baca
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write // Mengirim operasi tulis ke channel writes
				<-write.resp // Menerima respons dari channel
				atomic.AddUint64(&writeOps, 1) // Menambah jumlah operasi tulis
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps) // Mengambil jumlah operasi baca akhir
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps) // Mengambil jumlah operasi tulis akhir
	fmt.Println("writeOps:", writeOpsFinal)
}
