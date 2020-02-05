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

	fmt.Printf("%d", runAmplifiers(string(program)))
}

func amplify(program string, input int, phases []int) int {
	var output int
	_, output = run(program, []int{phases[0], 0})
	_, output = run(program, []int{phases[1], output})
	_, output = run(program, []int{phases[2], output})
	_, output = run(program, []int{phases[3], output})
	_, output = run(program, []int{phases[4], output})
	return output
}

func runAmplifiers(program string) int {
	phaseList := generatePhases([]int{0, 1, 2, 3, 4})
	maxOutput := 0
	for _, phases := range phaseList {
		output := amplify(program, 0, phases)
		if output > maxOutput {
			maxOutput = output
		}
	}
	return maxOutput
}

func run(program string, input []int) ([]int, int) {
	tokens := strings.Split(program, ",")
	codes := make([]int, len(tokens))
	for i, token := range tokens {
		code, err := strconv.Atoi(token)
		if err != nil {
			log.Fatalf("Could not convert to int %v - %s", token, err)
		}
		codes[i] = code
	}
	ic, pc, output := 0, 0, 0
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
			codes[reg1] = input[ic]
			ic++
			// fmt.Printf("%v - pc: %d, modes: %d, reg: %v, input: %d, reg1: %d\n", codes[pc:pc+registers+1], pc, modes, []int{reg1}, input, codes[reg1])
			pc += registers + 1
			continue
		}

		if op == 4 {
			registers := 1
			reg1 := read(codes, modes, pc, 1)
			// fmt.Printf("%v - pc: %d, modes: %d, reg: %v\n", codes[pc:pc+registers+1], pc, modes, []int{reg1})
			output = reg1
			// fmt.Printf("pc: %d - output: %d\n", pc, output)
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

func generatePhases(input []int) [][]int {
	result := [][]int{}
	input5 := input
	for i1, val1 := range input5 {
		input4 := append([]int(nil), input5[:i1]...)
		input4 = append(input4, input5[i1+1:]...)
		if len(input4) == 0 {
			result = append(result, []int{val1})
		}
		for i2, val2 := range input4 {
			input3 := append([]int(nil), input4[:i2]...)
			input3 = append(input3, input4[i2+1:]...)
			if len(input3) == 0 {
				result = append(result, []int{val1, val2})
			}
			for i3, val3 := range input3 {
				input2 := append([]int(nil), input3[:i3]...)
				input2 = append(input2, input3[i3+1:]...)
				if len(input2) == 0 {
					result = append(result, []int{val1, val2, val3})
				}
				for i4, val4 := range input2 {
					input1 := append([]int(nil), input2[:i4]...)
					input1 = append(input1, input2[i4+1:]...)
					for _, val5 := range input1 {
						result = append(result, []int{val1, val2, val3, val4, val5})
					}
				}
			}
		}
	}
	return result
}
