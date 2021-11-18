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
	_, err = fmt.Scan(&u.User)
	if err != nil {
		log.Println("Error to add name", err)
		return
	}

	fmt.Println("Enter password")
	_, err = fmt.Scan(&u.Pass)
	if err != nil {
		log.Println("Error to add password", err)
		return
	}

	return
}

func LogIn(users []User) (n int, err error) {
	var aux User

	fmt.Println("Enter user name")
	_, err = fmt.Scan(&aux.User)
	if err != nil {
		log.Println("Error to find name", err)
		return
	}

	fmt.Println("Enter password")
	_, err = fmt.Scan(&aux.Pass)
	if err != nil {
		log.Println("Error to find password", err)
		return
	}

	for n = range users {
		if users[n].User == aux.User && users[n].Pass == aux.Pass {
			return
		}
	}

	err = errors.New("User not found")

	return
}

func AddPost() (post []byte, err error) {

	r := bufio.NewReader(os.Stdin)
	r.ReadString('\n')

	fmt.Print("Enter text: ")
	post, _, err = r.ReadLine()
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

	if i > len(users[0].Post)-1 {
		err = errors.New("number out of range")
		return
	}

	r := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	aux, _, err := r.ReadLine()
	if err != nil {
		log.Println("Error to edit post", err)
		return
	}

	users[0].Post[i] = string(aux)

	return
}

func DeletePost(users []User, n int) (err error) {
	var i int
	fmt.Scanln(&i)
	i--

	if i > len(users[0].Post)-1 {
		err = errors.New("number out of range")
		return
	}

	users[n].Post = append(users[0].Post[:i], users[0].Post[i+1:]...)

	return
}

func ShowAllPosts(u User) {
	for i, v := range u.Post {
		fmt.Printf("%d. %s\n", i+1, v)
	}
}

func ShowUserPost(users []User) (err error) {
	var aux string

	fmt.Println("Enter user name")
	_, err = fmt.Scan(&aux)
	if err != nil {
		log.Println("Error to find name", err)
		return
	}

	for _, u := range users {
		if u.User == aux {
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
