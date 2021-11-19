package user

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name string
	Pass string
	Post []string
}

func CreateUser() (u User, err error) {
	fmt.Println("Enter user name")
	aux, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		log.Println("Error to add name", err)
		return
	}
	u.Name = string(aux)

	fmt.Println()

	fmt.Println("Enter password")
	aux, _, err = bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		log.Println("Error to add password", err)
		return
	}
	u.Pass = string(aux)

	return
}

func LogIn(users []User) (n int, err error) {
	var aux User
	n = -1

	fmt.Println("Enter user name")
	name, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		log.Println("Error to find name", err)
		return
	}
	aux.Name = string(name)

	fmt.Println()

	fmt.Println("Enter password")
	pass, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		log.Println("Error to find password", err)
		return
	}
	aux.Pass = string(pass)

	for i := range users {
		if users[i].Name == aux.Name && users[i].Pass == aux.Pass {
			n = i
			break
		}
	}

	return
}

func CreatePost() (post []byte, err error) {

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
		if u.Name == string(aux) {
			fmt.Printf("- %s's posts -\n", u.Name)

			for i, v := range u.Post {
				fmt.Printf("%d. %s\n", i+1, v)
			}

			return
		}
	}

	err = errors.New("User not found")

	return
}
