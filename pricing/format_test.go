package pricing

import (
	"testing"

	"golang.org/x/text/language"
)

func TestFormat(t *testing.T) {
	if got := Format(1250, language.English); got != "€ 12.50" {
		t.Fatalf("Format = %q, want %q", got, "€ 12.50")
	}
}
