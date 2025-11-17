package core

type Transaction struct {
    From   string
    To     string
    Amount float64
    GasFee float64
    Type   string  // "normal", "bridge"
}
