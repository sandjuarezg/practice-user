package functionality

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func CleanConsole() {
	fmt.Println(". . . . . . . . . . . .")
	time.Sleep(3 * time.Second)

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
