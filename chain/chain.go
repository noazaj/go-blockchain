package chain

import (
	"fmt"
	"log"

	"github.com/noazaj/go-blockchain/block"
)

type Blockchain struct {
	blocks []*block.Block
}

func NewBlockchain() *Blockchain {
	genBlock := genesisBlock()
	return &Blockchain{blocks: []*block.Block{genBlock}}
}

func (bc *Blockchain) AddBlock(block *block.Block) {
	if len(bc.blocks) == 0 {
		log.Print("No blocks in blockchain or blockchain not instantiated. Try creating a blockchain.")
		return
	}

	prevBlock := bc.blocks[len(bc.blocks)-1]
	err := validateHash(prevBlock, block)
	if err != nil {
		log.Print("Hashes do not match")
		return
	}

	bc.blocks = append(bc.blocks, block)
}

func (bc *Blockchain) PrintBlockchain() {
	fmt.Print("Blockchain: ")
	for _, block := range bc.blocks {
		fmt.Printf("%+v", block)
	}
}

func genesisBlock() *block.Block {
	genesis := block.NewBlock("")
	return genesis
}

func validateHash(prevBlock, newBlock *block.Block) error {
	if prevBlock.Hash != newBlock.PrevHash {
		return fmt.Errorf("error: previous hash does not match new hash")
	}
	return nil
}
