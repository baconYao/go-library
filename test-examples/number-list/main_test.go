package main

import "testing"

func TestExtractList(t *testing.T) {

	// helper function to check result
	checkExpection := func(t *testing.T, numbers []int, mean, wantMean, median, wantMedian float64, mode, wantMode int) {
		t.Helper()
		if mean != wantMean {
			t.Errorf("got %.3f wantMean %.3f gievn, %v", mean, wantMean, numbers)
		}
		if median != wantMedian {
			t.Errorf("got %.3f wantMedian %.3f gievn, %v", median, wantMedian, numbers)
		}
		if mode != wantMode {
			t.Errorf("got %d wantMode %d gievn, %v", mode, wantMode, numbers)
		}
	}

	t.Run("Test normal case", func(t *testing.T) {
		numbers := []int{40, 21, 59, 13, 44, 9, 9, 9, 21}
		mean, median, mode, _ := ExtractList(numbers)
		checkExpection(t, numbers, mean, 25.0, median, 21.0, mode, 9)
	})

	t.Run("Test negative number", func(t *testing.T) {
		numbers := []int{-11, -22, 0, -11, 3}
		mean, median, mode, _ := ExtractList(numbers)
		checkExpection(t, numbers, mean, -8.2, median, -11, mode, -11)
	})

	t.Run("Get median when even number", func(t *testing.T) {
		numbers := []int{99, 22, 39, 48, 22, 22}
		mean, median, mode, _ := ExtractList(numbers)
		checkExpection(t, numbers, mean, 42, median, 30.5, mode, 22)
	})

	t.Run("Test only one value", func(t *testing.T) {
		numbers := []int{3}
		mean, median, mode, _ := ExtractList(numbers)
		checkExpection(t, numbers, mean, 3.0, median, 3.0, mode, 3)
	})

	t.Run("Test only one value and it's 0", func(t *testing.T) {
		numbers := []int{0}
		mean, median, mode, _ := ExtractList(numbers)
		checkExpection(t, numbers, mean, 0.0, median, 0.0, mode, 0)
	})

	t.Run("Test empty slice and error string", func(t *testing.T) {
		numbers := []int{}
		_, _, _, err := ExtractList(numbers)
		if err.Error() != "Should not be empty slice" {
			t.Fatal("message error")
		}
	})

	// FIXME
	// t.Run("Get multiple modes", func(t *testing.T) {
	// 	numbers := []int{11, 22, 11, 22}
	// 	mean, median, mode := ExtractList(numbers)
	// 	checkExpection(t, numbers, mean, 16.5, median, 16.5, mode, []int{11, 22})
	// })

}
