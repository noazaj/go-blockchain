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

func Transaction(data interface{}) interface{} {
	return data
}

func NewBlock(data string) *Block {
	block := &Block{
		Data:      data,
		Timestamp: time.Now(),
	}
	return block
}

func calculateHash(block *Block) string {
	record := fmt.Sprintf("%d%d%v%v", block.BlockNum, block.Nonce, block.Data, block.Timestamp)
	hash := sha256.New()
	hash.Write([]byte(record))
	return hex.EncodeToString(hash.Sum(nil))
}

func run(block *Block) {
	i := 0
	for {
		hash := calculateHash(block)
		fmt.Println(hash)
		if strings.HasPrefix(hash, strings.Repeat("0", difficulty)) {
			break
		}
		i++
	}
	block.Nonce = i
	fmt.Println(block.Nonce)
}
