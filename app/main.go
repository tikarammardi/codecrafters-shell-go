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

var builtIncommandList = []string{"echo", "exit", "type", "pwd"}

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

		handleCommands(cmd, args)
	}
}

func isCommandExecutableInPath(cmd string) (string, bool) {
	path, err := exec.LookPath(cmd)
	if err != nil {
		return "", false
	}
	return path, true
}

func executeCommand(cmd string, args []string) {
	c := exec.Command(cmd, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Run()
	if err != nil {
		fmt.Println(cmd + ": command not found")
	}
}

func handleCommands(cmd string, args []string) {
	switch cmd {
	case "exit":

		executeExitCommand(cmd, args)
	case "echo":
		executeEchoCommand(cmd, args)
	case "pwd":

		dir, _ := os.Getwd()
		fmt.Println(dir)
	case "type":
		executeTypeCommand(cmd, args)
	default:
		executeCommand(cmd, args)
	}
}

func executeExitCommand(cmd string, args []string) {
	var firstArg string
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
}

func executeEchoCommand(cmd string, args []string) {
	fmt.Print(strings.Join(args, " ") + "\n")
}

func executeTypeCommand(cmd string, args []string) {
	var firstArg string = args[0]
	isBuiltInCommandInList := slices.Contains(builtIncommandList, firstArg)
	if isBuiltInCommandInList {
		fmt.Println(firstArg + " is a shell builtin")
	} else if path, ok := isCommandExecutableInPath(firstArg); ok {
		fmt.Println(firstArg + " is " + path)
	} else {
		fmt.Println(firstArg + ": not found")
	}
}
