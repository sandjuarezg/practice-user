package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/sandjuarezg/practice-user/functionality"
	"github.com/sandjuarezg/practice-user/user"
)

func main() {
	var (
		opc  int
		exit bool
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

			var aux user.User

			fmt.Println("Enter user name")
			name, _, err := bufio.NewReader(os.Stdin).ReadLine()
			if err != nil {
				log.Println("Error to find name", err)
				continue
			}
			aux.Name = string(name)

			fmt.Println()

			fmt.Println("Enter password")
			pass, _, err := bufio.NewReader(os.Stdin).ReadLine()
			if err != nil {
				log.Println("Error to find password", err)
				continue
			}
			aux.Pass = string(pass)

			u, err := user.LogIn(aux)
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

				fmt.Printf("- Welcome %s -\n", u.Name)
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

					fmt.Print("Enter text: ")
					post, _, err := bufio.NewReader(os.Stdin).ReadLine()
					if err != nil {
						log.Println("Error to add post", err)
						continue
					}

					user.AddPost(u, string(post))

					fmt.Println()
					fmt.Println("Post added successfully")

				case 2:

					var i int

					fmt.Println("- Enter num of post to edit -")
					user.ShowPosts(u)
					fmt.Scanln(&i)
					i--
					if i > len(u.Post)-1 {
						log.Println("number out of range")
						continue
					}

					fmt.Println()
					fmt.Print("Enter text: ")
					aux, _, err := bufio.NewReader(os.Stdin).ReadLine()
					if err != nil {
						log.Println("Error to edit post", err)
						continue
					}

					user.EditPost(u, i, string(aux))
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

					var i int

					fmt.Println("- Enter num of post to delete -")
					user.ShowPosts(u)
					fmt.Scanln(&i)
					i--

					if i > len(u.Post)-1 {
						log.Panicln("number out of range")
						continue
					}

					user.DeletePost(u, i)

					fmt.Println()
					fmt.Println("Post deleted successfully")

				case 4:

					fmt.Println("- All your post -")
					user.ShowPosts(u)

					fmt.Println()
					fmt.Println("Press ENTER to continue")
					fmt.Scanln()

				case 5:

					fmt.Println("Enter user name")
					name, _, err := bufio.NewReader(os.Stdin).ReadLine()
					if err != nil {
						log.Println("Error to find name", err)
						continue
					}

					u, err := user.GetUser(string(name))
					if err != nil {
						log.Println(err)

						err = functionality.CleanConsole()
						if err != nil {
							log.Println(err)
						}

						continue
					}

					fmt.Println()
					fmt.Printf("- %s's posts -\n", u.Name)
					user.ShowPosts(u)

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
			var u user.User

			fmt.Println("Enter user name")
			aux, _, err := bufio.NewReader(os.Stdin).ReadLine()
			if err != nil {
				log.Println("Error to add name", err)
				continue
			}
			u.Name = string(aux)

			fmt.Println()

			fmt.Println("Enter password")
			aux, _, err = bufio.NewReader(os.Stdin).ReadLine()
			if err != nil {
				log.Println("Error to add password", err)
				continue
			}
			u.Pass = string(aux)

			user.AddUser(u)

			fmt.Println()
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
