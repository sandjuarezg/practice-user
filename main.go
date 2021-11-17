package main

import (
	"fmt"
)

func main() {
	var opc int
	var exit bool

	for !exit {

		fmt.Println("- Registro de Usuarios -")
		fmt.Println("0. Exit")
		fmt.Println("1. Log in")
		fmt.Println("2. Sign up")
		fmt.Scan(&opc)

		switch opc {
		case 0:

			fmt.Println(". . . B Y E . . .")
			exit = true

		case 1:

		case 2:

		default:

			fmt.Println("Option not valid")

		}

	}

}
