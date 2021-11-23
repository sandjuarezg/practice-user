package user

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name string
	Pass string
	Post []string
}

type file struct {
	f    *os.File
	path string
}

var (
	usersFile file = file{path: "./user/files/users/users.txt"}
	postsFile file
)

func GetUser(name string) (u *User, err error) {
	usersFile.f, err = os.Open(usersFile.path)
	if err != nil {
		return
	}
	defer usersFile.f.Close()

	u, err = usersFile.SearchUser(name)
	if err != nil {
		return
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

	usersFile.f, err = os.OpenFile(usersFile.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer usersFile.f.Close()

	err = usersFile.WriteUserPass(u)
	if err != nil {
		return
	}

	return
}

func (aux User) LogIn() (u *User, err error) {
	usersFile.f, err = os.Open(usersFile.path)
	if err != nil {
		return
	}
	defer usersFile.f.Close()

	u, err = usersFile.SearchUserPass(aux)
	if err != nil {
		return
	}

	if u.isEmpty() {
		err = errors.New("user not found")
		return
	}

	return
}

func (u *User) AddPost(post string) (err error) {
	if u.isEmpty() {
		err = errors.New("user is empty")
		return
	}

	postsFile.path = fmt.Sprintf("./user/files/posts/%s.txt", u.Name)

	postsFile.f, err = os.OpenFile(postsFile.path, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer usersFile.f.Close()

	_, err = postsFile.f.WriteString(post + "\n")
	if err != nil {
		return
	}

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

	postsFile.path = fmt.Sprintf("./user/files/posts/%s.txt", u.Name)

	postsFile.f, err = os.Create(postsFile.path)
	if err != nil {
		return
	}
	defer usersFile.f.Close()

	u.Post[postIndex] = newPost

	err = postsFile.WritePosts(u.Post)
	if err != nil {
		return
	}

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

	postsFile.path = fmt.Sprintf("./user/files/posts/%s.txt", u.Name)

	postsFile.f, err = os.Create(postsFile.path)
	if err != nil {
		return
	}
	defer usersFile.f.Close()

	u.Post = append(u.Post[:postIndex], u.Post[postIndex+1:]...)

	err = postsFile.WritePosts(u.Post)
	if err != nil {
		return
	}

	return
}

func (u *User) GetPosts() (post []string, err error) {
	if u.isEmpty() {
		err = errors.New("user is empty")
		return
	}

	postsFile.path = fmt.Sprintf("./user/files/posts/%s.txt", u.Name)

	postsFile.f, err = os.Open(postsFile.path)
	if err != nil {
		return
	}
	defer usersFile.f.Close()

	scanner := bufio.NewScanner(postsFile.f)

	for scanner.Scan() {
		post = append(post, scanner.Text())
	}

	u.Post = post

	return
}

func (file *file) WriteUserPass(u *User) (err error) {
	_, err = file.f.WriteString("u:" + u.Name + "\n")
	if err != nil {
		return
	}

	_, err = usersFile.f.WriteString("p:" + u.Pass + "\n\n")
	if err != nil {
		return
	}

	return
}

func (file *file) SearchUser(name string) (u *User, err error) {
	var (
		aux     User
		scanner *bufio.Scanner = bufio.NewScanner(file.f)
	)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "u:") {

			if name == strings.TrimPrefix(line, "u:") {

				aux.Name = name
				u = &aux
				break

			}

		}

	}

	if u.isEmpty() {
		err = errors.New("user not found")
		return
	}

	return
}

func (file *file) SearchUserPass(aux User) (u *User, err error) {
	scanner := bufio.NewScanner(file.f)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "u:") {

			if aux.Name == strings.TrimPrefix(line, "u:") {

				if scanner.Scan() {
					line = scanner.Text()

					if strings.HasPrefix(line, "p:") {

						if aux.Pass == strings.TrimPrefix(line, "p:") {
							u = &aux
							break
						}

					}

				}
			}

		}

	}

	if u.isEmpty() {
		err = errors.New("user not found")
		return
	}

	return
}

func (file *file) WritePosts(post []string) (err error) {
	for _, v := range post {
		_, err = file.f.WriteString(v + "\n")
		if err != nil {
			return
		}
	}

	return
}
