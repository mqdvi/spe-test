package main

import (
	"reflect"
	"testing"
)

func TestNarcissistic(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{
			n:        1634,
			expected: true,
		},
		{
			n:        153,
			expected: true,
		},
		{
			n:        111,
			expected: false,
		},
		{
			n:        9474,
			expected: true,
		},
		{
			n:        8208,
			expected: true,
		},
	}

	for _, testCase := range testCases {
		t.Run("NarcissisticNumber", func(t *testing.T) {
			result := NarcissisticNumber(testCase.n)
			if result != testCase.expected {
				t.Errorf("Expected %v for input %d, but got %v", testCase.expected, testCase.n, result)
			}
		})
	}
}

func TestParityOutlier(t *testing.T) {
	testCases := []struct {
		numbers  []int
		expected interface{}
	}{
		{
			numbers:  []int{2, 4, 0, 100, 4, 11, 2602, 36},
			expected: 11,
		},
		{
			numbers:  []int{160, 3, 1719, 19, 11, 13, -21},
			expected: 160,
		},
		{
			numbers:  []int{11, 13, 15, 19, 9, 13, -21},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run("ParityOutlier", func(t *testing.T) {
			result := ParityOutlier(testCase.numbers)
			if result != testCase.expected {
				t.Errorf("Expected %v for input %v, but got %v", testCase.expected, testCase.numbers, result)
			}
		})
	}
}

func TestFindNeedleInHaystack(t *testing.T) {
	testCases := []struct {
		data     []string
		needle   string
		expected int
	}{
		{
			data:     []string{"red", "blue", "yellow", "black", "grey"},
			needle:   "blue",
			expected: 1,
		},
		{
			data:     []string{"apple", "banana", "orange"},
			needle:   "orange",
			expected: 2,
		},
		{
			data:     []string{"dog", "cat", "fish"},
			needle:   "bird",
			expected: -1,
		},
	}

	for _, testCase := range testCases {
		t.Run("FindNeedleInHaystack", func(t *testing.T) {
			result := FindNeedleInHaystack(testCase.data, testCase.needle)
			if result != testCase.expected {
				t.Errorf("Expected %d for input %v with needle %s, but got %d", testCase.expected, testCase.data, testCase.needle, result)
			}
		})
	}
}

func TestBlueOceanReverse(t *testing.T) {
	testCases := []struct {
		list1    []int
		list2    []int
		expected []int
	}{
		{
			list1:    []int{1, 2, 3, 4, 6, 10},
			list2:    []int{1},
			expected: []int{2, 3, 4, 6, 10}},
		{
			list1:    []int{1, 5, 5, 5, 5, 3},
			list2:    []int{5},
			expected: []int{1, 3}},
		{
			list1:    []int{1, 2, 3},
			list2:    []int{4},
			expected: []int{1, 2, 3}},
		{
			list1:    []int{},
			list2:    []int{1, 2, 3},
			expected: []int{}},
		{
			list1:    []int{1, 2, 3},
			list2:    []int{},
			expected: []int{1, 2, 3}},
	}

	for _, testCase := range testCases {
		t.Run("BlueOceanReverse", func(t *testing.T) {
			result := BlueOceanReverse(testCase.list1, testCase.list2)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Expected %v for input %v and %v, but got %v", testCase.expected, testCase.list1, testCase.list2, result)
			}
		})
	}
}
