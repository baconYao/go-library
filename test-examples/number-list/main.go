package main

import (
	"errors"
	"math"
	"sort"
)

// ExtractList ,
// Given a list of integers, return the mean (the average value), median (when sorted, the
// value in the middle position), and mode (the value that occurs most often; a hash map
// will be helpful here) of the list.
func ExtractList(numbers []int) (mean float64, median float64, mode int, err error) {
	numbersLength := len(numbers)
	if numbersLength == 0 {
		err = errors.New("Should not be empty slice")
		return 0.0, 0.0, 0, err
	}
	indexOfMedian := int(math.Floor(float64(numbersLength) / 2))
	statistic := make(map[int]int)
	mean = 0
	mode = 0
	mostNumber := 0

	for _, number := range numbers {
		mean += float64(number)
		if _, ok := statistic[number]; ok {
			statistic[number]++
		} else {
			statistic[number] = 1
		}
	}
	// count average value
	mean = mean / float64(numbersLength)

	// sort slice, this will change the original slice
	sort.Ints(numbers)
	if numbersLength%2 == 0 {
		median = float64((numbers[indexOfMedian] + numbers[indexOfMedian-1])) / 2.0

	} else {
		median = float64(numbers[indexOfMedian])
	}

	// get mode value by using map to statistic each number
	for key, value := range statistic {
		if value >= mostNumber {
			mostNumber = value
			mode = key
		}
	}
	return mean, median, mode, nil
}
