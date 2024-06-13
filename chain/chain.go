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

func (bc *Blockchain) AddBlock(newBlock *block.Block) {
	if len(bc.blocks) == 0 {
		log.Print("No blocks in blockchain or blockchain not instantiated. Try creating a blockchain.")
		return
	}

	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock.BlockNum = prevBlock.BlockNum + 1
	newBlock.PrevHash = prevBlock.Hash
	block.Run(newBlock)

	err := validateHash(prevBlock, newBlock)
	if err != nil {
		log.Print("Hashes do not match")
		return
	}

	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) PrintBlockchain() {
	fmt.Println("Blockchain: ")
	for _, block := range bc.blocks {
		fmt.Printf("%+v\n", block)
	}
}

func genesisBlock() *block.Block {
	genBlock := block.NewBlock("")
	block.Run(genBlock)
	return genBlock
}

func validateHash(prevBlock, newBlock *block.Block) error {
	if prevBlock.Hash != newBlock.PrevHash {
		return fmt.Errorf("error: previous hash does not match new hash")
	}
	return nil
}
