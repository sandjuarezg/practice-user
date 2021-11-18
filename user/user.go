package user

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type User struct {
	User string
	Pass string
	Post []string
}

func AddUser() (u User, err error) {
	fmt.Println("Enter user name")
	aux, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		log.Println("Error to add name", err)
		return
	}
	u.User = string(aux)

	fmt.Println()

	fmt.Println("Enter password")
	aux, _, err = bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		log.Println("Error to add password", err)
		return
	}
	u.Pass = string(aux)

	fmt.Println()

	return
}

func LogIn(users []User) (n int, err error) {
	var aux User

	fmt.Println("Enter user name")
	name, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		log.Println("Error to find name", err)
		return
	}
	aux.User = string(name)

	fmt.Println()

	fmt.Println("Enter password")
	pass, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		log.Println("Error to find password", err)
		return
	}
	aux.Pass = string(pass)

	for n = range users {
		if users[n].User == aux.User && users[n].Pass == aux.Pass {
			return
		}
	}

	fmt.Println()
	err = errors.New("User not found")

	return
}

func AddPost() (post []byte, err error) {

	fmt.Print("Enter text: ")
	post, _, err = bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		log.Println("Error to add post", err)
		return
	}

	return
}

func EditPost(users []User, n int) (err error) {
	var i int
	fmt.Scanln(&i)
	i--

	if i > len(users[n].Post)-1 {
		err = errors.New("number out of range")
		return
	}

	fmt.Print("Enter text: ")
	aux, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		log.Println("Error to edit post", err)
		return
	}

	users[n].Post[i] = string(aux)

	return
}

func DeletePost(users []User, n int) (err error) {
	var i int
	fmt.Scanln(&i)
	i--

	if i > len(users[n].Post)-1 {
		err = errors.New("number out of range")
		return
	}

	users[n].Post = append(users[n].Post[:i], users[n].Post[i+1:]...)

	return
}

func ShowAllPosts(u User) {
	for i, v := range u.Post {
		fmt.Printf("%d. %s\n", i+1, v)
	}
}

func ShowUserPost(users []User) (err error) {
	fmt.Println("Enter user name")
	aux, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		log.Println("Error to find name", err)
		return
	}

	for _, u := range users {
		if u.User == string(aux) {
			fmt.Printf("- %s's posts -\n", u.User)

			for i, v := range u.Post {
				fmt.Printf("%d. %s\n", i+1, v)
			}

			return
		}
	}

	err = errors.New("User not found")

	return
}
