package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDiscount(t *testing.T) {
	tests := []struct {
		customerType CustomerType
		amount       float64
		expected     float64
	}{
		{Regular, 100, 5},   // 5% dari 100
		{Member, 200, 20},   // 10% dari 200
		{VIP, 500, 100},     // 20% dari 500
		{"Unknown", 100, 0}, // Tidak ada diskon untuk tipe tidak dikenal
	}

	for _, test := range tests {
		result := CalculateDiscount(test.customerType, test.amount)
		assert.Equal(t, test.expected, result, "Discount calculation should be correct")
	}
}
