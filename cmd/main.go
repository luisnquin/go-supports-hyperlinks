package main

import (
	"fmt"
	"os"

	shl "github.com/luisnquin/go-supports-hyperlinks"
)

func main() {
	if shl.Stdout() {
		fmt.Println("Terminal stdout supports hyperlinks")
	}

	if shl.Stderr() {
		fmt.Println("Terminal stderr supports hyperlinks")
	}

	if shl.SupportsHyperlinks(os.Stdin) {
		fmt.Println("Terminal stdin supports hyperlinks")
	}
}
