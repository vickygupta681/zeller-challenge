package pricing

func (r MultiplebuyRule) Apply(cart map[string]int, prices map[string]float64) float64 {
	SKUQuantity := cart[r.SKU]
	if SKUQuantity <= 0 || r.PickItems <= r.PayFor || r.PickItems <= 0 {
		return 0
	}
	groups := SKUQuantity / r.PickItems
	free := groups * (r.PickItems - r.PayFor)
	unitPrice := prices[r.SKU]
	return float64(free) * unitPrice
}
