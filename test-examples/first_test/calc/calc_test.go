package calculate

import (
	"testing"
)

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 5},
		{-1, 2},
		{0, 3},
		{-5, -2},
		{99999, 100001}, // Error case
	}

	for _, test := range tests {
		if output := Calculate3(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
