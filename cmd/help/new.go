package help

import (
	"fmt"
	"os"
)

// New help
func New() {
	fmt.Printf("Usage: %s new <name>\r\n", os.Args[0])
}
