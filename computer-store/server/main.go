package main

import (
	"fmt"

	"zeller-challenge/cart"
	"zeller-challenge/pricing"
)

func main() {
	rules := []pricing.PricingRule{
		pricing.MultiplebuyRule{SKU: "atv", PickItems: 3, PayFor: 2},
		pricing.BulkDiscountRule{SKU: "ipd", Threshold: 4, DiscountedPrice: 499.99},
	}

	// cart 1
	c1 := cart.CheckOut(rules)
	c1.Scan(cart.Item{SKU: "atv"})
	c1.Scan(cart.Item{SKU: "atv"})
	c1.Scan(cart.Item{SKU: "atv"})
	c1.Scan(cart.Item{SKU: "vga"})
	fmt.Printf("Scenario 1 total: $%.2f", c1.Total())

	// cart 2
	c2 := cart.CheckOut(rules)
	c2.Scan(cart.Item{SKU: "atv"})
	c2.Scan(cart.Item{SKU: "ipd"})
	c2.Scan(cart.Item{SKU: "ipd"})
	c2.Scan(cart.Item{SKU: "atv"})
	c2.Scan(cart.Item{SKU: "ipd"})
	c2.Scan(cart.Item{SKU: "ipd"})
	c2.Scan(cart.Item{SKU: "ipd"})
	fmt.Printf("\nScenario 2 total: $%.2f", c2.Total())

	// Cart 3
	c3 := cart.CheckOut(rules)
	for i := 0; i < 5; i++ {
		c3.Scan(cart.Item{SKU: "ipd"})
	}
	fmt.Printf("\nScenario 3 total: $%.2f", c3.Total())
}
