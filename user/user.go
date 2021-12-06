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
	XMLName xml.Name `xml:"users"`

	Users []User `xml:"user"`
}

type User struct {
	Name   string `xml:"name"`
	Passwd string `xml:"passwd"`
	Posts  []post `xml:"posts"`
}

type post struct {
	Post string `xml:"post"`
}

func GetAllDataFromTag(r io.Reader, tag string) (data []string) {
	decoder := xml.NewDecoder(r)

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		s, ok := t.(xml.StartElement)
		if !ok {
			continue
		}

		if s.Name.Local == tag {
			t, _ = decoder.Token()
			if t == nil {
				break
			}

			s, ok := t.(xml.CharData)
			if ok {
				data = append(data, string(s))
			}
		}
	}

	return
}

func AddUserToFile(u User) (err error) {
	file, err := os.OpenFile("./data/users.xml", os.O_CREATE|os.O_RDONLY, 0600)
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

	_, err = file.WriteAt(b, 0)
	if err != nil {
		return
	}

	return
}

func LogIn(name, passwd string) (u User, err error) {
	file, err := os.Open("./data/users.xml")
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
	file, err := os.Open("./data/users.xml")
	if err != nil {
		return
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		s, ok := t.(xml.StartElement)
		if !ok {
			continue
		}

		if s.Name.Local == "name" {
			t, _ = decoder.Token()
			if t == nil {
				break
			}

			s, ok := t.(xml.CharData)
			if !ok {
				break
			}

			if string(s) != name {
				continue
			}

			fmt.Printf("Post's %s:\n", name)
			var i int

			for {
				t, _ = decoder.Token()
				if t == nil {
					break
				}

				data, ok := t.(xml.StartElement)
				if !ok {
					continue
				}

				if data.Name.Local == "user" {
					break
				}

				if data.Name.Local == "post" {
					t, _ = decoder.Token()
					if t == nil {
						break
					}

					s, ok := t.(xml.CharData)
					if !ok {
						break
					}

					fmt.Printf("Key: %d\n", i+1)
					fmt.Printf("Content: %s\n", s)
					fmt.Println()
					i++
				}
			}

		}
	}

	return
}

func (u User) AddPostToFile(postText string) (err error) {
	file, err := os.OpenFile("./data/users.xml", os.O_RDWR, 0600)
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

	err = errors.New("user not found")
	for n, v := range us.Users {
		if v.Name == u.Name {
			us.Users[n].Posts = append(us.Users[n].Posts, post{postText})
			err = nil
			break
		}
	}

	b, err := xml.MarshalIndent(us, "", "\t")
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
	file, err := os.OpenFile("./data/users.xml", os.O_RDWR, 0600)
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

	b, err := xml.MarshalIndent(us, "", "\t")
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
	file, err := os.OpenFile("./data/users.xml", os.O_RDWR, 0600)
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

	b, err := xml.MarshalIndent(us, "", "\t")
	if err != nil {
		return
	}

	_, err = file.WriteAt(b, 0)
	if err != nil {
		return
	}

	return
}
