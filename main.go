package main

import (
	"fmt"
	"log"

	"github.com/sandjuarezg/practice-user-memory/user"
)

func main() {
	var opc int
	var exit bool
	var users []user.User

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

			u, err := user.AddUser()
			if err != nil {
				log.Println("User couldn't be added", err)
				return
			}
			fmt.Println("User added successfully")

			users = append(users, u)

		default:

			fmt.Println("Option not valid")

		}

	}

}
