package main

import (
    "fmt"
    "web4-architecture/internal/core"
    "web4-architecture/internal/gas"
)

func main() {
    fmt.Println("🚀 Starting Web4 Node...")

    bc := core.NewBlockchain()

    tx := core.Transaction{
        From: "Alice",
        To:   "Bob",
        Amount: 10,
        Type: "normal",
    }

    usd := gas.CalculateGasUSD(tx.Type)
    tx.GasFee = gas.FeeInDAC(usd)

    block := bc.AddBlock([]core.Transaction{tx})

    fmt.Println("Block added:", block.Height)
    fmt.Println("Tx Gas Fee (in DAC):", tx.GasFee)
}
