package main

import (
	"errors"
	"fmt"
)

// Error
func f(arg int) (int, error) {
    if arg == 42 {

        return -1, errors.New("can't work with 42")
    }

    return arg + 3, nil
}

var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
    if arg == 2 {
        return ErrOutOfTea
    } else if arg == 4 {

        return fmt.Errorf("making tea: %w", ErrPower)
    }

	// A "nil" value in the error position indicates that there was no error.
    return nil
}



// Custom Error

type argError struct {
    arg     int
    message string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func fcustom(arg int) (int, error) {
    if arg == 42 {

        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}


func main() {

	fmt.Println("Errors:")
	// Make a loop to fetch [7,42]
    for _, i := range []int{7, 42} {

		// r, e := f(i); => Call function f(v) with return vallue r and e
		// e != nil => Check if function e is not "nil"
        if r, e := f(i); e != nil {
            fmt.Println("f failed:", e)
        } else {
            fmt.Println("f worked:", r)
        }
    }


	// range 5 => generate index from 0 to 4
    for i := range 5 {

		// err := makeTea(i); => call function makeTea dengan return err
		// err != nil => Check if err value is not nill
        if err := makeTea(i); err != nil {

			// Valaidate error with function errors.Is
			// errors.Is(err, ErrOutOfTea) => Check if err is equal to ErrOutOfTea
			// errors.Is(err, ErrPower) => Check if err is equal to ErrPower
            if errors.Is(err, ErrOutOfTea) {
                fmt.Println("We should buy new tea!")
            } else if errors.Is(err, ErrPower) {
                fmt.Println("Now it is dark.")
            } else {
                fmt.Printf("unknown error: %s\n", err)
            }
            continue
        }

        fmt.Println("Tea is ready!")
    }


	
	// Custom Errors
	fmt.Println("\nCustom Errors:")
	_, err := fcustom(42)
    var ae *argError
    if errors.As(err, &ae) {
        fmt.Println(ae.arg)
        fmt.Println(ae.message)
    } else {
        fmt.Println("err doesn't match argError")
    }
}

// Note :
// A "nil" value in the error position indicates that there was no error.