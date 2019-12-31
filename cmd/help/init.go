package help

import (
	"fmt"
	"os"
)

// Init help
func Init() {
	fmt.Printf("Usage: %s init\r\n", os.Args[0])
}
