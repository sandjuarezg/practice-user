package user

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
)

type users struct {
	XMLName xml.Name `xml:"Users"`

	Users []User `xml:"User"`
}

type User struct {
	Name   string `xml:"Name"`
	Passwd string `xml:"Passwd"`
	Posts  []post `xml:"Posts"`
}

type post struct {
	Post string
}

func GetIndexUser(u User, us []User) (i int) {
	i = -1
	for n, v := range us {
		if v.Name == u.Name && v.Passwd == u.Passwd {
			i = n
			break
		}
	}

	return
}

func AddUserToFile(u User) (err error) {
	file, err := os.OpenFile("./users.xml", os.O_CREATE|os.O_RDONLY, 0600)
	if err != nil {
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return
	}

	var us users

	if !bytes.Equal(content, []byte("")) {
		err = xml.Unmarshal(content, &us)
		if err != nil {
			return
		}
	}

	us.Users = append(us.Users, u)

	b, err := xml.MarshalIndent(us, "", "\t")
	if err != nil {
		return
	}

	err = os.WriteFile(file.Name(), b, 0)
	if err != nil {
		return
	}

	return
}

func LogIn(name, passwd string) (u User, err error) {
	file, err := os.Open("./users.xml")
	if err != nil {
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return
	}

	var us users
	err = xml.Unmarshal(content, &us)
	if err != nil {
		return
	}

	n := GetIndexUser(User{Name: name, Passwd: passwd}, us.Users)
	if n < 0 {
		err = errors.New("user not found")
		return
	}

	u = us.Users[n]

	return
}

func ShowPostByName(name string) (err error) {
	content, err := os.ReadFile("./users.xml")
	if err != nil {
		return
	}

	var us users
	err = xml.Unmarshal(content, &us)
	if err != nil {
		return
	}

	for _, v := range us.Users {
		if v.Name == name {
			fmt.Printf("Post's %s:\n", name)
			for i, v := range v.Posts {
				fmt.Printf("Key: %d\n", i+1)
				fmt.Printf("Content: %s\n", v)
				fmt.Println()
			}
			break
		}
	}

	return
}

func (u User) AddPostToFile(postText string) (err error) {
	file, err := os.Open("./users.xml")
	if err != nil {
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return
	}

	var us users

	err = xml.Unmarshal(content, &us)
	if err != nil {
		return
	}

	n := GetIndexUser(u, us.Users)
	if n < 0 {
		err = errors.New("user not found")
		return
	}

	us.Users[n].Posts = append(us.Users[n].Posts, post{postText})

	b, err := xml.MarshalIndent(us, "", "\t")
	if err != nil {
		return
	}

	err = os.WriteFile(file.Name(), b, 0)
	if err != nil {
		return
	}

	return

}

func (u User) EditPost(postIndex int, newPost string) (err error) {
	file, err := os.Open("./users.xml")
	if err != nil {
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return
	}

	var us users

	err = xml.Unmarshal(content, &us)
	if err != nil {
		return
	}

	n := GetIndexUser(u, us.Users)
	if n < 0 {
		err = errors.New("user not found")
		return
	}

	postIndex--
	if postIndex > len(us.Users[n].Posts)-1 || postIndex < 0 {
		err = errors.New("number out of range")
		return
	}

	us.Users[n].Posts[postIndex].Post = newPost

	b, err := xml.MarshalIndent(us, "", "\t")
	if err != nil {
		return
	}

	err = os.WriteFile(file.Name(), b, 0)
	if err != nil {
		return
	}

	return
}

func (u User) DeletePost(postIndex int) (err error) {
	file, err := os.Open("./users.xml")
	if err != nil {
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return
	}

	var us users

	err = xml.Unmarshal(content, &us)
	if err != nil {
		return
	}

	n := GetIndexUser(u, us.Users)
	if n < 0 {
		err = errors.New("user not found")
		return
	}

	postIndex--
	if postIndex > len(us.Users[n].Posts)-1 || postIndex < 0 {
		err = errors.New("number out of range")
		return
	}

	us.Users[n].Posts = append(us.Users[n].Posts[:postIndex], us.Users[n].Posts[postIndex+1:]...)

	b, err := xml.MarshalIndent(us, "", "\t")
	if err != nil {
		return
	}

	err = os.WriteFile(file.Name(), b, 0)
	if err != nil {
		return
	}

	return
}
