package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"

	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

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

	fmt.Printf("Min near crossing = %d\n", getCrossDistanceNear(pathString1, pathString2))
	fmt.Printf("Min path crossing = %d\n", getCrossDistancePath(pathString1, pathString2))

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
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
	seen     map[string]coord
}

func (w *wirePath) contains(pos coord) coord {
	key := strconv.Itoa(pos.x) + ":" + strconv.Itoa(pos.y)
	c, found := w.seen[key]
	if found {
		return c
	}
	return coord{}
}

func (w *wirePath) add(step string) {
	direction := step[0]
	length, err := strconv.Atoi(step[1:])
	if err != nil {
		log.Fatalf("error parsing %s", step)
	}

	for i := 0; i < length; i++ {
		w.distance++

		switch direction {
		case 'R':
			w.pos = coord{x: w.pos.x + 1, y: w.pos.y, d: w.distance}
		case 'L':
			w.pos = coord{x: w.pos.x - 1, y: w.pos.y, d: w.distance}
		case 'U':
			w.pos = coord{x: w.pos.x, y: w.pos.y + 1, d: w.distance}
		case 'D':
			w.pos = coord{x: w.pos.x, y: w.pos.y - 1, d: w.distance}
		}

		w.coords = append(w.coords, w.pos)
		key := strconv.Itoa(w.pos.x) + ":" + strconv.Itoa(w.pos.y)
		w.seen[key] = w.pos
	}
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
	path := wirePath{coord{0, 0, 0}, 0, nil, map[string]coord{}}
	steps := strings.Split(pathString, ",")
	for _, step := range steps {
		path.add(step)
	}
	return path
}
