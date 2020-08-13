package main

import (
	"testing"
)

func TestListLength(t *testing.T) {
	if insertListElements(5).Len() != 5 {
		t.Error("Expected length of list to equal 5")
	}
}

func TestTableListLength(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 2},
		{10, 10},
	}

	for _, test := range tests {
		if output := insertListElements(test.input).Len(); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
