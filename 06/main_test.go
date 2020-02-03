package main

import (
	"sort"
	"testing"

	"github.com/craigh1015/aoc-2019/06/tree"
)

func TestTreeMakeNode(t *testing.T) {
	root := tree.MakeNode("A")
	if !root.IsLeaf() {
		t.Fatal("single node should be leaf")
	}
	child := tree.MakeNode("B")
	root.Append(child)
	if root.IsLeaf() {
		t.Fatal("parent node should not be leaf")
	}
}

func TestTreeGetPathRoot(t *testing.T) {
	root := tree.MakeNode("A")
	if len(root.GetPaths("")) != 0 {
		t.Fatal("Single node has no paths")
	}
}

func TestTreeGetPathRootAndChild(t *testing.T) {
	root := tree.MakeNode("A")
	child := tree.MakeNode("B")
	root.Append(child)
	expected := []string{"A->B"}
	actual := root.GetPaths("")
	for i, path := range actual {
		if path != expected[i] {
			t.Fatalf("expected: %v got %v\n", expected, actual)
		}
	}
}

func TestTreeGetPathRootWithChildren(t *testing.T) {
	a := tree.MakeNode("A")
	aa := tree.MakeNode("AA")
	ab := tree.MakeNode("AB")
	aaa := tree.MakeNode("AAA")
	a.Append(aa)
	a.Append(ab)
	aa.Append(aaa)
	expected := []string{"A->AA", "A->AA->AAA", "A->AB"}
	actual := a.GetPaths("")
	for i, path := range actual {
		if path != expected[i] {
			t.Fatalf("expected: %v got %v\n", expected, actual)
		}
	}
}

func TestTreeMakeTree(t *testing.T) {
	testCases := []struct {
		desc   string
		orbits []string
		paths  []string
	}{
		{"ex01", []string{"COM)A"}, []string{"COM->A"}},
		{"ex02", []string{"COM)A", "A)B"}, []string{"COM->A", "COM->A->B"}},
		{"ex03", []string{"COM)A", "A)B", "COM)C", "B)D"}, []string{"COM->A", "COM->A->B", "COM->A->B->D", "COM->C"}},
		{"ex04", []string{"A)B", "COM)C", "B)D", "COM)A"}, []string{"COM->A", "COM->A->B", "COM->A->B->D", "COM->C"}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tree := makeTree(tC.orbits, "COM")
			paths := tree.GetPaths("")
			sort.Strings(paths)
			if len(paths) != len(tC.paths) {
				t.Fatalf("expected: %v got %v\n", tC.paths, paths)
			}
			for i, path := range tC.paths {
				if path != paths[i] {
					t.Fatalf("expected: %v got %v\n", tC.paths, paths)
				}
			}
		})
	}
}

func TestRun(t *testing.T) {
	testCases := []struct {
		desc   string
		orbits []string
		count  int
	}{
		{"ex01", []string{"COM)A"}, 1},
		{"ex02", []string{"COM)A", "A)B"}, 3},
		{"ex03", []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}, 42},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			count := run(tC.orbits)
			if count != tC.count {
				t.Fatalf("orbits: %v expected: %d got: %d", tC.orbits, tC.count, count)
			}
		})
	}
}
