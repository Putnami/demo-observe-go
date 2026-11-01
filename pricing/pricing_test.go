package pricing

import "testing"

func TestSubtotal(t *testing.T) {
	got, err := Subtotal([]Line{{SKU: "a", UnitCents: 250, Quantity: 2}, {SKU: "b", UnitCents: 100, Quantity: 1}})
	if err != nil {
		t.Fatal(err)
	}
	if got != 600 {
		t.Fatalf("Subtotal = %d, want 600", got)
	}
}

func TestSubtotalRejectsNegative(t *testing.T) {
	if _, err := Subtotal([]Line{{SKU: "a", UnitCents: -1, Quantity: 1}}); err != ErrNegative {
		t.Fatalf("err = %v, want ErrNegative", err)
	}
}

func TestWithTax(t *testing.T) {
	if got := WithTax(1000, 2000); got != 1200 {
		t.Fatalf("WithTax = %d, want 1200", got)
	}
}
