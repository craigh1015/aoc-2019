package main

import (
	"testing"
)

func TestReadPosition(t *testing.T) {
	testCases := []struct {
		desc     string
		tokens   []int
		input    int
		expected int
	}{
		{desc: "ex1", tokens: []int{1, 0}, input: 1, expected: 0},
		{desc: "ex1", tokens: []int{1, 2}, input: 0, expected: 1},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := readPosition(tC.tokens, tC.input)
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
		tokens   []int
		input    int
		expected int
	}{
		{desc: "ex1", tokens: []int{1, 0}, input: 1, expected: 1},
		{desc: "ex1", tokens: []int{1, 2}, input: 0, expected: 2},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := readRelative(tC.tokens, tC.input)
			if result != tC.expected {
				t.Fatalf("expected [%d] got [%d]", tC.expected, result)
				// assert.Equal(t, tC.expected, result)
			}
		})
	}
}

func TestRead(t *testing.T) {
	testCases := []struct {
		desc   string
		tokens []int
		modes  int
		pc     int
		offset int
		result int
	}{
		{
			"ex1", []int{1, 1, 1, 1, 99}, 0, 0, 1, 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := read(tC.tokens, tC.modes, tC.pc, tC.offset)
			if result != tC.result {
				t.Fatalf("expected [%d] got [%d]", tC.result, result)
			}
		})
	}
}

func TestDecomposeCommand(t *testing.T) {
	testCases := []struct {
		desc    string
		command int
		opCode  int
		modes   int
	}{
		{"ex1", 1, 1, 0},
		{"ex2", 99, 99, 0},
		{"ex3", 101, 1, 1},
		{"ex4", 1002, 2, 10},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			opCode, modes := decomposeCommand(tC.command)
			if opCode != tC.opCode {
				t.Fatalf("expected [%d] got [%d]", tC.opCode, opCode)
			}
			if modes != tC.modes {
				t.Fatalf("expected [%d] got [%d]", tC.modes, modes)
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
		expected []int
	}{
		{desc: "ex1", program: "1,0,0,0,99", noun: 0, verb: 0, expected: []int{2, 0, 0, 0, 99}},
		{desc: "ex2", program: "2,3,0,3,99", noun: 3, verb: 0, expected: []int{2, 3, 0, 6, 99}},
		{desc: "ex3", program: "2,4,4,5,99,0", noun: 4, verb: 4, expected: []int{2, 4, 4, 5, 99, 9801}},
		{desc: "ex4", program: "1,1,1,4,99,5,6,0,99", noun: 1, verb: 1, expected: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		// {desc: "ex5", program: "1002,4,3,4,33", noun: 4, verb: 3, expected: []int{1002, 4, 3, 4, 99}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := run(tC.program, tC.noun, tC.verb)
			for i, val := range result {
				if val != tC.expected[i] {
					t.Fatalf("expected [%v] got [%v]", tC.expected, result)
				}
			}
		})
	}
}
