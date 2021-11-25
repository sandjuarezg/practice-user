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

func GetUser(name string) (u User, err error) {
	file, err := os.Open("./files/users/users.txt")
	if err != nil {
		return
	}
	defer file.Close()

	u, err = SearchUserFromFile(name, file)
	if err != nil {
		return
	}

	if u.Name == "" {
		err = errors.New("user not found")
		return
	}

	return
}

func AddUser(name, passwd string) (err error) {
	if name == "" {
		err = errors.New("user not found")
		return
	}

	file, err := os.OpenFile("./files/users/users.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	err = WriteUserPasswdToFile(name, passwd, file)
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

	u, err = SearchUserPassFromFile(name, passwd, file)
	if err != nil {
		return
	}

	return
}

func (u User) AddPost(post string) (err error) {
	if u.Name == "" {
		err = errors.New("user not found")
		return
	}

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

func (u User) EditPost(postIndex int, newPost string) (err error) {
	if u.Name == "" {
		err = errors.New("user not found")
		return
	}

	if postIndex > len(u.Post)-1 {
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

func (u User) DeletePost(postIndex int) (err error) {
	if u.Name == "" {
		err = errors.New("user not found")
		return
	}

	if postIndex > len(u.Post)-1 {
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

func (u *User) GetPosts() (post []string, err error) {
	if u.Name == "" {
		err = errors.New("user not found")
		return
	}

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

func WriteUserPasswdToFile(name, passwd string, file *os.File) (err error) {
	_, err = file.WriteString("u:" + name + "\n")
	if err != nil {
		return
	}

	_, err = file.WriteString("p:" + passwd + "\n\n")
	if err != nil {
		return
	}

	return
}

func SearchUserFromFile(name string, file *os.File) (u User, err error) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "u:") {

			if name == strings.TrimPrefix(line, "u:") {
				u.Name = name
				break
			}

		}

	}

	if u.Name == "" {
		err = errors.New("user not found")
		return
	}

	return
}

func SearchUserPassFromFile(name, passwd string, file *os.File) (u User, err error) {
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
