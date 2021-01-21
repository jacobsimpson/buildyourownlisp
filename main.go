package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/* Print Version and Exit Information */
	fmt.Println("Lispy Version 0.0.0.0.1")
	fmt.Println("Press Ctrl+c to Exit\n")

	/* In a never ending loop */
	for {
		/* Output our prompt */
		fmt.Printf("lispy> ")

		/* Read a line of user input of maximum size 2048 */
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		/* Echo input back to user */
		fmt.Printf("No you're a %s\n", input)
	}
}
