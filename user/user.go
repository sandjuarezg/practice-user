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

func GetUserByName(name string) (u User, err error) {
	file, err := os.Open("./files/users/users.txt")
	if err != nil {
		return
	}
	defer file.Close()

	ban, err := ExistUserFromFile(name, file)
	if err != nil {
		return
	}

	if !ban {
		err = errors.New("user not found")
		return
	}

	u.Name = name

	return
}

func AddUserToFile(u User) (err error) {
	file, err := os.OpenFile("./files/users/users.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	err = WriteUserPasswdToFile(u, file)
	if err != nil {
		return
	}

	return
}

func LogIn(name, passwd string) (u User, err error) {
	file, err := os.Open("./files/users/users.txt")
	if err != nil {
		return
	}
	defer file.Close()

	u, err = SearchUserPasswdFromFile(name, passwd, file)
	if err != nil {
		return
	}

	return
}

func (u User) AddPostToFile(post string) (err error) {
	file, err := os.OpenFile(fmt.Sprintf("./files/posts/%s.txt", u.Name), os.O_APPEND|os.O_CREATE, 0666)
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

func (u User) EditPostFromFile(postIndex int, newPost string) (err error) {
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

	err = WritePosts(u.Post, file)
	if err != nil {
		return
	}

	return
}

func (u User) DeletePostFromFile(postIndex int) (err error) {
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

	err = WritePosts(u.Post, file)
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

func WriteUserPasswdToFile(u User, file *os.File) (err error) {
	_, err = file.WriteString("u:" + u.Name + "\n")
	if err != nil {
		return
	}

	_, err = file.WriteString("p:" + u.Passwd + "\n\n")
	if err != nil {
		return
	}

	return
}

func ExistUserFromFile(name string, file *os.File) (ban bool, err error) {
	scanner := bufio.NewScanner(file)

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

func SearchUserPasswdFromFile(name, passwd string, file *os.File) (u User, err error) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "u:") {

			if name == strings.TrimPrefix(line, "u:") {

				if scanner.Scan() {
					line = scanner.Text()

					if strings.HasPrefix(line, "p:") {

						if passwd == strings.TrimPrefix(line, "p:") {
							u.Name = name
							u.Passwd = passwd
							break
						}

					}

				}
			}

		}

	}

	if u.Name == "" {
		err = errors.New("user not found")
		return
	}

	return
}

func WritePosts(post []string, file *os.File) (err error) {
	for _, v := range post {
		_, err = file.WriteString(v + "\n")
		if err != nil {
			return
		}
	}

	return
}
