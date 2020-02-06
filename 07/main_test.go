package main

import (
	"sync"
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

func TestGeneratePhases(t *testing.T) {
	// t.Skip()
	testCases := []struct {
		desc   string
		input  []int
		output [][]int
	}{
		{desc: "ex01", input: []int{}, output: [][]int{}},
		{desc: "ex02", input: []int{0}, output: [][]int{{0}}},
		{desc: "ex03", input: []int{0, 1}, output: [][]int{{0, 1}, {1, 0}}},
		{desc: "ex04", input: []int{0, 1, 2}, output: [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := generatePhases(tC.input)
			if len(output) != len(tC.output) {
				t.Fatalf("expected output [%v] got [%v]", tC.output, output)
			}
			for i := range tC.output {
				if len(output[i]) != len(tC.output[i]) {
					t.Fatalf("expected output [%v] got [%v]", tC.output, output)
				}
				for j := range output[i] {
					if output[i][j] != tC.output[i][j] {
						t.Fatalf("expected output [%v] got [%v]", tC.output, output)
					}
				}
			}
		})
	}
}

func TestRunAmplifiers(t *testing.T) {
	// t.Skip()
	testCases := []struct {
		desc    string
		program string
		input   int
		phases  []int
		output  int
	}{
		{desc: "ex01", program: "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0", input: 0, phases: []int{4, 3, 2, 1, 0}, output: 43210},
		{desc: "ex02", program: "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0", input: 0, phases: []int{0, 1, 2, 3, 4}, output: 54321},
		{desc: "ex03", program: "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0", input: 0, phases: []int{1, 0, 4, 3, 2}, output: 65210},
		{desc: "ex04", program: "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5", input: 0, phases: []int{9, 8, 7, 6, 5}, output: 139629729},
		{
			desc:    "ex05",
			program: "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10",
			input:   0,
			phases:  []int{9, 7, 8, 5, 6},
			output:  18216,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := amplify(tC.program, tC.input, tC.phases)
			if output != tC.output {
				t.Fatalf("expected output [%d] got [%d]", tC.output, output)
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
		output  int
	}{
		{desc: "ex09", program: "3,9,8,9,10,9,4,9,99,-1,8", input: 8, output: 1},
		{desc: "ex10", program: "3,9,8,9,10,9,4,9,99,-1,8", input: 9, output: 0},
		{desc: "ex11", program: "3,9,7,9,10,9,4,9,99,-1,8", input: 7, output: 1},
		{desc: "ex12", program: "3,9,7,9,10,9,4,9,99,-1,8", input: 8, output: 0},
		{desc: "ex13", program: "3,3,1108,-1,8,3,4,3,99", input: 8, output: 1},
		{desc: "ex14", program: "3,3,1108,-1,8,3,4,3,99", input: 9, output: 0},
		{desc: "ex15", program: "3,3,1107,-1,8,3,4,3,99", input: 7, output: 1},
		{desc: "ex16", program: "3,3,1107,-1,8,3,4,3,99", input: 8, output: 0},
		{desc: "ex17", program: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", input: 0, output: 0},
		{desc: "ex18", program: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", input: 1, output: 1},
		{desc: "ex19", program: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", input: -11, output: 1},
		{desc: "ex20", program: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", input: 0, output: 0},
		{desc: "ex21", program: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", input: 1, output: 1},
		{desc: "ex22", program: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", input: -11, output: 1},

		{desc: "ex23", program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", input: 8, output: 1000},
		{desc: "ex24", program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", input: 7, output: 999},
		{desc: "ex25", program: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", input: 9, output: 1001},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			c1in := make(chan int)
			c1out := make(chan int)
			var wg sync.WaitGroup
			wg.Add(1)
			go run(tC.desc, tC.program, c1in, c1out, &wg)
			c1in <- tC.input
			output := <-c1out
			wg.Wait()
			if output != tC.output {
				t.Fatalf("expected output [%d] got [%d]", tC.output, output)
			}
		})
	}
}
