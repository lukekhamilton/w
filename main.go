package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	cmd := os.Args[1:]
	fmt.Println(cmd)

	out, err := exec.Command(cmd[0], cmd[1:]...).Output()

	// cmd := GetCommand()
	// fmt.Println("cmd: ", cmd)
	// out, err := exec.Command(cmd[0], cmd[1:]...).Output()
	fmt.Printf("out: %s", out)
	fmt.Println("err: ", err)
}

// GetCommand ...
// func GetCommand() []string {
// 	return []string{"ifconfig", "-a"}
// }
