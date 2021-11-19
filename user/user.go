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

func LogIn(users []User) (u *User, err error) {
	var aux User

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
			u = &users[i]
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

func EditPost(u *User) (err error) {
	var i int
	fmt.Scanln(&i)
	i--

	if i > len(u.Post)-1 {
		err = errors.New("number out of range")
		return
	}

	fmt.Print("Enter text: ")
	aux, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		log.Println("Error to edit post", err)
		return
	}

	u.Post[i] = string(aux)

	return
}

func DeletePost(u *User) (err error) {
	var i int
	fmt.Scanln(&i)
	i--

	if i > len(u.Post)-1 {
		err = errors.New("number out of range")
		return
	}

	u.Post = append(u.Post[:i], u.Post[i+1:]...)

	return
}

func ShowAllPosts(u *User) {
	for i, v := range u.Post {
		fmt.Printf("%d. %s\n", i+1, v)
	}
}

func GetUser(users []User) (u *User, err error) {
	fmt.Println("Enter user name")
	aux, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		log.Println("Error to find name", err)
		return
	}

	for i := range users {
		if users[i].Name == string(aux) {
			u = &users[i]
			break
		}
	}

	return
}
