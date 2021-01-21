package main

import (
	"fmt"
	"log"

	"github.com/peterh/liner"
)

func main() {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	// Print Version and Exit Information.
	fmt.Println("Lispy Version 0.0.0.0.1")
	fmt.Println("Press Ctrl+c to Exit\n")

	// In a never ending loop.
	for {
		// Output our prompt and read a line of user input.
		if input, err := line.Prompt("lispy> "); err == nil {
			// Echo input back to user
			fmt.Printf("No you're a %s\n", input)
			line.AppendHistory(input)
		} else if err == liner.ErrPromptAborted {
			break
		} else {
			log.Print("Error reading line: ", err)
		}
	}
}
