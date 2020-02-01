package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	program, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			result := run(string(program), noun, verb)
			fmt.Printf("noun: %d, verb: %d -> %d\n", noun, verb, result[0])
			if result[0] == 19690720 {
				return
			}
		}
	}
}

func run(program string, noun int, verb int) []int {
	tokens := strings.Split(program, ",")
	codes := make([]int, len(tokens))
	for i, token := range tokens {
		code, err := strconv.Atoi(token)
		if err != nil {
			log.Fatalf("Could not convert to int %v - %s", token, err)
		}
		codes[i] = code
	}
	codes[1] = noun
	codes[2] = verb
	pc := 0
	for {
		op, modes := decomposeCommand(codes[pc])
		if op == 99 {
			break
		}

		_ = modes

		// reg1, err := read(tokens, modes, pc, 1)

		reg1 := readRelative(codes, pc+1)
		reg2 := readRelative(codes, pc+2)
		reg3 := readPosition(codes, pc+3)

		if op == 1 {
			codes[reg3] = reg1 + reg2
			pc += 4
		}

		if op == 2 {
			codes[reg3] = reg1 * reg2
			pc += 4
		}
		// fmt.Printf("pc:%d op:%s tokens:%v\n", pc, op, tokens)
	}
	return codes
}

func decomposeCommand(command int) (int, int) {
	opCode := command % 100
	modes := command / 100
	return opCode, modes
}

func read(codes []int, modes, pc, offset int) int {
	return 0
}

func readPosition(codes []int, position int) int {
	return codes[position]
}

func readRelative(codes []int, position int) int {
	return codes[codes[position]]
}
