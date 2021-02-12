//go:generate pigeon -o parser/main.go parser/grammar.peg
//go:generate goimports -w parser/main.go
package main

import (
	"fmt"

	"buildyourownlisp/evaluator"
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

	environment := evaluator.NewRootEnvironment()

	// In a never ending loop.
	for {
		// Output our prompt and read a line of user input.
		if input, err := line.Prompt("lispy> "); err == nil {
			line.AppendHistory(input)

			ast, err := parser.Parse("interpreter", []byte(input))
			if err != nil {
				fmt.Printf("err = %+v\n", err)
				continue
			}
			result, err := evaluator.Evaluate(environment, ast)
			if err != nil {
				fmt.Printf("ERROR: %+v\n", err)
				continue
			}
			fmt.Printf("%v\n", result)
		} else if err == liner.ErrPromptAborted {
			break
		} else {
			fmt.Printf("Error reading line: %v\n", err)
		}
	}
}
