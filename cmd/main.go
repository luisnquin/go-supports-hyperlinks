package main

import (
	"fmt"

	shl "github.com/luisnquin/go-supports-hyperlinks"
)

func main() {
	if shl.Stdout {
		fmt.Println("Terminal stdout supports hyperlinks")
	}

	if shl.Stderr {
		fmt.Println("Terminal stderr supports hyperlinks")
	}
}
