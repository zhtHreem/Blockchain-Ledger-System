
package blockchain

import (
    "crypto/sha256"
    "encoding/hex"
    "strings"
)

func CalculateHash(inputBlock *Block) string {
    transactionStr := strings.Join(inputBlock.Transactions, ",")
    input := transactionStr + inputBlock.PrevHash
    hash := sha256.New()
    hash.Write([]byte(input))
    return hex.EncodeToString(hash.Sum(nil))
}