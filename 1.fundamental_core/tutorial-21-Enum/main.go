package main

import "fmt"

// Enumerasi Status
const (
    Pending = iota
    Approved
    Rejected
    Cancelled
)

const (
    ReadMode = 1 << iota
    WriteMode
    ExecuteMode
)


func main() {
    fmt.Println("Status Pending:", Pending)
    fmt.Println("Status Approved:", Approved)
    fmt.Println("Status Rejected:", Rejected)
    fmt.Println("Status Cancelled:", Cancelled)

	// Example Enumerate bit shifting
    var mode = ReadMode | ExecuteMode
    fmt.Printf("Mode: %b\n", mode)

    // Memeriksa apakah mode tertentu diaktifkan
    fmt.Println("Read Mode Activated:", mode&ReadMode == ReadMode)
    fmt.Println("Write Mode Activated:", mode&WriteMode == WriteMode)
    fmt.Println("Execute Mode Activated:", mode&ExecuteMode == ExecuteMode)
}
