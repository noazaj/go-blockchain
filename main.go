package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/noazaj/go-blockchain/block"
)

var (
	transactionPool []interface{}
)

func main() {
	fmt.Printf("Try adding data to the block with a transaction. Once you have a block, add it to the chain.\n\n")

	fmt.Printf("Menu:\n\t1. New Transaction\n\t2. New Block\n\t3. Add to Blockchain\n\t4. Transaction Pool\n\t5. Exit\n\n")

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
			addTransaction()
		case input == "2":
			fmt.Println("Option not yet implemented!")
		case input == "3":
			fmt.Println("Option not yet implemented!")
		case input == "4":
			viewTransactionPool()
		case input == "5" || input == "exit":
			fmt.Println("Goodbye. Thanks for trying my program out!")
			os.Exit(0)
			return
		default:
			fmt.Printf("Error in input or no value given\n")
		}
	}
}

func addTransaction() {
	data, err := block.Transaction()
	if err != nil {
		log.Print(err)
		return
	}
	transactionPool = append(transactionPool, data)
	fmt.Println("Transaction added to pool. View transaction(s) with other option.")
}

func viewTransactionPool() {
	if len(transactionPool) == 0 {
		fmt.Println("No transactions present in transaction pool.")
		return
	}

	for i, transaction := range transactionPool {
		fmt.Printf("Transaction %d: %v", i, transaction)
	}
}
