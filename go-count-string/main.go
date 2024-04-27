package main

import (
	"errors"
	"fmt"
)

func countString(s string) (int, error) {
	// checks if the string is empty
	if (s == "") {
		// returns an error
		return 0, errors.New("empty string")
	}

	// converts the string to a rune array containing all the characters of the string
	r := []rune(s)
	// gets the length of the rune array
	l := len(r)

	// returns the length of the string
	return l, nil

	/*
	references:
	https://www.w3schools.com/go/go_function_parameters.php
	https://golangbyexample.com/length-of-string-golang/
	https://go.dev/doc/tutorial/handle-errors
	*/
}

func main() {
	// l, e := countString("")
	l, e := countString("test")

	if (e != nil) {
		fmt.Println(e)
		return
	}

	fmt.Println(l)
}
