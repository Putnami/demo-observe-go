package pricing

import (
	"testing"

	"go.putnami.dev/protocol/features/spectest"
)

func TestSubtotal(t *testing.T) {
	spectest.Proves(t, "pricing/order-total", "the-subtotal-sums-the-lines", "two-lines-sum-to-their-total")
	got, err := Subtotal([]Line{{SKU: "a", UnitCents: 250, Quantity: 2}, {SKU: "b", UnitCents: 100, Quantity: 1}})
	if err != nil {
		t.Fatal(err)
	}
	if got != 600 {
		t.Fatalf("Subtotal = %d, want 600", got)
	}
}

func TestSubtotalRejectsNegative(t *testing.T) {
	spectest.Proves(t, "pricing/order-total", "a-negative-line-is-refused", "a-negative-unit-price-returns-err-negative")
	if _, err := Subtotal([]Line{{SKU: "a", UnitCents: -1, Quantity: 1}}); err != ErrNegative {
		t.Fatalf("err = %v, want ErrNegative", err)
	}
}

func TestWithTax(t *testing.T) {
	spectest.Proves(t, "pricing/order-total", "tax-rounds-half-up", "twenty-percent-of-ten-euros-is-two-euros")
	if got := WithTax(1000, 2000); got != 1200 {
		t.Fatalf("WithTax = %d, want 1200", got)
	}
}
