package main

import "strconv"

import "fmt"

func main() {
	count := 0
	for index := 183564; index <= 657474; index++ {
		value := strconv.Itoa(index)
		if hasPair(value) && increases(value) {
			count++
		}
	}
	fmt.Printf("%d\n", count)
}

func hasDoubleDigits(bytes string) bool {
	len := len(bytes)
	for index := 0; index < len-1; index++ {
		if bytes[index] == bytes[index+1] {
			return true
		}
	}
	return false
}

func hasPair(bytes string) bool {
	m := make(map[byte]int)
	for index := 0; index < len(bytes); index++ {
		m[bytes[index]]++
	}
	for _, v := range m {
		if v == 2 {
			return true
		}
	}
	return false
}

func increases(bytes string) bool {
	for index := 0; index < len(bytes)-1; index++ {
		if bytes[index] > bytes[index+1] {
			return false
		}
	}
	return true
}
