package gas

import (
	"errors"
	"math/big"
	"strings"
)

// CalculateGasInDAC returns gas amount expressed in DAC (as decimal string with 8 decimals).
// gasUSD and dacPriceUSD are decimal strings like "0.03" or "0.01".
// It uses big.Rat for exact rational arithmetic and formats result with decimals decimals.
func CalculateGasInDAC(gasUSDStr, dacPriceUSDStr string, decimals int) (string, error) {
	gasUSDStr = strings.TrimSpace(gasUSDStr)
	dacPriceUSDStr = strings.TrimSpace(dacPriceUSDStr)
	if gasUSDStr == "" || dacPriceUSDStr == "" {
		return "", errors.New("empty gasUSD or dacPriceUSD")
	}

	gasRat, ok := new(big.Rat).SetString(gasUSDStr)
	if !ok {
		return "", errors.New("invalid gasUSD string")
	}
	priceRat, ok := new(big.Rat).SetString(dacPriceUSDStr)
	if !ok {
		return "", errors.New("invalid dacPriceUSD string")
	}

	// gasInDac = gasUSD / dacPriceUSD
	if priceRat.Sign() == 0 {
		return "", errors.New("dac price cannot be zero")
	}
	gasInDac := new(big.Rat).Quo(gasRat, priceRat)

	// Format as decimal string with rounding up to 'decimals'
	str := CeilToDecimals(gasInDac, decimals)
	return str, nil
}

// CeilToDecimals returns decimal string of rat rounded UP (ceiling) at 'decimals' places.
// Example: CeilToDecimals(1.234567891, 8) -> "1.23456789" (if exact) or "1.23456790" if need ceil.
func CeilToDecimals(r *big.Rat, decimals int) string {
	// Multiply by 10^decimals and perform ceiling on numerator/denominator
	mul := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil) // 10^decimals

	// scaled = r * mul
	num := new(big.Int).Mul(r.Num(), mul)
	den := new(big.Int).Set(r.Denom())

	// integer division and remainder
	intPart := new(big.Int).Quo(num, den)
	rem := new(big.Int).Rem(num, den)

	// if remainder > 0 then ceil => intPart + 1
	if rem.Sign() > 0 {
		intPart.Add(intPart, big.NewInt(1))
	}

	// Convert intPart back to string with decimal point
	intStr := intPart.String()
	// ensure length > decimals to place decimal point
	if len(intStr) <= decimals {
		// pad with leading zeros
		pad := strings.Repeat("0", decimals+1-len(intStr))
		intStr = pad + intStr
	}
	// split
	intLen := len(intStr)
	intPartPart := intStr[:intLen-decimals]
	fracPart := intStr[intLen-decimals:]
	// trim trailing zeros in fraction? we keep full decimals for consistency
	return intPartPart + "." + fracPart
}
