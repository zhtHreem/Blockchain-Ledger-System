// File: internal/blockchain/storage.go
package blockchain

import (
    "encoding/json"
    "fmt"
    "os"
)

const dataFile = "blockchain.json"

type BlockData struct {
    Transactions []string `json:"transactions"`
    PrevHash    string   `json:"prevHash"`
    CurrentHash string   `json:"currentHash"`
    Timestamp   int64    `json:"timestamp"`
}

type ChainData struct {
    Blocks []BlockData `json:"blocks"`
}

func (c *Chain) SaveChain() error {
    var chainData ChainData
    current := c.Head
    
    // Convert chain to serializable format
    for current != nil {
        blockData := BlockData{
            Transactions: current.Transactions,
            PrevHash:    current.PrevHash,
            CurrentHash: current.CurrentHash,
            Timestamp:   current.Timestamp,
        }
        // Prepend to maintain order
        chainData.Blocks = append([]BlockData{blockData}, chainData.Blocks...)
        current = current.PrevPointer
    }
    
    // Save to file
    data, err := json.MarshalIndent(chainData, "", "  ")
    if err != nil {
        return fmt.Errorf("error marshaling chain: %v", err)
    }
    
    return os.WriteFile(dataFile, data, 0644)
}

func (c *Chain) LoadChain() error {
    // Check if file exists
    data, err := os.ReadFile(dataFile)
    if err != nil {
        if os.IsNotExist(err) {
            fmt.Println("DEBUG: No existing blockchain found, starting fresh")
            return nil
        }
        return fmt.Errorf("error reading chain file: %v", err)
    }
    
    var chainData ChainData
    if err := json.Unmarshal(data, &chainData); err != nil {
        return fmt.Errorf("error unmarshaling chain: %v", err)
    }
    
    // Rebuild chain
    c.Head = nil
    c.BlockCount = 0
    
    // Add blocks in correct order
    for _, blockData := range chainData.Blocks {
        c.InsertBlock(blockData.Transactions)
    }
    
    fmt.Printf("DEBUG: Loaded %d blocks from storage\n", c.BlockCount)
    return nil
}