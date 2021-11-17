package user

import "fmt"

type User struct {
	User string
	Pass string
	Post []string
}

func AddUser() (u User, err error) {
	fmt.Println("Enter user name")
	_, err = fmt.Scan(&u.User)

	fmt.Println("Enter password")
	_, err = fmt.Scan(&u.Pass)

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
