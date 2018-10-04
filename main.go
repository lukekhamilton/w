package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := GetCommand()
	fmt.Println("cmd: ", cmd)
	out, err := exec.Command(cmd[0], cmd[1:]...).Output()
	fmt.Printf("out: %s", out)
	fmt.Println("err: ", err)
}

// GetCommand ...
func GetCommand() []string {
	return []string{"ifconfig", "-a"}
}
