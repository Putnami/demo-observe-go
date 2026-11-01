package pricing

import (
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Format renders an amount in cents as a localized euro amount.
func Format(cents int64, tag language.Tag) string {
	amount := currency.EUR.Amount(float64(cents) / 100)
	return message.NewPrinter(tag).Sprint(currency.Symbol(amount))
}
