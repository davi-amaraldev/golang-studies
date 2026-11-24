package auxiliar

import "fmt"

// A função escrever2 escreve outro texto, mas é privada.
func escrever2() {
	fmt.Println("Escrevendo segundo texto")
	escrever2()
}