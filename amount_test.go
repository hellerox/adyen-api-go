package adyen

import (
	"testing"
)

func TestNewAmount(t *testing.T) {
	cases := []struct {
		name     string
		currency string
		amount   float32
		expected Amount
	}{
		{
			name:     "Test EUR currency",
			currency: "EUR",
			amount:   10.50,
			expected: Amount{Currency: "EUR", Value: 1050},
		},
		{
			name:     "Test EUR currency, zero case",
			currency: "EUR",
			amount:   0,
			expected: Amount{Currency: "EUR", Value: 0},
		},
		{
			name:     "Test unknown (UKN) currency, default should be used",
			currency: "UKN",
			amount:   10.60,
			expected: Amount{Currency: "UKN", Value: 1060},
		},
		{
			name:     "Test CVE currency with zero decimal adjustment",
			currency: "CVE",
			amount:   150,
			expected: Amount{Currency: "CVE", Value: 150},
		},
		{
			name:     "Test BHD currency with 3 decimal adjustment points",
			currency: "BHD",
			amount:   150.050,
			expected: Amount{Currency: "BHD", Value: 150050},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			a := NewAmount(c.currency, c.amount)

			equals(t, c.expected, *a)
		})
	}
}
