//go:generate pigeon -o parser/main.go parser/grammar.peg
//go:generate goimports -w parser/main.go
package main

import (
	"fmt"
	"log"

	"buildyourownlisp/parser"

	"github.com/peterh/liner"
)

func main() {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	// Print Version and Exit Information.
	fmt.Println("Lispy Version 0.0.0.0.1")
	fmt.Println("Press Ctrl+c to Exit")
	fmt.Println()

	// In a never ending loop.
	for {
		// Output our prompt and read a line of user input.
		if input, err := line.Prompt("lispy> "); err == nil {
			// Echo input back to user
			fmt.Printf("No you're a %s\n", input)
			line.AppendHistory(input)

			ast, err := parser.Parse("interpreter", []byte(input))
			fmt.Printf("ast = %+v\n", ast)
			fmt.Printf("err = %+v\n", err)
		} else if err == liner.ErrPromptAborted {
			break
		} else {
			log.Print("Error reading line: ", err)
		}
	}
}
