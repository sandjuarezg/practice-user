package user

import (
	"fmt"
	"log"
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

func LogIn() {
	fmt.Println("Enter user name")
	fmt.Println("Enter password")
}

func AddPost() {

}

func EditPost() {

}

func DeletePost() {

}

func ShowAllPosts() {

}

func ShowUserPost() {

}
