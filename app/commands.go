package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// CmdResult represents the outcome of running a builtin command.
type CmdResult struct {
	Handled bool // whether the command was a builtin and handled
	Exit    bool // whether the command requested the shell to exit
	Code    int  // exit code to use when exiting
}

// builtins is a registry of builtin command handlers.
var builtins map[string]func([]string) CmdResult

func init() {
	builtins = map[string]func([]string) CmdResult{
		"echo": func(args []string) CmdResult {
			fmt.Print(strings.Join(args, " ") + "\n")
			return CmdResult{Handled: true}
		},
		"pwd": func(args []string) CmdResult {
			dir, _ := os.Getwd()
			fmt.Println(dir)
			return CmdResult{Handled: true}
		},
		"type": func(args []string) CmdResult {
			if len(args) == 0 {
				fmt.Println("type: missing operand")
				return CmdResult{Handled: true}
			}
			name := args[0]
			// check if it's a builtin
			if _, ok := builtins[name]; ok {
				fmt.Println(name + " is a shell builtin")
				return CmdResult{Handled: true}
			}
			if path, ok := isCommandExecutableInPath(name); ok {
				fmt.Println(name + " is " + path)
			} else {
				fmt.Println(name + ": not found")
			}
			return CmdResult{Handled: true}
		},
		"exit": func(args []string) CmdResult {
			code := 0
			if len(args) > 0 {
				if v, err := strconv.Atoi(args[0]); err == nil {
					code = v
				}
			}
			return CmdResult{Handled: true, Exit: true, Code: code}
		},
	}
}

// RunBuiltin attempts to run a builtin command by name. It returns whether
// the command was handled, whether the handler requested an exit, and the exit code.
func RunBuiltin(cmd string, args []string) (handled bool, exit bool, code int) {
	if h, ok := builtins[cmd]; ok {
		res := h(args)
		return res.Handled, res.Exit, res.Code
	}
	return false, false, 0
}

// isCommandExecutableInPath checks whether an executable exists in PATH.
func isCommandExecutableInPath(cmd string) (string, bool) {
	path, err := exec.LookPath(cmd)
	if err != nil {
		return "", false
	}
	return path, true
}

// executeExternalCommand runs a program from the filesystem.
func executeExternalCommand(cmd string, args []string) error {
	c := exec.Command(cmd, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
