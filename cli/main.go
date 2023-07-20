package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided.")
		os.Exit(1)
	}

	cmd := os.Args[1]
	name := os.Args[2]

	switch cmd {
	case "hello":
		hello(name)
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		os.Exit(1)
	}
}

func hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
