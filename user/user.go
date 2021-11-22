package user

import (
	"errors"
)

type User struct {
	Name string
	Pass string
	Post []string
}

type users struct {
	users []User
}

var us users

func GetUser(name string) (u *User, err error) {
	for i := range us.users {
		if us.users[i].Name == name {
			u = &us.users[i]
			break
		}
	}

	if u.isEmpty() {
		err = errors.New("user not found")
		return
	}

	return
}

func (u *User) isEmpty() bool {
	return u == nil
}

func (u *User) AddUser() (err error) {
	if u.isEmpty() {
		err = errors.New("user is empty")
		return
	}

	us.users = append(us.users, *u)

	return
}

func (aux User) LogIn() (u *User, err error) {
	for i := range us.users {
		if us.users[i].Name == aux.Name && us.users[i].Pass == aux.Pass {
			u = &us.users[i]
			break
		}
	}

	if u.isEmpty() {
		err = errors.New("user is empty")
		return
	}

	return
}

func (u *User) AddPost(post string) (err error) {
	if u.isEmpty() {
		err = errors.New("user is empty")
		return
	}

	u.Post = append(u.Post, post)

	return
}

func (u *User) EditPost(postIndex int, newPost string) (err error) {
	if u.isEmpty() {
		err = errors.New("user is empty")
		return
	}

	if postIndex > len(u.Post)-1 {
		err = errors.New("number out of range")
		return
	}

	u.Post[postIndex] = newPost

	return
}

func (u *User) DeletePost(postIndex int) (err error) {
	if u.isEmpty() {
		err = errors.New("user is empty")
		return
	}

	if postIndex > len(u.Post)-1 {
		err = errors.New("number out of range")
		return
	}

	u.Post = append(u.Post[:postIndex], u.Post[postIndex+1:]...)

	return
}

func (u *User) GetPosts() (post []string, err error) {
	if u.isEmpty() {
		err = errors.New("user is empty")
		return
	}

	post = u.Post

	return
}
