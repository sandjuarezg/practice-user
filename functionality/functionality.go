package functionality

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func CleanConsole() (err error) {
	fmt.Println(". . . . . . . . . . . .")
	time.Sleep(3 * time.Second)

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		return
	}

	return
}

func ScanNamePasswd() (name, passwd string, err error) {
	fmt.Println("Enter user name")
	aux, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		err = errors.New("error to find name")
		return
	}
	name = string(aux)

	fmt.Println()

	fmt.Println("Enter password")
	aux, _, err = bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		err = errors.New("error to find password")
		return
	}
	passwd = string(aux)

	return
}

func ScanName() (name string, err error) {
	fmt.Println("Enter user name")
	aux, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		err = errors.New("error to find name")
		return
	}
	name = string(aux)

	return
}

func ScanPostText() (post string, err error) {
	fmt.Print("Enter text: ")
	aux, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		err = errors.New("error to write post")
		return
	}

	post = string(aux)

	return
}
