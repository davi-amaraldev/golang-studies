package main

import "fmt"

func main() {
	var variable1 string = "Variável 1"
	variable2 := "Variável 2" // Inferência de Tipo
	var (
		variable3 string = "a"
		variable4 string = "b"
	)

	variable5, variable6 := "Váriavel 5", "Variável 6"
	
	const constant1 = "Constante 1"

	variable5, variable6 = variable6, variable5

	fmt.Println(variable1)
	fmt.Println(variable2)
	fmt.Println(variable3, variable4)
	fmt.Println(variable5, variable6)
	fmt.Println(constant1)
}