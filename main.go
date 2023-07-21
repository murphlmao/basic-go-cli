// main.go

package main

import (
    "fmt"
	"os"
	"strconv"
	"github.com/murphlmao/basic-go-cli/cli"
)


func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command provided.")
		os.Exit(1)
	}

	cmd := os.Args[1]
	name := os.Args[2]

	switch cmd {
	case "--hello":
		result := basics.Hello(name)
		fmt.Println(result)

	case "--math":
		num, err := strconv.Atoi(name)
		if err != nil || num < 0 || num > 255 {
			fmt.Println("Please provide a number between 0 and 255 for the math command.")
			os.Exit(1)
		}
		
		result := basics.IAmDoMathing(uint8(num))
		fmt.Println(result)
		
	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		os.Exit(1)
	}
}

/*
// do_the_basics(10, "Purr")
func do_the_basics(num uint8, name string) {
	result := basics.IAmDoMathing(num) // does the most complicated math on planet earth
    fmt.Println(result)

	var result2 string = basics.Hello(name) // says some shit
	fmt.Println(result2)
}
*/