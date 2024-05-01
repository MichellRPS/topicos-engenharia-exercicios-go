package main

import "fmt"

func main() {
	a := [10]int{43, 67, 125, 96, 78, 104, 45, 99, 81, 55}
	// creates an empty map with key and value of type int
	// m := map[int]int{}
	m := make(map[int]int)

	// loops through the array
	// for i, v := range a {
	for i := 0; i < len(a); i++ {
		// assigns the current index and and value of the array to a key and value in the map, respectively
		m[i] = a[i]
	}

	fmt.Println(m)

	/*
		references:
		https://www.w3schools.com/go/go_maps.php
		https://golangbyexample.com/length-of-string-golang/
	*/
}
