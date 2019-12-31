package cli

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

// Prompt ask for input based upon a prompt
func Prompt(prompt string) string {
	fmt.Printf("%s", prompt)
	line, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	if line[len(line)-2:len(line)] == "\r\n" {
		line = line[0 : len(line)-2]
	} else if line[len(line)-1] == '\n' {
		line = line[0 : len(line)-1]
	}

	return line
}
