package gas

import (
	"fmt"
)

// Configurable values
var (
	NormalTxUSD   = 0.03
	BridgeTxUSD   = 0.05
)

// Calculate fee in DAC based on USD-fixed model:
func CalculateFee(usdPrice float64, txType string) (float64, error) {
	var usdFee float64

	switch txType {
	case "normal":
		usdFee = NormalTxUSD
	case "bridge":
		usdFee = BridgeTxUSD
	default:
		return 0, fmt.Errorf("unknown tx type: %s", txType)
	}

	// DAC price in USD is usdPrice
	if usdPrice <= 0 {
		return 0, fmt.Errorf("invalid DAC price")
	}

	// Fee in DAC = USD fee / DAC price
	return usdFee / usdPrice, nil
}
