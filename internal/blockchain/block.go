package blockchain

import (
    "fmt"
    "time"
)

type Block struct {
    Transactions []string
    PrevPointer  *Block
    PrevHash     string
    CurrentHash  string
    Timestamp    int64
}

func NewBlock(transactions []string, prevBlock *Block) *Block {
    fmt.Printf("DEBUG: Creating new block with transactions: %v\n", transactions)
    
    block := &Block{
        Transactions: transactions,
        PrevPointer:  prevBlock,
        PrevHash:     "",
        Timestamp:    time.Now().Unix(),
    }
    
    if prevBlock != nil {
        block.PrevHash = prevBlock.CurrentHash
    }
    
    block.CurrentHash = CalculateHash(block)
    return block
}
