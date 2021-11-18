package main

import (
	"fmt"
	"log"

	"github.com/sandjuarezg/practice-user-memory/functionality"
	"github.com/sandjuarezg/practice-user-memory/user"
)

func main() {
	functionality.CleanConsole()

	var opc int
	var exit bool
	var users []user.User

	for !exit {

		fmt.Println("- Registro de Usuarios -")
		fmt.Println("0. Exit")
		fmt.Println("1. Log in")
		fmt.Println("2. Sign up")
		fmt.Scanln(&opc)

		functionality.CleanConsole()

		switch opc {
		case 0:

			fmt.Println(". . . .  B Y E  . . . .")
			exit = true

			functionality.CleanConsole()

		case 1:

			n, err := user.LogIn(users)
			if err != nil {
				log.Println(err)
				functionality.CleanConsole()
				continue
			}

			var back bool

			for !back {

				functionality.CleanConsole()

				fmt.Printf("- Welcome %s -\n", users[n].User)
				fmt.Println("0. Sign off")
				fmt.Println("1. Add post")
				fmt.Println("2. Edit post")
				fmt.Println("3. Delete post")
				fmt.Println("4. Show your posts")
				fmt.Println("5. Show user's posts")
				fmt.Scanln(&opc)

				functionality.CleanConsole()

				switch opc {
				case 0:

					fmt.Println(". . . .  B Y E  . . . .")
					back = true

					functionality.CleanConsole()

				case 1:

					post, err := user.AddPost()
					if err != nil {
						log.Println(err)
						functionality.CleanConsole()
						continue
					}

					users[n].Post = append(users[n].Post, string(post))
					fmt.Println()
					fmt.Println("Post added successfully")

				case 2:

					fmt.Println("- Enter num of post to edit -")
					user.ShowAllPosts(users[n])

					err = user.EditPost(users, n)
					if err != nil {
						log.Println(err)
						functionality.CleanConsole()
						continue
					}

					fmt.Println()
					fmt.Println("Post edited successfully")

				case 3:

					fmt.Println("- Enter num of post to delete -")
					user.ShowAllPosts(users[n])

					err = user.DeletePost(users, n)
					if err != nil {
						log.Println(err)
						functionality.CleanConsole()
						continue
					}

					fmt.Println()
					fmt.Println("Post deleted successfully")

				case 4:

					fmt.Println("- All your post -")
					user.ShowAllPosts(users[n])

					fmt.Println()
					fmt.Println("Press ENTER to continue")
					fmt.Scanln()

				case 5:

					err := user.ShowUserPost(users)
					if err != nil {
						log.Println(err)
						functionality.CleanConsole()
						continue
					}

					fmt.Println()
					fmt.Println("Press ENTER to continue")
					fmt.Scanln()

				default:

					fmt.Println("Option not valid")
					functionality.CleanConsole()

				}
			}

		case 2:

			u, err := user.AddUser()
			if err != nil {
				log.Println("User couldn't be added", err)
				functionality.CleanConsole()
				continue
			}

			users = append(users, u)
			fmt.Println("User added successfully")

			functionality.CleanConsole()

		default:

			fmt.Println("Option not valid")
			functionality.CleanConsole()

		}

	}

}
