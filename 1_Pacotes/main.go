package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main(){
	fmt.Println("eai dog!")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("davicd@regm.com")
	fmt.Println(erro)
}
