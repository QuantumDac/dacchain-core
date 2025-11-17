package tests

import (
	"testing"

	"github.com/QuantumDac/web4-architecture/internal/gas"
)

func TestCalculateGasInDAC(t *testing.T) {
	tests := []struct {
		gasUsd string
		price  string
		want   string
	}{
		{"0.03", "0.01", "3.00000000"},
		{"0.03", "0.015", "2.00000000"},
		{"0.05", "0.02", "2.50000000"}, // ceiling will produce 2.50000000
	}
	for _, tt := range tests {
		got, err := gas.CalculateGasInDAC(tt.gasUsd, tt.price, 8)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != tt.want {
			t.Fatalf("gas(%s,%s) = %s, want %s", tt.gasUsd, tt.price, got, tt.want)
		}
	}
}
