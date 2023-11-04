package shorter

import (
	"testing"
)

func TestURLToID(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"https://www.example.com", "y5xd_ubeTb"},
		{"https://www.example.org", "x07zFQGz3A"},
		{"https://www.example.net", "BvToCBk1Cr"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := URLToID(test.input)
			if actual != test.expected {
				t.Errorf("URLToID(%s) = %s; want %s", test.input, actual, test.expected)
			}
		})
	}
}
