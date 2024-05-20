package main

import (
	"fmt"
)

func main() {
    // String
    greeting := "Hello, 世界"
    fmt.Println("String:", greeting)

    // Menghitung panjang string dalam byte
    fmt.Println("Length in bytes:", len(greeting))

    // Iterasi string sebagai byte
    fmt.Println("Bytes in string:")
    for i := 0; i < len(greeting); i++ {
        fmt.Printf("%x ", greeting[i])
    }
    fmt.Println()

    // Iterasi string sebagai rune
    fmt.Println("Runes in string:")
    for i, r := range greeting {
        fmt.Printf("%d: %c ", i, r)
    }
    fmt.Println()

    // Rune
    ch := 'A'
    fmt.Printf("Rune: %c, Unicode code point: %U\n", ch, ch)

    // Mengubah string menjadi slice of rune
    runes := []rune(greeting)
    fmt.Println("Slice of runes:")
    for i, r := range runes {
        fmt.Printf("%d: %c ", i, r)
    }
    fmt.Println()
}
