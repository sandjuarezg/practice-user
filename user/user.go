package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type users struct {
	Users []User `json:"user"`
}

type User struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
	Posts  []post `json:"posts"`
}

type post struct {
	Post string `json:"post"`
}

func GetAllDataFromTag(r io.Reader, tag string) (data []json.Token) {
	decoder := json.NewDecoder(r)

	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}

		_, ok := t.(json.Delim)
		if ok {
			continue
		}

		if t == tag {

			s, err := decoder.Token()
			if err == io.EOF {
				break
			}

			data = append(data, s)

		}

	}

	return
}

func AddUserToFile(u User) (err error) {
	file, err := os.OpenFile("./data/users.json", os.O_CREATE|os.O_RDONLY, 0600)
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
		err = json.Unmarshal(content, &us)
		if err != nil {
			return
		}
	}

	us.Users = append(us.Users, u)

	b, err := json.MarshalIndent(us, "", "\t")
	if err != nil {
		return
	}

	_, err = file.WriteAt(b, 0)
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

	names := GetAllDataFromTag(file, "name")
	file.Seek(0, 0)

	passwds := GetAllDataFromTag(file, "passwd")
	err = errors.New("user not found")

	for i := range names {

		for y := range passwds {
			if i != y {
				continue
			}

			if names[i] == name && passwds[y] == passwd {
				u.Name = name
				u.Passwd = passwd

				err = nil
				break
			}
		}
	}

	return
}

func ShowPostByName(name string) (err error) {
	file, err := os.Open("./data/users.json")
	if err != nil {
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}

		if t == "name" {

			s, _ := decoder.Token()
			if s == nil {
				break
			}

			if s != name {
				continue
			}

			fmt.Printf("Post's %s:\n", name)
			var i int

			for {
				t, _ = decoder.Token()
				if t == nil {
					break
				}

				if t == "name" {
					break
				}

				if t == "post" {
					t, _ = decoder.Token()
					if t == nil {
						break
					}

					fmt.Printf("Key: %d\n", i+1)
					fmt.Printf("Content: %s\n", t)
					fmt.Println()
					i++
				}
			}

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

	content, err := io.ReadAll(file)
	if err != nil {
		return
	}

	var us users

	err = json.Unmarshal(content, &us)
	if err != nil {
		return
	}

	err = errors.New("user not found")
	for n, v := range us.Users {
		if v.Name == u.Name {
			us.Users[n].Posts = append(us.Users[n].Posts, post{postText})
			err = nil
			break
		}
	}

	b, err := json.MarshalIndent(us, "", "\t")
	if err != nil {
		return
	}

	_, err = file.WriteAt(b, 0)
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

	content, err := io.ReadAll(file)
	if err != nil {
		return
	}

	var us users

	err = json.Unmarshal(content, &us)
	if err != nil {
		return
	}

	err = errors.New("user not found")
	for n, v := range us.Users {
		if v.Name == u.Name {

			postIndex--
			if postIndex > len(us.Users[n].Posts)-1 || postIndex < 0 {
				err = errors.New("number out of range")
				return
			}

			us.Users[n].Posts[postIndex].Post = newPost
			err = nil

			break
		}
	}

	b, err := json.MarshalIndent(us, "", "\t")
	if err != nil {
		return
	}

	err = file.Truncate(0)
	if err != nil {
		return
	}

	_, err = file.WriteAt(b, 0)
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

	content, err := io.ReadAll(file)
	if err != nil {
		return
	}

	var us users

	err = json.Unmarshal(content, &us)
	if err != nil {
		return
	}

	err = errors.New("user not found")
	for n, v := range us.Users {
		if v.Name == u.Name {

			postIndex--
			if postIndex > len(us.Users[n].Posts)-1 || postIndex < 0 {
				err = errors.New("number out of range")
				return
			}

			us.Users[n].Posts = append(us.Users[n].Posts[:postIndex], us.Users[n].Posts[postIndex+1:]...)
			err = nil

			break
		}
	}

	b, err := json.MarshalIndent(us, "", "\t")
	if err != nil {
		return
	}

	err = file.Truncate(0)
	if err != nil {
		return
	}

	_, err = file.WriteAt(b, 0)
	if err != nil {
		return
	}

	return
}
