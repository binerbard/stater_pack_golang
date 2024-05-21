package main

import (
	"bytes"
	"fmt"
	"text/template"
)

// Basic - Template Text

// 1. Define Struct for Person
type Person struct {
	Name string
	Age  int
}

func BasicTemplate(person Person) string {
	// 2. Make a Format Template
	// {{.Name}} : Person.Name
	// {{.Age}} : Person.Age
	const basicTemplate = `Hello, my name is {{.Name}} and I'm ({{.Age}}) years old.`

	// 3. Register template
	templateReg := template.New("BASIC_TEMPLATE")

	// 4. Parse format to register template
	templateParse, err := templateReg.Parse(basicTemplate)
	if err != nil {
		panic(err)
	}

	// 5. Execute the template with the data
	var output bytes.Buffer
	err = templateParse.Execute(&output, person)
	if err != nil {
		panic(err)
	}

	// 6. Return the generated string
	return output.String()
}

// 2. Conditional Text Template

type Item struct {
	Name  string
	Price int
}

type Cart struct {
	Items      []Item
	TotalPrice int
}

func ConditionalTemplate(cart Cart) string {
	// 2. Make a Format Template
	// {{- range .Items }} is mean fetch data from Items
	// {{ .Name }} : Item.Name
	// {{ .Price }} : Item.Price
	// {{ if gt .TotalPrice 100 }} is mean if .TotalPrice > 100
	// {{- end }}
	const conditionalTemplate = `Daftar Barang:
{{- range .Items }}
- Nama: {{ .Name }}, Harga: {{ .Price }}
{{- end }}

Total Harga: {{ .TotalPrice }}
{{ if gt .TotalPrice 100 }}
Diskon diberikan!
{{ else }}
Tidak ada diskon.
{{ end }}`

	// 3. Register template
	templateReg := template.New("CONDITIONAL_TEMPLATE")

	// 4. Parse format to register template
	templateParse, err := templateReg.Parse(conditionalTemplate)
	if err != nil {
		panic(err)
	}

	// 5. Execute the template with the data
	var output bytes.Buffer
	err = templateParse.Execute(&output, cart)
	if err != nil {
		panic(err)
	}

	// 6. Return the generated string
	return output.String()
}

func main() {
	// Call BasicTemplate and print the result
	// Create an instance of Person with data

	fmt.Println("BASIC TEXT TEMPLATE EXAMPLE :")
	person := Person{
		Name: "John Doe",
		Age:  30,
	}
	result := BasicTemplate(person)
	fmt.Println(result)
	
	// Call ConditionalTemplate and print the result
	// Create an instance of Cart with data
	fmt.Println("\nCONDITIONAL TEXT TEMPLATE : ")
	cart := Cart{
		Items: []Item{
			{Name: "Item1", Price: 20},
			{Name: "Item2", Price: 50},
			{Name: "Item3", Price: 100},
		},
		TotalPrice: 170,
	}
	result = ConditionalTemplate(cart)
	fmt.Println(result)
}

// Note :
// If you understand this concept 
// Maybe you can implement this case to create file template like 
// file.go, file.js or other