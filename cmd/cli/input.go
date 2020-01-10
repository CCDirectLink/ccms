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
