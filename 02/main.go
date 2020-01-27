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
			fmt.Printf("noun: %d, verb: %d -> %s\n", noun, verb, result[0])
			if result[0] == "19690720" {
				return
			}
		}
	}
}

func run(program string, noun int, verb int) []string {
	tokens := strings.Split(program, ",")
	tokens[1] = strconv.Itoa(noun)
	tokens[2] = strconv.Itoa(verb)
	pc := 0
	for true {
		op := tokens[pc]
		if op == "99" {
			break
		}

		reg1, err := readRelative(tokens, pc+1)
		if err != nil {
			log.Fatal(err)
		}
		reg2, err := readRelative(tokens, pc+2)
		if err != nil {
			log.Fatal(err)
		}
		reg3, err := readPosition(tokens, pc+3)
		if err != nil {
			log.Fatal(err)
		}

		if op == "1" {
			tokens[reg3] = strconv.Itoa(reg1 + reg2)
			pc += 4
		}

		if op == "2" {
			tokens[reg3] = strconv.Itoa(reg1 * reg2)
			pc += 4
		}
		// fmt.Printf("pc:%d op:%s tokens:%v\n", pc, op, tokens)
	}
	return tokens
}

func readPosition(tokens []string, position int) (int, error) {
	result, err := strconv.Atoi(tokens[position])
	if err != nil {
		return 0, err
	}
	return result, nil
}

func readRelative(tokens []string, position int) (int, error) {
	target, err := strconv.Atoi(tokens[position])
	if err != nil {
		return 0, err
	}
	result, err := strconv.Atoi(tokens[target])
	if err != nil {
		return 0, err
	}
	// fmt.Printf("%v [%d] = %d\n", tokens, position, result)
	return result, nil
}
