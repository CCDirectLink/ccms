package help

import (
	"fmt"
	"os"
)

// Install help
func Install() {
	fmt.Printf("Usage: %s <options> install [packages...]\r\n", os.Args[0])
}
