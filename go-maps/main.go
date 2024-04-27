package main

import "fmt"

func main() {
	a := [10]int{43, 67, 125, 96, 78, 104, 45, 99, 81, 55}
	// cria um mapa vazio com chave e valor do tipo int
	m := map[int]int{}

	// percorre a lista de inteiros
	for i, v := range a {
		// coloca o índice e o valor atuais da lista como chave e valor no mapa, respectivamente
		m[i] = v
	}

	fmt.Println(m)

	/*
		references:
		https://www.w3schools.com/go/go_maps.php
		https://golangbyexample.com/length-of-string-golang/
	*/
}
