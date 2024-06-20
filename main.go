package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/noazaj/go-blockchain/block"
)

var (
	transactionPool []interface{}
	blockPool       []*block.Block
)

func main() {
	fmt.Printf("Try adding data to the block with a transaction. Once you have a block, add it to the chain.\n\n")

	fmt.Printf("Menu:\n\t1. New Transaction\n\t2. New Block\n\t3. Add to Blockchain\n\t4. Transaction Pool\n\t5. Block Pool\n\t6. Exit\n\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("\nOption: ")
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
			addBlock()
		case input == "3":
			fmt.Println("Option not yet implemented!")
		case input == "4":
			viewTransactionPool()
		case input == "5":
			viewBlockPool()
		case input == "6" || input == "exit":
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

func addBlock() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Transaction Pool or New -> (1/2): ")
	choice, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		reader := bufio.NewReader(os.Stdin)
		viewTransactionPool()

		fmt.Printf("\nSelect a transaction number: ")
		num, _ := reader.ReadString('\n')
		num = strings.TrimSpace(num)

		for i, transaction := range transactionPool {
			if strconv.Itoa(i) == num {
				block := block.NewBlock(transaction)
				blockPool = append(blockPool, block)

				index, _ := strconv.Atoi(num)
				transactionPool[index] = transactionPool[len(transactionPool)-1]
				transactionPool = transactionPool[:len(transactionPool)-1]

				fmt.Println("Block added to pool. View block(s) with other option.")
				return
			}
		}

		fmt.Printf("Your number selection was not a given transaction")
		return

	case "2":
		data, err := block.Transaction()
		if err != nil {
			log.Printf("error: %v", err)
			return
		}
		block := block.NewBlock(data)
		blockPool = append(blockPool, block)
		fmt.Println("Block added to pool. View block(s) with other option.")
		return
	default:
		fmt.Printf("Error in input or no value given\n")
	}
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

func viewBlockPool() {
	if len(blockPool) == 0 {
		fmt.Println("No blocks present in block pool.")
		return
	}

	for i, block := range blockPool {
		fmt.Printf("Block %d: %+v", i, block)
	}
}
