package pricing

// If quantity is greater than Threshold, reduce the unit price
type BulkDiscountRule struct {
	SKU             string
	Threshold       int
	DiscountedPrice float64
}

func (r BulkDiscountRule) Apply(cart map[string]int, prices map[string]float64) float64 {
	count := cart[r.SKU]
	if count <= r.Threshold {
		return 0
	}

	originalPrice := prices[r.SKU]
	d := (originalPrice - r.DiscountedPrice) * float64(count)
	if d < 0 {
		return 0
	}
	return d
}
