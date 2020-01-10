package help

import (
	"fmt"
	"os"
)

// Default help
func Default() {
	fmt.Printf("Usage: %s <options> <operation> args\r\n", os.Args[0])
	fmt.Println("<operation> = new, init, install")
}
