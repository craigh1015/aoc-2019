package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func main() {
	program, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	minZeroCount := 10000
	minZeroLayer := 0
	minZeroLayerOnesCount := 0
	minZeroLayerTwosCount := 0

	image := [6][25]byte{}
	for row := 0; row < len(image); row++ {
		for p := 0; p < len(image[row]); p++ {
			image[row][p] = '2'
		}
	}

	for layer := 0; layer < 100; layer++ {
		zeroCount, onesCount, twosCount := 0, 0, 0
		// fmt.Printf("\nlayer: %d\n", layer)
		for row := 0; row < 6; row++ {
			for p := 0; p < 25; p++ {
				pos := (layer * 25 * 6) + (row * 25) + p
				// val, _ := strconv.Atoi(string(program[pos]))
				// fmt.Print(val)
				switch program[pos] {
				case '0':
					zeroCount++
				case '1':
					onesCount++
				case '2':
					twosCount++
				}
				if image[row][p] == '2' {
					image[row][p] = program[pos]
				}
			}
			// fmt.Print("\n")
		}
		if zeroCount < minZeroCount {
			minZeroCount = zeroCount
			minZeroLayer = layer
			minZeroLayerOnesCount = onesCount
			minZeroLayerTwosCount = twosCount
		}
	}
	fmt.Printf("minZeroLayer: %d, result: %d\n", minZeroLayer, (minZeroLayerOnesCount * minZeroLayerTwosCount))

	for row := 0; row < len(image); row++ {
		for p := 0; p < len(image[row]); p++ {
			val, _ := strconv.Atoi(string(image[row][p]))
			if val == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(val)
			}
		}
		fmt.Print("\n")
	}

}
