package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	program, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	run(string(program), 5)
}

func run(program string, input int) ([]int, int) {
	tokens := strings.Split(program, ",")
	codes := make([]int, len(tokens))
	for i, token := range tokens {
		code, err := strconv.Atoi(token)
		if err != nil {
			log.Fatalf("Could not convert to int %v - %s", token, err)
		}
		codes[i] = code
	}
	output := 0
	pc := 0
	for {
		op, modes := decomposeCommand(codes[pc])
		if op == 99 {
			break
		}

		if op == 1 {
			registers := 3
			reg1 := read(codes, modes, pc, 1)
			reg2 := read(codes, modes, pc, 2)
			reg3 := read(codes, 100, pc, 3)
			codes[reg3] = reg1 + reg2
			// fmt.Printf("%v - pc: %d, modes: %d, reg: %v, reg3: %d\n", codes[pc:pc+registers+1], pc, modes, []int{reg1, reg2, reg3}, codes[reg3])
			pc += registers + 1
			continue
		}

		if op == 2 {
			registers := 3
			reg1 := read(codes, modes, pc, 1)
			reg2 := read(codes, modes, pc, 2)
			reg3 := read(codes, 100, pc, 3)
			codes[reg3] = reg1 * reg2
			// fmt.Printf("%v - pc: %d, modes: %d, reg: %v, reg3: %d\n", codes[pc:pc+registers+1], pc, modes, []int{reg1, reg2, reg3}, codes[reg3])
			pc += registers + 1
			continue
		}

		if op == 3 {
			registers := 1
			reg1 := read(codes, 1, pc, 1)
			codes[reg1] = input
			// fmt.Printf("%v - pc: %d, modes: %d, reg: %v, input: %d, reg1: %d\n", codes[pc:pc+registers+1], pc, modes, []int{reg1}, input, codes[reg1])
			pc += registers + 1
			continue
		}

		if op == 4 {
			registers := 1
			reg1 := read(codes, modes, pc, 1)
			// fmt.Printf("%v - pc: %d, modes: %d, reg: %v\n", codes[pc:pc+registers+1], pc, modes, []int{reg1})
			output = reg1
			fmt.Printf("pc: %d - output: %d\n", pc, output)
			pc += registers + 1
			continue
		}

		if op == 5 {
			registers := 2
			reg1 := read(codes, modes, pc, 1)
			reg2 := read(codes, modes, pc, 2)
			// fmt.Printf("%v - pc: %d, modes: %d, reg: %v, reg3: %d\n", codes[pc:pc+registers+1], pc, modes, []int{reg1, reg2, reg3}, codes[reg3])
			if reg1 != 0 {
				pc = reg2
			} else {
				pc += registers + 1
			}
			continue
		}

		if op == 6 {
			registers := 2
			reg1 := read(codes, modes, pc, 1)
			reg2 := read(codes, modes, pc, 2)
			// fmt.Printf("%v - pc: %d, modes: %d, reg: %v, reg3: %d\n", codes[pc:pc+registers+1], pc, modes, []int{reg1, reg2, reg3}, codes[reg3])
			if reg1 == 0 {
				pc = reg2
			} else {
				pc += registers + 1
			}
			continue
		}

		if op == 7 {
			registers := 3
			reg1 := read(codes, modes, pc, 1)
			reg2 := read(codes, modes, pc, 2)
			reg3 := read(codes, 100, pc, 3)
			if reg1 < reg2 {
				codes[reg3] = 1
			} else {
				codes[reg3] = 0
			}
			// fmt.Printf("%v - pc: %d, modes: %d, reg: %v, reg3: %d\n", codes[pc:pc+registers+1], pc, modes, []int{reg1, reg2, reg3}, codes[reg3])
			pc += registers + 1
			continue
		}

		if op == 8 {
			registers := 3
			reg1 := read(codes, modes, pc, 1)
			reg2 := read(codes, modes, pc, 2)
			reg3 := read(codes, 100, pc, 3)
			if reg1 == reg2 {
				codes[reg3] = 1
			} else {
				codes[reg3] = 0
			}
			// fmt.Printf("%v - pc: %d, modes: %d, reg: %v, reg3: %d\n", codes[pc:pc+registers+1], pc, modes, []int{reg1, reg2, reg3}, codes[reg3])
			pc += registers + 1
			continue
		}

		log.Fatalf("invalid opcode %d", op)
	}
	return codes, output
}

func decomposeCommand(command int) (int, int) {
	opCode := command % 100
	modes := command / 100
	return opCode, modes
}

func read(codes []int, modes, pc, offset int) int {
	div := int(math.Pow10(offset - 1))
	mode := (modes / div) % 10
	value := codes[pc+offset]
	if mode == 0 {
		return codes[value]
	}
	return value
}
