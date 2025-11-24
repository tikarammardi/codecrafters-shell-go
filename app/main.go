package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var (
	_ = fmt.Fprint
	_ = os.Stdout
)

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
		parts := strings.Split(input, " ")

		cmd := parts[0]
		args := parts[1:]

		var firstArg string

		builtIncommandList := []string{"echo", "exit", "type"}

		if cmd == "exit" {

			if len(args) != 0 {
				firstArg = args[0]
			} else {
				firstArg = ""
			}

			if firstArg == "0" {
				os.Exit(0)
			}
			if firstArg == "1" {
				os.Exit(1)
			}
			os.Exit(0)
		} else if cmd == "echo" {
			fmt.Print(strings.Join(args, " ") + "\n")
		} else if cmd == "type" {
			firstArg = args[0]
			isBuiltInCommandInList := slices.Contains(builtIncommandList, firstArg)
			if isBuiltInCommandInList {
				fmt.Println(firstArg + " is a shell builtin")
			} else if path, ok := isCommandExecutableInPath(firstArg); ok {
				fmt.Println(firstArg + " is " + path)
			} else {
				fmt.Println(firstArg + ": not found")
			}

		} else {

			err := executeCommand(cmd, args)
			if err != nil {
				fmt.Println(cmd + ": command not found")
			}

		}
	}
}

func isCommandExecutableInPath(cmd string) (string, bool) {
	path, err := exec.LookPath(cmd)
	if err != nil {
		return "", false
	}
	return path, true
}

func executeCommand(cmd string, args []string) error {
	c := exec.Command(cmd, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Run()
	return err
}
