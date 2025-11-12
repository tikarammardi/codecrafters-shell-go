package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		input := command[:len(command)-1]

		parts := strings.SplitN(input, " ", 2)

		cmd := parts[0]
		args := parts[1:]

		var firstArg string

		if len(args) != 0 {
			firstArg = args[0]
		} else {
			firstArg = ""
		}

		if cmd == "exit" {
			if firstArg == "0" {
				os.Exit(0)
			}
			if firstArg == "1" {
				os.Exit(1)
			}
		} else if cmd == "echo" {
			fmt.Print(firstArg + "\n")
		} else {

			fmt.Println(cmd + ": command not found")
		}

	}
}
