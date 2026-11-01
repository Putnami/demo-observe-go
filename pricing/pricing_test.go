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

func TestDiscount(t *testing.T) {
	cases := []struct{ cents, bp, want int64 }{
		{1000, 1000, 900},
		{1000, 0, 1000},
		{1000, 12000, 0},
	}
	for _, c := range cases {
		if got := Discount(c.cents, c.bp); got != c.want {
			t.Fatalf("Discount(%d, %d) = %d, want %d", c.cents, c.bp, got, c.want)
		}
	}
}
