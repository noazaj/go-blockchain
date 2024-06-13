package main

import (
	"fmt"

	"github.com/noazaj/go-blockchain/block"
)

func main() {
	transaction1 := block.Transaction("Here is some data!")
	transaction2 := block.Transaction(5)

	fmt.Printf("%v\n%v", transaction1, transaction2)
}
