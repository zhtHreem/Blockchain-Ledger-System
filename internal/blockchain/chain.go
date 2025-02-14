

package blockchain

import "fmt"

type Chain struct {
    Head *Block
    BlockCount int
}

func NewChain() *Chain {
    return &Chain{
        Head: nil,
        BlockCount: 0,
    }
}

func (c *Chain) InsertBlock(transactions []string) {
    newBlock := NewBlock(transactions, c.Head)
    c.Head = newBlock
    c.BlockCount++
}

func (c *Chain) ListBlocks() {
    if c.Head == nil {
        fmt.Println("Blockchain is empty")
        return
    }

    current := c.Head
    blockNum := c.BlockCount

    for current != nil {
        fmt.Printf("\nBlock %d:\n", blockNum)
        fmt.Printf("Transactions: %v\n", current.Transactions)
        fmt.Printf("Hash: %s\n", current.CurrentHash[:16]) // Show only first 16 chars of hash for readability
        if current.PrevHash != "" {
            fmt.Printf("Previous Block's Hash: %s\n", current.PrevHash[:16])
        }
        fmt.Println("------------------------")
        
        current = current.PrevPointer
        blockNum--
    }
    
    fmt.Printf("\nTotal blocks: %d\n", c.BlockCount)
}


func (c *Chain) ChangeBlock(oldTrans string, newTrans string) {
	current := c.Head
	
	for current != nil {
		for i, trans := range current.Transactions {
			if trans == oldTrans {
				current.Transactions[i] = newTrans
				return
			}
		}
		current = current.PrevPointer
	}
}

func (c *Chain) VerifyChain() bool {
	current := c.Head
	
	for current != nil {
		expectedHash := CalculateHash(current)
		
		if expectedHash != current.CurrentHash {
			fmt.Println("Block chain is compromised (hash mismatch)")
			return false
		}
		
		if current.PrevPointer != nil {
			if current.PrevHash != current.PrevPointer.CurrentHash {
				fmt.Println("Block chain is compromised (prevHash mismatch)")
				return false
			}
		}
		
		current = current.PrevPointer
	}
	
	fmt.Println("Block chain is unchanged")
	return true
}







