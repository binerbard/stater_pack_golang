package main

import (
	"cmp"    // Import paket cmp untuk perbandingan
	"fmt"    // Import paket fmt untuk mencetak output
	"slices" // Import paket slices untuk operasi pada slice
)

func main() {
	// Membuat slice fruits dengan beberapa buah
	fruits := []string{"peach", "banana", "kiwi"}

	// Mendefinisikan fungsi perbandingan berdasarkan panjang string
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b)) // Membandingkan panjang string a dan b
	}

	// Mengurutkan slice fruits berdasarkan panjang string menggunakan lenCmp
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits) // Mencetak hasil urutan fruits

	// Mendefinisikan tipe data Person dengan field name dan age
	type Person struct {
		name string
		age  int
	}

	// Membuat slice people dengan beberapa objek Person
	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}

	// Mengurutkan slice people berdasarkan umur menggunakan fungsi anonim
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age) // Membandingkan umur a dan b
		})
	fmt.Println(people) // Mencetak hasil urutan people
}
