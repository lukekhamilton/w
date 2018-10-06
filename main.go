package main

import (
	"fmt"

	git "gopkg.in/src-d/go-git.v4"
)

func main() {

	// cmd := os.Args[1:]
	// fmt.Println(cmd)

	// out, err := exec.Command(cmd[0], cmd[1:]...).Output()

	// // cmd := GetCommand()
	// // fmt.Println("cmd: ", cmd)
	// // out, err := exec.Command(cmd[0], cmd[1:]...).Output()
	// fmt.Printf("out: %s", out)
	// fmt.Println("err: ", err)

	// Info("git clone https://github.com/src-d/go-git")

	_, err := git.PlainClone("./go-git", false, &git.CloneOptions{
		URL: "https://github.com/src-d/go-git",
		// Progress: os.Stdout,
	})

	// CheckIfError(err)
	fmt.Println("err: ", err)
}

// GetCommand ...
// func GetCommand() []string {
// 	return []string{"ifconfig", "-a"}
// }
