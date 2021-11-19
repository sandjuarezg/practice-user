package functionality

import (
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
