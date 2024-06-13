package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

const (
	difficulty = 4
)

type Block struct {
	BlockNum  int       `json:"block_num"`
	Nonce     int       `json:"nonce"`
	Data      string    `json:"data"`
	PrevHash  string    `json:"prev_hash"`
	Hash      string    `json:"hash"`
	Timestamp time.Time `json:"timestamp"`
}

func NewBlock(data string) *Block {
	block := &Block{
		Data:      data,
		Timestamp: time.Now(),
	}
	block.Nonce = run(block)
	return block
}

func CalculateHash(block *Block) string {
	record := fmt.Sprintf("%d%d%v%v", block.BlockNum, block.Nonce, block.Data, block.Timestamp)
	hash := sha256.New()
	hash.Write([]byte(record))
	return hex.EncodeToString(hash.Sum(nil))
}

func run(block *Block) int {
	for {
		hash := CalculateHash(block)
		fmt.Println(hash)
		if strings.HasPrefix(hash, strings.Repeat("0", difficulty)) {
			break
		}
		block.Nonce++
	}
	fmt.Println(block.Nonce)
	return block.Nonce
}
