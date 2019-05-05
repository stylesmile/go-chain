package core

import (
	"fmt"
	"log"
)

type Blockchain struct {
	// 集合 区块链
	Blocks []*Block
}

// 区块链
func NewBlockchain() *Blockchain {
	genesisBlock := GenerateGenesisBlock()
	blockchain := Blockchain{}
	blockchain.ApendBlock(&genesisBlock)
	return &blockchain
}

// 发送数据
func (bc *Blockchain) SendData(data string) {
	proBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*proBlock, data)
	bc.ApendBlock(&newBlock)
}

// 添加区块
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

// 后台数据结构
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

// 验证相邻的2个区块数据是否正确
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
