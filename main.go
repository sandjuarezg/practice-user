package main

import (
	"fmt"
	"log"

	"github.com/sandjuarezg/practice-user/functionality"
	"github.com/sandjuarezg/practice-user/user"
)

func main() {
	var (
		opc  int
		exit bool
	)

	err := functionality.PrepareFilePaths()
	if err != nil {
		log.Fatal(err)
	}

	for !exit {
		err := functionality.CleanConsole()
		if err != nil {
			log.Println(err)
			return
		}

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

			err := functionality.CleanConsole()
			if err != nil {
				log.Println(err)
				return
			}

		case 1:

			name, passwd, err := functionality.ScanNamePasswd()
			if err != nil {
				log.Println(err)
				continue
			}

			u, err := user.LogIn(name, passwd)
			if err != nil {
				log.Println(err)
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

				case 1:

					post, err := functionality.ScanPostText()
					if err != nil {
						log.Println(err)
						continue
					}

					err = u.AddPostToFile(post)
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println()
					fmt.Println("Post added successfully")

				case 2:

					var i int

					fmt.Println("- Enter num of post -")

					posts, err := u.SyncPosts()
					if err != nil {
						return
					}

					err = functionality.PrintUserPost(posts)
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Scanln(&i)
					i--

					fmt.Println()
					aux, err := functionality.ScanPostText()
					if err != nil {
						log.Println(err)
						continue
					}

					err = u.EditPost(i, aux)
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println()
					fmt.Println("Post edited successfully")

				case 3:

					var i int

					fmt.Println("- Enter num of post -")

					posts, err := u.SyncPosts()
					if err != nil {
						return
					}

					err = functionality.PrintUserPost(posts)
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Scanln(&i)
					i--

					err = u.DeletePost(i)
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println()
					fmt.Println("Post deleted successfully")

				case 4:

					fmt.Println("- All your post -")

					posts, err := u.SyncPosts()
					if err != nil {
						return
					}

					err = functionality.PrintUserPost(posts)
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println()
					fmt.Println("Press ENTER to continue")
					fmt.Scanln()

				case 5:

					name, err := functionality.ScanName()
					if err != nil {
						log.Println(err)
						continue
					}

					u, err := user.GetUserByName(name)
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println()
					fmt.Printf("- %s's posts -\n", u.Name)

					posts, err := u.SyncPosts()
					if err != nil {
						return
					}

					err = functionality.PrintUserPost(posts)
					if err != nil {
						log.Println(err)
						continue
					}

					fmt.Println()
					fmt.Println("Press ENTER to continue")
					fmt.Scanln()

				default:

					fmt.Println("Option not valid")

				}
			}

		case 2:
			name, passwd, err := functionality.ScanNamePasswd()
			if err != nil {
				log.Println(err)
				continue
			}

			u := user.User{Name: name, Passwd: passwd}

			err = user.AddUserToFile(u)
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println()
			fmt.Println("User added successfully")

		default:

			fmt.Println("Option not valid")

		}

	}

}
