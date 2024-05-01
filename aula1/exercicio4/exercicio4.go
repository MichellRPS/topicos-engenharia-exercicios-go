package main

import (
	"errors"
	"fmt"
)

func divide(i1 int, i2 int) (int, error) {
	if i2 == 0 {
		return 0, errors.New("division by zero")
	}

	return (i1 / i2), nil
}

func main() {
	l1 := [10]int{43, 67, 125, 96, 78, 104, 45, 99, 81, 55}
	l2 := [10]int{11, 29, 52, 47, 30, 0, 18, 66, 2, 88}

	for i := 0; i < len(l1); i++ {
		q, e := divide(l1[i], l2[i])

		if e != nil {
			fmt.Println(e)
			// fmt.Printf("%d / %d \n", l1[i], l2[i])
			fmt.Println(l1[i], "/", l2[i], "=", q)
			continue
		}

		fmt.Println(l1[i], "/", l2[i], "=", q)
	}

	/*
		references:
		https://pkg.go.dev/fmt#Sprintf
	*/
}
