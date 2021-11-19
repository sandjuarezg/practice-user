package main

import (
	"fmt"
	"log"

	"github.com/sandjuarezg/practice-user-memory/functionality"
	"github.com/sandjuarezg/practice-user-memory/user"
)

func main() {
	var (
		opc   int
		exit  bool
		users []user.User
	)

	err := functionality.CleanConsole()
	if err != nil {
		log.Println(err)
		return
	}

	for !exit {

		fmt.Println("- Registro de Usuarios -")
		fmt.Println("0. Exit")
		fmt.Println("1. Log in")
		fmt.Println("2. Sign up")
		fmt.Scanln(&opc)

		err = functionality.CleanConsole()
		if err != nil {
			log.Println(err)
			continue
		}

		switch opc {
		case 0:

			fmt.Println(". . . .  B Y E  . . . .")
			exit = true

			err = functionality.CleanConsole()
			if err != nil {
				log.Println(err)
				continue
			}

		case 1:

			n, err := user.LogIn(users)
			if n < 0 {
				log.Println("User not found")

				err = functionality.CleanConsole()
				if err != nil {
					log.Println(err)
				}
				continue
			}

			if err != nil {
				log.Println(err)

				err = functionality.CleanConsole()
				if err != nil {
					log.Println(err)
				}
				continue
			}

			var back bool

			for !back {

				err = functionality.CleanConsole()
				if err != nil {
					log.Println(err)
					continue
				}

				fmt.Printf("- Welcome %s -\n", users[n].Name)
				fmt.Println("0. Sign off")
				fmt.Println("1. Add post")
				fmt.Println("2. Edit post")
				fmt.Println("3. Delete post")
				fmt.Println("4. Show your posts")
				fmt.Println("5. Show user's posts")
				fmt.Scanln(&opc)

				err = functionality.CleanConsole()
				if err != nil {
					log.Println(err)
					continue
				}

				switch opc {
				case 0:

					fmt.Println(". . . .  B Y E  . . . .")
					back = true

					err = functionality.CleanConsole()
					if err != nil {
						log.Println(err)
						continue
					}

				case 1:

					post, err := user.CreatePost()
					if err != nil {
						log.Println(err)

						err = functionality.CleanConsole()
						if err != nil {
							log.Println(err)
						}

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

						err = functionality.CleanConsole()
						if err != nil {
							log.Println(err)
						}

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

						err = functionality.CleanConsole()
						if err != nil {
							log.Println(err)
						}

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

						err = functionality.CleanConsole()
						if err != nil {
							log.Println(err)
						}

						continue
					}

					fmt.Println()
					fmt.Println("Press ENTER to continue")
					fmt.Scanln()

				default:

					fmt.Println("Option not valid")

					err = functionality.CleanConsole()
					if err != nil {
						log.Println(err)
						continue
					}

				}
			}

		case 2:

			u, err := user.CreateUser()
			if err != nil {
				log.Println("User couldn't be added", err)

				err = functionality.CleanConsole()
				if err != nil {
					log.Println(err)
				}

				continue
			}

			users = append(users, u)
			fmt.Println("User added successfully")

			err = functionality.CleanConsole()
			if err != nil {
				log.Println(err)
				continue
			}

		default:

			fmt.Println("Option not valid")

			err = functionality.CleanConsole()
			if err != nil {
				log.Println(err)
				continue
			}

		}

	}

}
