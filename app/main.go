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

		parts := strings.SplitN(input, " ", 2)

		cmd := parts[0]
		args := parts[1:]

		var firstArg string

		if len(args) != 0 {
			firstArg = args[0]
		} else {
			firstArg = ""
		}

		builtIncommandList := []string{"echo", "exit", "type"}

		isBuiltInCommandInList := slices.Contains(builtIncommandList, firstArg)

		if cmd == "exit" {
			if firstArg == "0" {
				os.Exit(0)
			}
			if firstArg == "1" {
				os.Exit(1)
			}
		} else if cmd == "echo" {
			fmt.Print(firstArg + "\n")
		} else if cmd == "type" {
			if isBuiltInCommandInList {
				fmt.Println(firstArg + " is a shell builtin")
			} else if isCommandExecutableInPath(firstArg) {
				fmt.Println(firstArg + " is a shell builtin")
			} else {
				fmt.Println(firstArg + ": not found")
			}
		} else {
			fmt.Println(cmd + ": command not found")
		}
	}
}

func isExecutable(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	mode := info.Mode()
	return mode&0o111 != 0, nil
}

func isCommandExecutableInPath(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
