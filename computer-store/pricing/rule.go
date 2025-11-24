package pricing

// Interface that is defining rules applying to the cart for skus pricing.
type PricingRule interface {
	Apply(cart map[string]int, prices map[string]float64) float64
}
