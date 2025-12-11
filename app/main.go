package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var (
	_ = fmt.Fprint
	_ = os.Stdout
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		_, err := fmt.Fprint(os.Stdout, "$ ")
		if err != nil {
			return
		}

		line, err := reader.ReadString('\n')
		if err != nil {
			// EOF or read error - exit loop
			return
		}

		input := strings.TrimSpace(line)
		if input == "" {
			continue
		}

		// simple splitting on whitespace
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}
		cmd := parts[0]
		args := parts[1:]

		handled, exit, code := RunBuiltin(cmd, args)
		if handled {
			if exit {
				os.Exit(code)
			}
			continue
		}

		// Not a builtin - try to execute external command
		err = executeExternalCommand(cmd, args)
		if err != nil {
			// Keep message similar to previous behavior
			fmt.Println(cmd + ": command not found")
		}
	}
}
