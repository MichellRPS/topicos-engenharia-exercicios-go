package main

import "fmt"

func legalAge(age int) bool {
	return age >= 18
}

func main() {
	// isLegalAge := legalAge(17)
	isLegalAge := legalAge(19)

	if isLegalAge {
		fmt.Println("É maior")
		return
	}

	fmt.Println("É menor")
}
