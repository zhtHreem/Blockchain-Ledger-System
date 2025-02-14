package main

import (
    "blockchain/internal/blockchain"
    "flag"
    "fmt"
    "os"
)

const version = "1.0.0"

// Create a global chain instance
var chain = blockchain.NewChain()

func main() {
    initCmd := flag.NewFlagSet("init", flag.ExitOnError)
    addCmd := flag.NewFlagSet("add", flag.ExitOnError)
    verifyCmd := flag.NewFlagSet("verify", flag.ExitOnError)
    listCmd := flag.NewFlagSet("list", flag.ExitOnError)

    fromAddr := addCmd.String("from", "", "From address")
    toAddr := addCmd.String("to", "", "To address")
    amount := addCmd.Float64("amount", 0.0, "Transaction amount")

    if len(os.Args) < 2 {
        printUsage()
        os.Exit(1)
    }

    fmt.Printf("DEBUG: Command received: %s\n", os.Args[1])

    switch os.Args[1] {
    case "init":
        handleInit(initCmd)
    case "add":
        handleAdd(addCmd, fromAddr, toAddr, amount)
    case "verify":
        handleVerify(verifyCmd)
    case "list":
        handleList(listCmd)
    case "version":
        fmt.Printf("Blockchain Ledger System v%s\n", version)
    case "help":
        printUsage()
    default:
        fmt.Printf("Unknown command: %s\n", os.Args[1])
        printUsage()
        os.Exit(1)
    }
}

func printUsage() {
    fmt.Println("Usage:")
    fmt.Println("  blockchain-ledger <command> [arguments]")
    fmt.Println("\nCommands:")
    fmt.Println("  init                Initialize a new blockchain")
    fmt.Println("  add -from <addr> -to <addr> -amount <value>   Add a new transaction")
    fmt.Println("  verify              Verify the blockchain integrity")
    fmt.Println("  list                List all blocks in the chain")
    fmt.Println("  version             Show version information")
    fmt.Println("  help                Show this help message")
}

// File: cmd/ledger/main.go
// Update these functions:

func handleInit(cmd *flag.FlagSet) {
    cmd.Parse(os.Args[2:])
    chain = blockchain.NewChain()
    if err := chain.SaveChain(); err != nil {
        fmt.Printf("Error saving chain: %v\n", err)
        return
    }
    fmt.Println("Initialized new blockchain")
}

func handleList(cmd *flag.FlagSet) {
    cmd.Parse(os.Args[2:])
    
    if err := chain.LoadChain(); err != nil {
        fmt.Printf("Error loading blockchain: %v\n", err)
        return
    }
    
    fmt.Println("\nBlockchain contents:")
    fmt.Println("==================")
    chain.ListBlocks()
}

func handleAdd(cmd *flag.FlagSet, from, to *string, amount *float64) {
    cmd.Parse(os.Args[2:])
    
    if *from == "" || *to == "" || *amount == 0.0 {
        fmt.Println("Error: --from, --to, and --amount are required")
        cmd.Usage()
        os.Exit(1)
    }
    
    if err := chain.LoadChain(); err != nil {
        fmt.Printf("Error loading blockchain: %v\n", err)
        return
    }
    
    transaction := fmt.Sprintf("%s->%s:%.2f", *from, *to, *amount)
    chain.InsertBlock([]string{transaction})
    
    if err := chain.SaveChain(); err != nil {
        fmt.Printf("Error saving blockchain: %v\n", err)
        return
    }
    
    fmt.Printf("Successfully added transaction: %s\n", transaction)
    fmt.Printf("Current blockchain size: %d blocks\n", chain.BlockCount)
}

func handleVerify(cmd *flag.FlagSet) {
    cmd.Parse(os.Args[2:])
    isValid := chain.VerifyChain()
    if isValid {
        fmt.Println("Blockchain verification successful: Chain is valid")
    }
}