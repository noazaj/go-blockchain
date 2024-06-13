package block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

const (
	difficulty = 1
)

type Block struct {
	BlockNum  int         `json:"block_num"`
	Nonce     int         `json:"nonce"`
	Data      interface{} `json:"data"`
	PrevHash  string      `json:"prev_hash"`
	Hash      string      `json:"hash"`
	Timestamp time.Time   `json:"timestamp"`
}

func Transaction(data interface{}) interface{} {
	return data
}

func NewBlock(transaction interface{}) *Block {
	block := &Block{
		Data:      transaction,
		Timestamp: time.Now(),
	}
	return block
}

func calculateHash(block *Block) string {
	record := fmt.Sprintf("%d%d%v%v%v", block.BlockNum, block.Nonce, block.Data, block.PrevHash, block.Timestamp)
	hash := sha256.New()
	hash.Write([]byte(record))
	return hex.EncodeToString(hash.Sum(nil))
}

func Run(block *Block) {
	for {
		if strings.HasPrefix(calculateHash(block), strings.Repeat("0", difficulty)) {
			block.Hash = calculateHash(block)
			break
		}
		block.Nonce++
	}
}
