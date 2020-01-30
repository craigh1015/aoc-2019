package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan()
	pathString1 := s.Text()
	s.Scan()
	pathString2 := s.Text()

	fmt.Printf("Min near crossing = %d", getCrossDistanceNear(pathString1, pathString2))
	fmt.Printf("Min path crossing = %d", getCrossDistancePath(pathString1, pathString2))
}

func run(program string, noun int, verb int) []string {
	return nil
}

type coord struct {
	x, y, d int
}

func (lhs *coord) equal(rhs coord) bool {
	return lhs.x == rhs.x && lhs.y == rhs.y
}

type wirePath struct {
	pos      coord
	distance int
	coords   []coord
}

func (w *wirePath) contains(pos coord) coord {
	for _, coord := range w.coords {
		if coord.equal(pos) {
			return coord
		}
	}
	return coord{}
}

func getCrossDistancePath(pathString1, pathString2 string) int {
	return getMinDistancePath(getCrossings(pathString1, pathString2))
}

func getMinDistancePath(coords []coord) int {
	if len(coords) == 0 {
		return 0
	}
	minPath := math.MaxInt32
	for _, coord := range coords {
		if coord.d < minPath {
			minPath = coord.d
		}
	}
	return minPath
}

func getCrossDistanceNear(pathString1, pathString2 string) int {
	return getMinDistanceNear(getCrossings(pathString1, pathString2))
}

func getMinDistanceNear(coords []coord) int {
	if len(coords) == 0 {
		return 0
	}
	minDistance := math.MaxFloat64
	for _, coord := range coords {
		distance := math.Abs(float64(coord.x)) + math.Abs(float64(coord.y))
		if distance < minDistance {
			minDistance = distance
		}
	}
	return int(minDistance)
}

func getCrossings(pathString1, pathString2 string) []coord {
	path1 := makePath(pathString1)
	path2 := makePath(pathString2)

	crossings := []coord{}

	for _, pos := range path1.coords {
		otherPos := path2.contains(pos)
		if (otherPos != coord{}) {
			crossings = append(crossings, coord{x: pos.x, y: pos.y, d: pos.d + otherPos.d})
		}
	}

	return crossings
}

func makePath(pathString string) wirePath {
	path := wirePath{coord{0, 0, 0}, 0, nil}
	steps := strings.Split(pathString, ",")
	for _, step := range steps {
		direction := step[0]
		length, err := strconv.Atoi(step[1:])
		if err != nil {
			log.Fatalf("error parsing %s", step)
		}
		switch direction {
		case 'R':
			for i := 0; i < length; i++ {
				path.distance += 1
				path.pos = coord{x: path.pos.x + 1, y: path.pos.y, d: path.distance}
				path.coords = append(path.coords, path.pos)
			}
		case 'L':
			for i := 0; i < length; i++ {
				path.distance += 1
				path.pos = coord{x: path.pos.x - 1, y: path.pos.y, d: path.distance}
				path.coords = append(path.coords, path.pos)
			}
		case 'U':
			for i := 0; i < length; i++ {
				path.distance += 1
				path.pos = coord{x: path.pos.x, y: path.pos.y + 1, d: path.distance}
				path.coords = append(path.coords, path.pos)
			}
		case 'D':
			for i := 0; i < length; i++ {
				path.distance += 1
				path.pos = coord{x: path.pos.x, y: path.pos.y - 1, d: path.distance}
				path.coords = append(path.coords, path.pos)
			}
		}
	}
	return path
}
