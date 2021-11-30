package user

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type User struct {
	Name   string
	Passwd string
	Post   []string
}

func GetUserByName(name string) (u User, err error) {
	file, err := os.Open(fmt.Sprintf("./files/users/%s.txt", name))
	if err != nil {
		err = errors.New("user not found")
		return
	}
	defer file.Close()

	u.Name = name

	return
}

func AddUserToFile(u User) (err error) {
	file, err := os.Create(fmt.Sprintf("./files/users/%s.txt", u.Name))
	if err != nil {
		err = errors.New("error to add user")
		return
	}
	defer file.Close()

	_, err = file.WriteString(u.Passwd)
	if err != nil {
		err = errors.New("error to write password")
		return
	}

	return
}

func LogIn(name, passwd string) (u User, err error) {
	file, err := os.Open(fmt.Sprintf("./files/users/%s.txt", name))
	if err != nil {
		err = errors.New("user not found")
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		err = errors.New("error to read password")
		return
	}

	if string(content) != passwd {
		err = errors.New("incorrect password")
		return
	}

	u.Name = name
	u.Passwd = passwd

	return
}

func (u User) AddPostToFile(post string) (err error) {
	file, err := os.OpenFile(fmt.Sprintf("./files/posts/%s.txt", u.Name), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = file.WriteString(post + "\n")
	if err != nil {
		return
	}

	return
}

func (u User) EditPost(postIndex int, newPost string) (err error) {
	if postIndex > len(u.Post)-1 && postIndex < 0 {
		err = errors.New("number out of range")
		return
	}

	file, err := os.Create(fmt.Sprintf("./files/posts/%s.txt", u.Name))
	if err != nil {
		return
	}
	defer file.Close()

	u.Post[postIndex] = newPost

	err = WritePostsToFile(u.Post, file)
	if err != nil {
		return
	}

	return
}

func (u User) DeletePost(postIndex int) (err error) {
	if postIndex > len(u.Post)-1 && postIndex < 0 {
		err = errors.New("number out of range")
		return
	}

	file, err := os.Create(fmt.Sprintf("./files/posts/%s.txt", u.Name))
	if err != nil {
		return
	}
	defer file.Close()

	u.Post = append(u.Post[:postIndex], u.Post[postIndex+1:]...)

	err = WritePostsToFile(u.Post, file)
	if err != nil {
		return
	}

	return
}

func (u *User) SyncPosts() (post []string, err error) {
	file, err := os.Open(fmt.Sprintf("./files/posts/%s.txt", u.Name))
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		post = append(post, scanner.Text())
	}

	u.Post = post

	return
}

func WritePostsToFile(post []string, file *os.File) (err error) {
	for _, v := range post {
		_, err = file.WriteString(v + "\n")
		if err != nil {
			return
		}
	}

	return
}
