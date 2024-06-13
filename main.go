package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Try adding data to the block with a transaction. Once you have a block, add it to the chain.\n\n")

	fmt.Printf("Menu:\n\t1. New Transaction\n\t2. New Block\n\t3. Add to Blockchain\n\t4. Exit\n\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Option: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input: ", err)
			continue
		}

		input = strings.TrimSpace(input)

		switch {
		case input == "1":
			return
		case input == "2":
			return
		case input == "3":
			return
		case input == "4":
			fmt.Println("Goodbye. Thanks for trying my program out!")
			os.Exit(0)
			return
		default:
			fmt.Printf("Error in input or no value given\n")
		}
	}
}
