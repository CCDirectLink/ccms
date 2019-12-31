package util

import (
	"fmt"
	"os"
)

// PrintHelp prints the help message
func PrintHelp() {
	fmt.Printf("Usage: %s <operation>\r\n", os.Args[0])
	fmt.Println("<operation> = init")
}
