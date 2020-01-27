package main

import (
	"strings"
	"testing"
)

func TestReadPosition(t *testing.T) {
	testCases := []struct {
		desc     string
		tokens   []string
		input    int
		expected int
	}{
		{desc: "ex1", tokens: []string{"1", "0"}, input: 1, expected: 0},
		{desc: "ex1", tokens: []string{"1", "2"}, input: 0, expected: 1},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := readPosition(tC.tokens, tC.input)
			if err != nil {
				t.Error(err)
			}
			if result != tC.expected {
				t.Fatalf("expected [%d] got [%d]", tC.expected, result)
				// assert.Equal(t, tC.expected, result)
			}
		})
	}
}

func TestReadRelative(t *testing.T) {
	testCases := []struct {
		desc     string
		tokens   []string
		input    int
		expected int
	}{
		{desc: "ex1", tokens: []string{"1", "0"}, input: 1, expected: 1},
		{desc: "ex1", tokens: []string{"1", "2"}, input: 0, expected: 2},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, err := readRelative(tC.tokens, tC.input)
			if err != nil {
				t.Error(err)
			}
			if result != tC.expected {
				t.Fatalf("expected [%d] got [%d]", tC.expected, result)
				// assert.Equal(t, tC.expected, result)
			}
		})
	}
}

func TestRun(t *testing.T) {
	// t.Skip()
	testCases := []struct {
		desc     string
		program  string
		noun     int
		verb     int
		expected string
	}{
		{desc: "ex1", program: "1,0,0,0,99", noun: 0, verb: 0, expected: "2,0,0,0,99"},
		{desc: "ex2", program: "2,3,0,3,99", noun: 3, verb: 0, expected: "2,3,0,6,99"},
		{desc: "ex3", program: "2,4,4,5,99,0", noun: 4, verb: 4, expected: "2,4,4,5,99,9801"},
		{desc: "ex4", program: "1,1,1,4,99,5,6,0,99", noun: 1, verb: 1, expected: "30,1,1,4,2,5,6,0,99"},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := strings.Join(run(tC.program, tC.noun, tC.verb), ",")
			if result != tC.expected {
				t.Fatalf("expected [%s] got [%s]", tC.expected, result)
				// assert.Equal(t, tC.expected, result)
			}
		})
	}
}
