package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"

	"github.com/craigh1015/aoc-2019/06/tree"
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

	orbits := []string{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		orbits = append(orbits, s.Text())
	}
	if err = s.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Total orbits: %d\n", run(orbits))

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

func makeTree(orbits []string, rootName string) *tree.Node {
	seen := make(map[string]*tree.Node)
	root := tree.MakeNode(rootName)
	seen[rootName] = root

	todo := orbits

	for {
		log.Printf("items to process - %d\n", len(todo))

		if len(todo) == 0 {
			break
		}

		unprocessed := []string{}

		for _, orbit := range todo {
			nodes := strings.Split(orbit, ")")
			if len(nodes) != 2 {
				log.Fatalf("unable to parse %s", nodes)
			}

			parentName := nodes[0]
			childName := nodes[1]

			_, found := seen[childName]
			if found {
				log.Fatalf("node %s from orbit %s is duplicate", childName, orbit)
			}

			node, found := seen[parentName]
			if found {
				child := tree.MakeNode(childName)
				seen[childName] = child
				node.Append(child)
				continue
			}

			unprocessed = append(unprocessed, orbit)
		}

		todo = unprocessed
	}

	return root
}

func run(orbits []string) int {
	tree := makeTree(orbits, "COM")
	paths := tree.GetPaths("")
	count := 0
	for _, path := range paths {
		count += strings.Count(path, "->")
		if strings.Contains(path, "YOU") || strings.Contains(path, "SAN") {
			log.Println(path)
		}
	}
	return count
}
