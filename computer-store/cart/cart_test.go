package cart

import (
	"testing"

	"zeller-challenge/pricing"
)

func TestScenario1_ThreeForTwo(t *testing.T) {
	rules := []pricing.PricingRule{
		pricing.MultiplebuyRule{SKU: "atv", PickItems: 3, PayFor: 2},
		pricing.BulkDiscountRule{SKU: "ipd", Threshold: 4, DiscountedPrice: 499.99},
	}
	co := CheckOut(rules)
	co.Scan(Item{SKU: "atv"})
	co.Scan(Item{SKU: "atv"})
	co.Scan(Item{SKU: "atv"})
	co.Scan(Item{SKU: "vga"})
	exp := 249.00
	if got := co.Total(); got != exp {
		t.Fatalf("expected %.2f, got %.2f", exp, got)
	}
}

func TestScenario2_BulkAndThreeForTwo(t *testing.T) {
	rules := []pricing.PricingRule{
		pricing.MultiplebuyRule{SKU: "atv", PickItems: 3, PayFor: 2},
		pricing.BulkDiscountRule{SKU: "ipd", Threshold: 4, DiscountedPrice: 499.99},
	}
	co := CheckOut(rules)
	co.Scan(Item{SKU: "atv"})
	co.Scan(Item{SKU: "ipd"})
	co.Scan(Item{SKU: "ipd"})
	co.Scan(Item{SKU: "atv"})
	co.Scan(Item{SKU: "ipd"})
	co.Scan(Item{SKU: "ipd"})
	co.Scan(Item{SKU: "ipd"})
	exp := 2718.95
	if got := co.Total(); got != exp {
		t.Fatalf("expected %.2f, got %.2f", exp, got)
	}
}

func TestBulkDiscountFiveIPD(t *testing.T) {
	rules := []pricing.PricingRule{
		pricing.BulkDiscountRule{SKU: "ipd", Threshold: 4, DiscountedPrice: 499.99},
	}
	co := CheckOut(rules)
	for i := 0; i < 5; i++ {
		co.Scan(Item{SKU: "ipd"})
	}
	exp := 5 * 499.99
	if got := co.Total(); got != exp {
		t.Fatalf("expected %.2f, got %.2f", exp, got)
	}
}
