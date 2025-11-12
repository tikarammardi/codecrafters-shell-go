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
	fmt.Fprint(os.Stdout, "$ ")
	command, error := bufio.NewReader(os.Stdin).ReadString('\n')
	if error != nil {
		panic(error)
	}
	cmd := command[:len(command)-1]

	fmt.Println(cmd + ": command not found")
}
