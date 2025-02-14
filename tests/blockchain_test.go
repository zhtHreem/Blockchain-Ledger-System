package tests

import (
	"testing"
	"/internal/blockchain"
)

func TestBlockchain(t *testing.T) {
	// Test 1: Create and verify empty blockchain
	chain := blockchain.NewChain()
	chain.InsertBlock([]string{"Genesis Block"})
	chain.ListBlocks()
	if !chain.VerifyChain() {
		t.Error("Test 1 failed: Chain verification failed for genesis block")
	}
	
	// Test 2: Multiple insertions
	chain.InsertBlock([]string{"Alice pays Bob 50"})
	chain.InsertBlock([]string{"Bob pays Charlie 30", "Charlie pays David 10"})
	if !chain.VerifyChain() {
		t.Error("Test 2 failed: Chain verification failed after multiple insertions")
	}
}