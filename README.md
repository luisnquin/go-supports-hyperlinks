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
```
## API
Returns a boolean API indicating whether or not hyperlinks are supported in the current terminal

## CLI
Still, you can force the behavior with ```--hyperlink```, otherwise ```--no-hyperlink``` or ```--no-hyperlinks```

## License 
MIT © [Luis Quiñones Requelme](https://github.com/luisnquin)
