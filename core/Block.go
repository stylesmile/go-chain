package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index        int64  // 区块编号
	Timestmp     int64  // 区块时间戳
	PreBolckHash string // 上一个区块的哈希值
	Hash         string // 当前区块哈希值
	Data         string // 区块数据
}

// 计算哈希
func CalculateHash(b Block) string {
	// 数据不允许修改
	// 把 数据 包含在hash运算的数值里面， ，任何一个模块数据的修改，都会导致区块链， 链式数据结构的破坏
	blockData := string(b.Index) + string(b.Timestmp) + string(b.PreBolckHash)
	HashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(HashInBytes[:])
}

// 生成新区块
func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	// 新区块的index = 父区块的index+1
	newBlock.Index = preBlock.Index + 1
	// 哈希值 = 父区块的哈希值
	newBlock.PreBolckHash = preBlock.Hash
	// 新区块的时间戳，==当前时间
	newBlock.Timestmp = time.Now().Unix()
	newBlock.Data = data
	// 新区块的哈希 是计算得来的
	newBlock.Hash = CalculateHash(newBlock)
	return newBlock
}

// 生成创始区块
func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	GenerateNewBlock(preBlock, "Genesis Block")
	return preBlock
}
