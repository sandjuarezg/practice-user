package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type users []User

type User struct {
	Name   string   `json:"name"`
	Passwd string   `json:"passwd"`
	Posts  []string `json:"posts"`
}

func AddUserToFile(u User) (err error) {
	file, err := os.OpenFile("./data/users.json", os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return
	}
	defer file.Close()

	s, err := file.Stat()
	if err != nil {
		return
	}

	var us users

	if s.Size() != 0 {
		err = json.NewDecoder(file).Decode(&us)
		if err != nil {
			return
		}

		_, err = file.Seek(0, 0)
		if err != nil {
			return
		}

	}

	us = append(us, u)

	err = file.Truncate(0)
	if err != nil {
		return
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(us)
	if err != nil {
		return
	}

	return
}

func LogIn(name, passwd string) (u User, err error) {
	file, err := os.Open("./data/users.json")
	if err != nil {
		return
	}
	defer file.Close()

	var us users

	err = json.NewDecoder(file).Decode(&us)
	if err != nil {
		return
	}

	var ban bool
	for _, v := range us {
		if v.Name == name && v.Passwd == passwd {
			ban = true
			u = v
			break
		}
	}

	if !ban {
		err = errors.New("user not found")
		return
	}

	return
}

func ShowPostByName(name string) (err error) {
	file, err := os.Open("./data/users.json")
	if err != nil {
		return
	}
	defer file.Close()

	var us users

	err = json.NewDecoder(file).Decode(&us)
	if err != nil {
		return
	}

	for _, v := range us {
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
	file, err := os.OpenFile("./data/users.json", os.O_RDWR, 0600)
	if err != nil {
		return
	}
	defer file.Close()

	var us users

	err = json.NewDecoder(file).Decode(&us)
	if err != nil {
		return
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return
	}

	var ban bool
	for n, v := range us {
		if v.Name == u.Name {
			ban = true
			us[n].Posts = append(v.Posts, postText)
			break
		}
	}

	if !ban {
		err = errors.New("user not found")
		return
	}

	err = file.Truncate(0)
	if err != nil {
		return
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(us)
	if err != nil {
		return
	}

	return
}

func (u User) EditPost(postIndex int, newPost string) (err error) {
	file, err := os.OpenFile("./data/users.json", os.O_RDWR, 0600)
	if err != nil {
		return
	}
	defer file.Close()

	var us users

	err = json.NewDecoder(file).Decode(&us)
	if err != nil {
		return
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return
	}

	var ban bool
	for n, v := range us {
		if v.Name == u.Name {

			ban = true
			if postIndex > len(v.Posts)-1 || postIndex < 0 {
				err = errors.New("number out of range")
				return
			}

			us[n].Posts[postIndex] = newPost

			break
		}
	}

	if !ban {
		err = errors.New("user not found")
		return
	}

	err = file.Truncate(0)
	if err != nil {
		return
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(us)
	if err != nil {
		return
	}

	return
}

func (u User) DeletePost(postIndex int) (err error) {
	file, err := os.OpenFile("./data/users.json", os.O_RDWR, 0600)
	if err != nil {
		return
	}
	defer file.Close()

	var us users

	err = json.NewDecoder(file).Decode(&us)
	if err != nil {
		return
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return
	}

	var ban bool
	for n, v := range us {
		if v.Name == u.Name {

			ban = true
			if postIndex > len(v.Posts)-1 || postIndex < 0 {
				err = errors.New("number out of range")
				return
			}

			us[n].Posts = append(v.Posts[:postIndex], v.Posts[postIndex+1:]...)

			break
		}
	}

	if !ban {
		err = errors.New("user not found")
		return
	}

	err = file.Truncate(0)
	if err != nil {
		return
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(us)
	if err != nil {
		return
	}

	return
}
