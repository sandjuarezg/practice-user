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

func EditPost() {

}

func DeletePost() {

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
