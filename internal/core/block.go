package core

type Block struct {
    Height        uint64
    PrevHash      string
    Hash          string
    Timestamp     int64
    Transactions  []Transaction
}
