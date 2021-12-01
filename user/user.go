package user

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type User struct {
	Name   string
	Passwd string
}

func PrepareUserPath(u User) (err error) {
	err = os.MkdirAll(fmt.Sprintf("./files/%s", u.Name), 0700)
	if err != nil {
		err = errors.New("error to create user dir")
		return
	}

	return
}

func PrepareUserPostsPath(u User) (err error) {
	err = os.MkdirAll(fmt.Sprintf("./files/%s/posts", u.Name), 0700)
	if err != nil {
		fmt.Println(err)
		err = errors.New("error to create user's posts dir")
		return
	}

	return
}

func AddUserToFile(u User) (err error) {
	file, err := os.Create(fmt.Sprintf("./files/%s/%s.txt", u.Name, u.Name))
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
	file, err := os.Open(fmt.Sprintf("./files/%s/%s.txt", name, name))
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

func ShowPostByName(name string) (err error) {
	files, err := os.ReadDir(fmt.Sprintf("./files/%s/posts", name))
	if err != nil {
		return
	}

	for _, f := range files {
		fmt.Printf("Key: %s\n", strings.TrimSuffix(f.Name(), ".txt"))

		data, err := os.ReadFile(fmt.Sprintf("./files/%s/posts/%s", name, f.Name()))
		if err != nil {
			break
		}

		if bytes.Equal(data[len(data)-1:], []byte("\n")) {
			data = data[:len(data)-1]
		}

		fmt.Printf("Content: %s\n", data)
		fmt.Println()
	}

	return
}

func (u User) AddPostToFile(post string) (err error) {
	files, err := os.ReadDir(fmt.Sprintf("./files/%s/posts", u.Name))
	if err != nil {
		return
	}

	file, err := os.Create(fmt.Sprintf("./files/%s/posts/%d.txt", u.Name, len(files)+1))
	if err != nil {
		return
	}
	defer file.Close()

	_, err = file.WriteString(post)
	if err != nil {
		return
	}

	return
}

func (u User) EditPost(postIndex int, newPost string) (err error) {
	files, err := os.ReadDir(fmt.Sprintf("./files/%s/posts", u.Name))
	if err != nil {
		return
	}

	if postIndex > len(files) || postIndex < 1 {
		err = errors.New("number out of range")
		return
	}

	file, err := os.Create(fmt.Sprintf("./files/%s/posts/%d.txt", u.Name, postIndex))
	if err != nil {
		return
	}
	defer file.Close()

	_, err = file.WriteString(newPost)
	if err != nil {
		return
	}

	return
}

func (u User) DeletePost(postIndex int) (err error) {
	files, err := os.ReadDir(fmt.Sprintf("./files/%s/posts", u.Name))
	if err != nil {
		return
	}

	if postIndex > len(files) || postIndex < 1 {
		err = errors.New("number out of range")
		return
	}

	err = os.Remove(fmt.Sprintf("./files/%s/posts/%d.txt", u.Name, postIndex))
	if err != nil {
		return
	}

	return

}
