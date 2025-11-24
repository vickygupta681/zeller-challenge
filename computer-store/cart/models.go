package cart

import "zeller-challenge/pricing"

type Item struct {
	SKU   string
	Name  string
	Price float64
}

type Cart struct {
	rules  []pricing.PricingRule
	items  []Item
	prices map[string]float64
}
