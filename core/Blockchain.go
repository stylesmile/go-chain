package core

import (
	"fmt"
	"log"
)

type Blockchain struct {
	// 数组
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	genesisBlock := GenerateGenesisBlock()
	blockchain := Blockchain{}
	blockchain.ApendBlock(&genesisBlock)
	return &blockchain
}

func (bc *Blockchain) SendData(data string) {
	proBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*proBlock, data)
	bc.ApendBlock(&newBlock)
}

func (bc *Blockchain) ApendBlock(newBlock *Block) {
	// 防止越界
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block")
	}
}

func (bc *Blockchain) Print() {
	// 下划线表示赋值是忽略的
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Prev.Hash: %s\n", block.PreBolckHash)
		fmt.Printf("Curr.Hash: %s\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timestmp: %d\n", block.Timestmp)
		fmt.Println("")
	}
}

func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	if newBlock.PreBolckHash != oldBlock.Hash {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}
