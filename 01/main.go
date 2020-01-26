package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sum := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text(), err)
		}
		sum += calculateFuelRecursive(n)
	}
	if err = s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total = %d", sum)
}

func calculateFuel(mass int) int {
	return (mass / 3) - 2
}

func calculateFuelRecursive(mass int) int {
	if mass <= 0 {
		return 0
	}
	fuel := (mass / 3) - 2
	result := fuel + calculateFuelRecursive(fuel)
	if result < 0 {
		result = 0
	}
	return result
}
