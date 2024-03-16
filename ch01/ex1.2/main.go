package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		fmt.Printf("%d - %s\n", i, arg)
	}
}
