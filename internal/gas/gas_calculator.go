package gas

// DACPriceUSD = price of DACCoin in USD (will be dynamic later)
var DACPriceUSD float64 = 0.015 // placeholder

func CalculateGasUSD(txType string) float64 {
    if txType == "bridge" {
        return 0.05
    }
    return 0.03
}

// Convert USD fee → DACCoin fee
func FeeInDAC(usd float64) float64 {
    if DACPriceUSD == 0 {
        return 0
    }
    return usd / DACPriceUSD
}
