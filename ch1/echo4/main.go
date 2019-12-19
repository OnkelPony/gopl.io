// Echo prints its command-line arguments on new lines including index
package main

import (
	"fmt"
	"os"
)

// !+
func main() {
	for i, argument := range os.Args[1:] {
		fmt.Printf("index: %d argument: %s\n", i, argument)
	}
}

// !-
