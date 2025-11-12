package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

func main() {

	for {
		_, err := fmt.Fprint(os.Stdout, "$ ")
		if err != nil {
			return
		}

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}
		cmd := command[:len(command)-1]
		if cmd == "exit" {
			break
		}
		fmt.Println(cmd + ": command not found")

	}
}
