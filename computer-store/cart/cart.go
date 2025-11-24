package cart

import (
	"fmt"

	"zeller-challenge/pricing"
)

func CheckOut(rules []pricing.PricingRule) *Cart {

	// Default Prices
	var DefaultPrices = map[string]float64{
		"ipd": 549.99,
		"mbp": 1399.99,
		"atv": 109.50,
		"vga": 30.00,
	}

	return &Cart{
		rules:  rules,
		items:  []Item{},
		prices: DefaultPrices,
	}
}

// Adding item to cart. If price is 0, get default price.
func (c *Cart) Scan(item Item) {
	if item.Price == 0 {
		if price, ok := c.prices[item.SKU]; ok {
			item.Price = price
		}
	}

	c.items = append(c.items, item)
}

// Total after applying discount.
func (c *Cart) Total() float64 {
	var sub float64
	cartCounts := make(map[string]int)
	for _, item := range c.items {
		sub += item.Price
		cartCounts[item.SKU]++
		if _, ok := c.prices[item.SKU]; !ok {
			c.prices[item.SKU] = item.Price
		}
	}

	var totalDiscount float64
	for _, r := range c.rules {
		discount := r.Apply(cartCounts, c.prices)
		if discount < 0 {
			fmt.Printf("warning: negative discount: %v", discount)
			continue
		}
		totalDiscount += discount
	}

	total := sub - totalDiscount
	return total
}
