package main

import (
	"fmt"
	"github.com/QuantumDac/web4-architecture/pkg/gas"
)

func main() {
	fmt.Println("Web4 Gas Fee Demo")

	// Example DAC price
	dacPrice := 0.012 // 1.2 cents

	normalFee, _ := gas.CalculateFee(dacPrice, "normal")
	bridgeFee, _ := gas.CalculateFee(dacPrice, "bridge")

	fmt.Printf("DAC Price: $%.4f\n", dacPrice)
	fmt.Printf("Normal Tx Fee: %.4f DAC\n", normalFee)
	fmt.Printf("Bridge Tx Fee: %.4f DAC\n", bridgeFee)
}
