package main

import (
	"strconv"
	"testing"
)

/*
Two adjacent digits are the same (like 22 in 122345).
Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
*/

func TestPair(t *testing.T) {
	testCases := []struct {
		desc   string
		value  int
		result bool
	}{
		{"ex1", 123456, false},
		{"ex2", 123455, true},
		{"ex3", 112345, true},
		{"ex4", 111234, false},
		{"ex5", 111244, true},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := hasPair(strconv.Itoa(tC.value))
			if tC.result != result {
				t.Fatalf("Expected [%v] got [%v]", tC.result, result)
			}
		})
	}
}

func TestDoubleDigits(t *testing.T) {
	testCases := []struct {
		desc   string
		value  int
		result bool
	}{
		{"ex1", 123456, false},
		{"ex2", 123455, true},
		{"ex3", 112345, true},
		{"ex4", 111234, true},
		{"ex5", 111244, true},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := hasDoubleDigits(strconv.Itoa(tC.value))
			if tC.result != result {
				t.Fatalf("Expected [%v] got [%v]", tC.result, result)
			}
		})
	}
}

func TestIncrease(t *testing.T) {
	testCases := []struct {
		desc   string
		value  int
		result bool
	}{
		{"ex1", 123456, true},
		{"ex2", 123454, false},
		{"ex3", 212345, false},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := increases(strconv.Itoa(tC.value))
			if tC.result != result {
				t.Fatalf("Expected [%v] got [%v]", tC.result, result)
			}
		})
	}
}
