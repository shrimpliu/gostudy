package main

import (
	"fmt"
	"os"
)

func main() {
	for key, value := range os.Args[1:] {
		fmt.Printf("%d : %s\n", key, value)
	}
}
