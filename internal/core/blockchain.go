package core

import (
    "time"
)

type Blockchain struct {
    Chain []*Block
}

func NewBlockchain() *Blockchain {
    genesis := &Block{
        Height:    0,
        PrevHash:  "",
        Hash:      "GENESIS",
        Timestamp: time.Now().Unix(),
    }
    return &Blockchain{Chain: []*Block{genesis}}
}

func (bc *Blockchain) AddBlock(txs []Transaction) *Block {
    prev := bc.Chain[len(bc.Chain)-1]

    block := &Block{
        Height:       uint64(len(bc.Chain)),
        PrevHash:     prev.Hash,
        Hash:         "BLOCK_HASH_PLACEHOLDER",
        Timestamp:    time.Now().Unix(),
        Transactions: txs,
    }

    bc.Chain = append(bc.Chain, block)
    return block
}
