package main

import (
	"fmt"
	"os"
	s "strings"
)



type point struct {
	x, y int
}
var p = fmt.Println


func main() {

	//  STRING FUNCTIONS  
    p("1. Contains:  ", s.Contains("test", "es"))
    p("2. Count:     ", s.Count("test", "t"))
    p("3. HasPrefix: ", s.HasPrefix("test", "te"))
    p("4. HasSuffix: ", s.HasSuffix("test", "st"))
    p("5. Index:     ", s.Index("test", "e"))
    p("6. Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("7. Repeat:    ", s.Repeat("a", 5))
    p("8. Replace:   ", s.Replace("foo", "o", "0", -1))
    p("9. Replace:   ", s.Replace("foo", "o", "0", 1))
    p("10. Split:     ", s.Split("a-b-c-d-e", "-"))
    p("11. ToLower:   ", s.ToLower("TEST"))
    p("12. ToUpper:   ", s.ToUpper("test"))


	// STRING FORMAT

    p := point{1, 2}
    fmt.Printf("13. struct1: %v\n", p)

    fmt.Printf("14. struct2: %+v\n", p)

    fmt.Printf("15. struct3: %#v\n", p)

    fmt.Printf("16. type: %T\n", p)

    fmt.Printf("17. bool: %t\n", true)

    fmt.Printf("18. int: %d\n", 123)

    fmt.Printf("19. bin: %b\n", 14)

    fmt.Printf("20. char: %c\n", 33)

    fmt.Printf("21. hex: %x\n", 456)

    fmt.Printf("22. float1: %f\n", 78.9)

    fmt.Printf("23. float2: %e\n", 123400000.0)
    fmt.Printf("24. float3: %E\n", 123400000.0)

    fmt.Printf("25. str1: %s\n", "\"string\"")

    fmt.Printf("26. str2: %q\n", "\"string\"")

    fmt.Printf("27. str3: %x\n", "hex this")

    fmt.Printf("28. pointer: %p\n", &p)

    fmt.Printf("29. width1: |%6d|%6d|\n", 12, 345)

    fmt.Printf("30. width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

    fmt.Printf("31. width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

    fmt.Printf("32. width4: |%6s|%6s|\n", "foo", "b")

    fmt.Printf("33. width5: |%-6s|%-6s|\n", "foo", "b")

    s := fmt.Sprintf("34. sprintf: a %s", "string")
    fmt.Println(s)

    fmt.Fprintf(os.Stderr, "35. Fprintf io: an %s\n", "error")
}