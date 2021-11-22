package user

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
	Pass string
	Post []string
}

var users []User

func AddUser(u User) {
	users = append(users, u)
}

func LogIn(aux User) (u *User, err error) {
	for i := range users {
		if users[i].Name == aux.Name && users[i].Pass == aux.Pass {
			u = &users[i]
			break
		}
	}

	if u == nil {
		err = errors.New("user not found")
		return
	}

	return
}

func AddPost(u *User, post string) {
	u.Post = append(u.Post, post)
}

func EditPost(u *User, postIndex int, newPost string) {
	u.Post[postIndex] = newPost
}

func DeletePost(u *User, postIntex int) {
	u.Post = append(u.Post[:postIntex], u.Post[postIntex+1:]...)
}

func ShowPosts(u *User) {
	for i, v := range u.Post {
		fmt.Printf("%d. %s\n", i+1, v)
	}
}

func GetUser(name string) (u *User, err error) {
	for i := range users {
		if users[i].Name == name {
			u = &users[i]
			break
		}
	}

	if u == nil {
		err = errors.New("user not found")
		return
	}

	return
}
