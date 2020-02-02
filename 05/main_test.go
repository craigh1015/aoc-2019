package main

import (
	"testing"
)

func TestRead(t *testing.T) {
	testCases := []struct {
		desc   string
		tokens []int
		modes  int
		pc     int
		offset int
		result int
	}{
		{"pos01", []int{1, 5, 6, 7, 99, 17, 18, 19}, 0, 0, 1, 17},
		{"pos02", []int{1, 5, 6, 7, 99, 17, 18, 19}, 0, 0, 2, 18},
		{"pos03", []int{1, 5, 6, 7, 99, 17, 18, 19}, 0, 0, 3, 19},
		{"dir01", []int{1, 5, 6, 7, 99, 17, 18, 19}, 111, 0, 1, 5},
		{"dir02", []int{1, 5, 6, 7, 99, 17, 18, 19}, 111, 0, 2, 6},
		{"dir03", []int{1, 5, 6, 7, 99, 17, 18, 19}, 111, 0, 3, 7},
		{"dir04", []int{1, 5, 6, 7, 99, 17, 18, 19}, 110, 0, 1, 17},
		{"dir05", []int{1, 5, 6, 7, 99, 17, 18, 19}, 110, 0, 2, 6},
		{"dir06", []int{1, 5, 6, 7, 99, 17, 18, 19}, 110, 0, 3, 7},
		{"dir07", []int{1, 5, 6, 7, 99, 17, 18, 19}, 101, 0, 1, 5},
		{"dir08", []int{1, 5, 6, 7, 99, 17, 18, 19}, 101, 0, 2, 18},
		{"dir09", []int{1, 5, 6, 7, 99, 17, 18, 19}, 101, 0, 3, 7},
		{"dir10", []int{1, 5, 6, 7, 99, 17, 18, 19}, 11, 0, 1, 5},
		{"dir11", []int{1, 5, 6, 7, 99, 17, 18, 19}, 11, 0, 2, 6},
		{"dir12", []int{1, 5, 6, 7, 99, 17, 18, 19}, 11, 0, 3, 19},
		{"dir13", []int{1, 5, 6, 7, 99, 17, 18, 19}, 10, 0, 1, 17},
		{"dir14", []int{1, 5, 6, 7, 99, 17, 18, 19}, 10, 0, 2, 6},
		{"dir15", []int{1, 5, 6, 7, 99, 17, 18, 19}, 10, 0, 3, 19},
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
		desc    string
		program string
		input   int
		result  []int
		output  int
	}{
		{desc: "ex01", program: "1,0,0,0,99", result: []int{2, 0, 0, 0, 99}},
		{desc: "ex02", program: "2,3,0,3,99", result: []int{2, 3, 0, 6, 99}},
		{desc: "ex03", program: "2,4,4,5,99,0", result: []int{2, 4, 4, 5, 99, 9801}},
		{desc: "ex04", program: "1,1,1,4,99,5,6,0,99", result: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{desc: "ex05", program: "1002,4,3,4,33", result: []int{1002, 4, 3, 4, 99}},
		{desc: "ex06", program: "1101,100,-1,4,0", result: []int{1101, 100, -1, 4, 99}},
		{desc: "ex07", program: "3,5,4,5,99,0", result: []int{3, 5, 4, 5, 99, 11}, input: 11, output: 11},
		{desc: "ex08", program: "3,5,104,5,99,0", result: []int{3, 5, 104, 5, 99, 11}, input: 11, output: 5},

		{desc: "ex09", program: "3,9,8,9,10,9,4,9,99,-1,8", result: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 1, 8}, input: 8, output: 1},
		{desc: "ex10", program: "3,9,8,9,10,9,4,9,99,-1,8", result: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 0, 8}, input: 9, output: 0},
		{desc: "ex11", program: "3,9,7,9,10,9,4,9,99,-1,8", result: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, 1, 8}, input: 7, output: 1},
		{desc: "ex12", program: "3,9,7,9,10,9,4,9,99,-1,8", result: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, 0, 8}, input: 8, output: 0},
		{desc: "ex13", program: "3,3,1108,-1,8,3,4,3,99", result: []int{3, 3, 1108, 1, 8, 3, 4, 3, 99}, input: 8, output: 1},
		{desc: "ex14", program: "3,3,1108,-1,8,3,4,3,99", result: []int{3, 3, 1108, 0, 8, 3, 4, 3, 99}, input: 9, output: 0},
		{desc: "ex15", program: "3,3,1107,-1,8,3,4,3,99", result: []int{3, 3, 1107, 1, 8, 3, 4, 3, 99}, input: 7, output: 1},
		{desc: "ex16", program: "3,3,1107,-1,8,3,4,3,99", result: []int{3, 3, 1107, 0, 8, 3, 4, 3, 99}, input: 8, output: 0},
		{desc: "ex17", program: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", result: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 0, 0, 1, 9}, input: 0, output: 0},
		{desc: "ex18", program: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", result: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 1, 1, 1, 9}, input: 1, output: 1},
		{desc: "ex19", program: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", result: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -11, 1, 1, 9}, input: -11, output: 1},
		{desc: "ex20", program: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", result: []int{3, 3, 1105, 0, 9, 1101, 0, 0, 12, 4, 12, 99, 0}, input: 0, output: 0},
		{desc: "ex21", program: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", result: []int{3, 3, 1105, 1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, input: 1, output: 1},
		{desc: "ex22", program: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", result: []int{3, 3, 1105, -11, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, input: -11, output: 1},

		{desc: "ex23", program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", input: 8, output: 1000, result: []int{}},
		{desc: "ex23", program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", input: 7, output: 999, result: []int{}},
		{desc: "ex23", program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", input: 9, output: 1001, result: []int{}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result, output := run(tC.program, tC.input)
			if output != tC.output {
				t.Fatalf("expected output [%d] got [%d]", tC.output, output)
			}
			for i, val := range tC.result {
				if val != result[i] {
					t.Fatalf("expected [%v] got [%v]", tC.result, result)
				}
			}
		})
	}
}
