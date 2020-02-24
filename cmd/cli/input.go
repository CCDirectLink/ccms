package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

// Prompt ask for input based upon a prompt
func Prompt(prompt string) string {
	fmt.Printf("%s", prompt)
	line, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}
	return strings.TrimRight(line, "\r\n")
}

// AskYesNo question to user
func AskYesNo(prompt string) bool {
	fmt.Printf("%s [Y/n]", prompt)
	line, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	line = strings.TrimRight(line, "\r\n")

	if line == "Y" {
		return true
	}

	return false
}
