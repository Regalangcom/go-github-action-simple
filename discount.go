package main

import "fmt"

// CustomerType menentukan jenis pelanggan
type CustomerType string

const (
	Regular CustomerType = "Regular"
	Member  CustomerType = "Member"
	VIP     CustomerType = "VIP"
)

// CalculateDiscount menghitung diskon berdasarkan jumlah pembelian dan jenis pelanggan
func CalculateDiscount(customerType CustomerType, amount float64) float64 {
	var discountRate float64

	switch customerType {
	case Regular:
		discountRate = 0.05 // 5% untuk pelanggan biasa
	case Member:
		discountRate = 0.1 // 10% untuk member
	case VIP:
		discountRate = 0.2 // 20% untuk VIP
	default:
		discountRate = 0.0
	}

	discount := amount * discountRate
	return discount
}

func main() {
	total := 500.0
	customer := Member
	discount := CalculateDiscount(customer, total)

	fmt.Printf("this is results : Customer Type: %s\nTotal: %.2f\nDiscount: %.2f\n", customer, total, discount)
}
