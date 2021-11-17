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

			u, err := user.LogIn(users)
			if err != nil {
				log.Println(err)
				return
			}

			var back bool

			for !back {

				fmt.Printf("- Welcome %s -\n", u.User)
				fmt.Println("0. Sign off")
				fmt.Println("1. Add post")
				fmt.Println("2. Edit post")
				fmt.Println("3. Delete post")
				fmt.Println("4. Show your posts")
				fmt.Println("5. Show user's posts")
				fmt.Scanln(&opc)

				switch opc {
				case 0:

					fmt.Println(". . . B Y E . . .")
					back = true

				case 1:

					post, err := user.AddPost()
					if err != nil {
						log.Println(err)
						return
					}
					fmt.Println("Post added successfully")

					u.Post = append(u.Post, string(post))

				case 2:

					user.EditPost()

				case 3:

					user.DeletePost()

				case 4:

					user.ShowAllPosts()

				case 5:

					user.ShowUserPost()

				default:

					fmt.Println("Option not valid")

				}
			}

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
