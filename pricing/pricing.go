// Package pricing computes order totals in cents.
package pricing

import "errors"

// ErrNegative reports a negative quantity or price.
var ErrNegative = errors.New("pricing: negative value")

// Line is one order line.
type Line struct {
	SKU       string
	UnitCents int64
	Quantity  int64
}

// Subtotal returns the sum of the lines, before tax.
func Subtotal(lines []Line) (int64, error) {
	var total int64
	for _, l := range lines {
		if l.UnitCents < 0 || l.Quantity < 0 {
			return 0, ErrNegative
		}
		total += l.UnitCents * l.Quantity
	}
	return total, nil
}

// WithTax adds a tax rate, in basis points, rounded half up.
func WithTax(cents int64, basisPoints int64) int64 {
	return cents + (cents*basisPoints+5000)/10000
}
