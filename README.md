# go-supports-hyperlinks

Fully based in https://github.com/jamestalmage/supports-hyperlinks. Detect whether a terminal emulator supports hyperlinks.

## Install

```console
$ go get -u github.com/luisnquin/go-supports-hyperlinks
```

## Usage

```go
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

```

## Platform support

- [x] Linux
- [x] Darwin
- [ ] Windows

Still, read the **CLI section** if you want to force it on Windows

## API

Returns a boolean API indicating whether or not hyperlinks are supported in the current terminal

## CLI

Still, you can force the behavior with

```console
$ go run program.go --hyperlink
$ go run program.go --hyperlinks
```

Otherwise

```console
$ go run program.go --no-hyperlink
$ go run program.go --no-hyperlinks
```

## License

MIT © [Luis Quiñones Requelme](https://github.com/luisnquin)
