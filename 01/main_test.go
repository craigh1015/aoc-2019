package main

import "testing"

func TestCalculateFuel(t *testing.T) {
	testCases := []struct {
		desc     string
		input    int
		expected int
	}{
		{desc: "ex1", input: 12, expected: 2},
		{desc: "ex2", input: 14, expected: 2},
		{desc: "ex2", input: 1969, expected: 654},
		{desc: "ex2", input: 100756, expected: 33583},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := calculateFuel(tC.input)
			if result != tC.expected {
				t.Fatalf("expected [%d] got [%d]", tC.expected, result)
			}
		})
	}
}

func TestCalculateFuelRecursive(t *testing.T) {
	testCases := []struct {
		desc     string
		input    int
		expected int
	}{
		{desc: "ex1", input: 12, expected: 2},
		{desc: "ex2", input: 14, expected: 2},
		{desc: "ex2", input: 1969, expected: 966},
		{desc: "ex2", input: 100756, expected: 50346},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := calculateFuelRecursive(tC.input)
			if result != tC.expected {
				t.Fatalf("expected [%d] got [%d]", tC.expected, result)
			}
		})
	}
}
