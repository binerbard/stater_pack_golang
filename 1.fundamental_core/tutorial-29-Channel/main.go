package main

import (
	"fmt"
)

// Struct untuk merepresentasikan data mahasiswa
type Student struct {
	Name  string
	Grade int
}

// Fungsi untuk mengirimkan data ke channel tanpa buffer
func sendData(ch chan<- Student, student Student) {
	ch <- student // Mengirimkan data ke channel
}

// Fungsi untuk menerima data dari channel
func receiveData(ch <-chan Student) Student {
	return <-ch // Menerima data dari channel
}

func main() {
	// Membuat channel tanpa buffer
	unbufferedCh := make(chan Student)
	// Membuat channel dengan buffer 2
	bufferedCh := make(chan Student, 2)

	// Goroutine untuk mengirimkan data ke channel tanpa buffer
	go sendData(unbufferedCh, Student{Name: "John", Grade: 90})
	// Goroutine untuk mengirimkan data ke channel dengan buffer
	go sendData(bufferedCh, Student{Name: "Doe", Grade: 85})

	// Menerima data dari channel tanpa buffer
	receivedStudent1 := receiveData(unbufferedCh)
	fmt.Println("Received from unbuffered channel:", receivedStudent1)

	// Menerima data dari channel dengan buffer
	receivedStudent2 := receiveData(bufferedCh)
	fmt.Println("Received from buffered channel:", receivedStudent2)
}
