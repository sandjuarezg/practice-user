package user

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name   string
	Passwd string
	Post   []string
}

type file struct {
	desc *os.File
}

func GetUserByName(name string) (u User, err error) {
	var f file
	f.desc, err = os.Open("./files/users/users.txt")
	if err != nil {
		return
	}
	defer f.desc.Close()

	ban := f.ExistUser(name)

	if !ban {
		err = errors.New("user not found")
		return
	}

	u.Name = name

	return
}

func AddUserToFile(u User) (err error) {
	var f file
	f.desc, err = os.OpenFile("./files/users/users.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return
	}
	defer f.desc.Close()

	err = f.WriteUserPasswd(u)
	if err != nil {
		return
	}

	return
}

func LogIn(name, passwd string) (u User, err error) {
	var f file
	f.desc, err = os.Open("./files/users/users.txt")
	if err != nil {
		return
	}
	defer f.desc.Close()

	ban := f.ExistUserPasswd(name, passwd)
	if !ban {
		err = errors.New("user not found")
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

	var f file
	f.desc, err = os.Create(fmt.Sprintf("./files/posts/%s.txt", u.Name))
	if err != nil {
		return
	}
	defer f.desc.Close()

	u.Post[postIndex] = newPost

	err = f.WritePosts(u.Post)
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

	var f file
	f.desc, err = os.Create(fmt.Sprintf("./files/posts/%s.txt", u.Name))
	if err != nil {
		return
	}
	defer f.desc.Close()

	u.Post = append(u.Post[:postIndex], u.Post[postIndex+1:]...)

	err = f.WritePosts(u.Post)
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

func (file file) WriteUserPasswd(u User) (err error) {
	_, err = file.desc.WriteString("u:" + u.Name + "\n")
	if err != nil {
		return
	}

	_, err = file.desc.WriteString("p:" + u.Passwd + "\n\n")
	if err != nil {
		return
	}

	return
}

func (file file) ExistUser(name string) (ban bool) {
	scanner := bufio.NewScanner(file.desc)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "u:") {

			if name == strings.TrimPrefix(line, "u:") {
				ban = true
				break
			}

		}

	}

	return
}

func (file file) ExistUserPasswd(name, passwd string) (ban bool) {
	scanner := bufio.NewScanner(file.desc)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "u:") {

			if name == strings.TrimPrefix(line, "u:") {

				if scanner.Scan() {
					line = scanner.Text()

					if strings.HasPrefix(line, "p:") {

						if passwd == strings.TrimPrefix(line, "p:") {
							ban = true
							break
						}

					}

				}
			}

		}

	}

	return
}

func (file file) WritePosts(post []string) (err error) {
	for _, v := range post {
		_, err = file.desc.WriteString(v + "\n")
		if err != nil {
			return
		}
	}

	return
}
