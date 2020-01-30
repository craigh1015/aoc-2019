package main

import (
	"testing"
)

func TestMakePath(t *testing.T) {
	testCases := []struct {
		desc   string
		path   string
		coords []coord
	}{
		{"ex1", "R1", []coord{{1, 0, 1}}},
		{"ex2", "R3", []coord{{1, 0, 1}, {2, 0, 2}, {3, 0, 3}}},
		{"ex3", "R1,U1,L1,D1", []coord{{1, 0, 1}, {1, 1, 2}, {0, 1, 3}, {0, 0, 4}}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := makePath(tC.path)
			if len(result.coords) != len(tC.coords) {
				t.Fatalf("[%v] is not [%v]", tC.coords, result.coords)
			}
			for i, val := range result.coords {
				if val.x != tC.coords[i].x || val.y != tC.coords[i].y {
					t.Fatalf("[%v] is not [%v]", tC.coords, result.coords)
				}
			}
		})
	}
}

func TestGetCrossings(t *testing.T) {
	testCases := []struct {
		desc      string
		path1     string
		path2     string
		crossings []coord
	}{
		{"ex1", "R3", "U3", []coord{}},
		{"ex2", "U2,R3", "R2,U3", []coord{{2, 2, 4}}},
		{"ex3", "U3,R3", "R1,U1,R1,U2", []coord{{2, 3, 5}}},
		{"ex3", "U3,R3,U3", "R1,U1,R1,U2,R2", []coord{{2, 3, 5}, {3, 3, 6}}},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := getCrossings(tC.path1, tC.path2)
			if len(result) != len(tC.crossings) {
				t.Fatalf("expected [%d] got [%d]", tC.crossings, result)
			}
			for i, pos := range result {
				if !pos.equal(tC.crossings[i]) {
					t.Fatalf("expected [%d] got [%d]", tC.crossings, result)
				}

			}
		})
	}
}

func TestGetMinDistance(t *testing.T) {
	testCases := []struct {
		desc      string
		crossings []coord
		distance  int
	}{
		{"ex1", []coord{}, 0},
		{"ex2", []coord{{2, 2, 4}}, 4},
		{"ex2", []coord{{2, 2, 4}, {2, 1, 3}}, 3},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := getMinDistanceNear(tC.crossings)
			if result != tC.distance {
				t.Fatalf("expected [%d] got [%d]", tC.distance, result)
			}
		})
	}
}

func TestGetCrossDistanceNear(t *testing.T) {
	testCases := []struct {
		desc     string
		path1    string
		path2    string
		distance int
	}{
		{"ex2", "R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"ex1", "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := getCrossDistanceNear(tC.path1, tC.path2)
			if result != tC.distance {
				t.Fatalf("expected [%d] got [%d]", tC.distance, result)
			}
		})
	}
}

func TestGetCrossDistancePath(t *testing.T) {
	testCases := []struct {
		desc     string
		path1    string
		path2    string
		distance int
	}{
		{"ex2", "R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 610},
		{"ex1", "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 410},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := getCrossDistancePath(tC.path1, tC.path2)
			if result != tC.distance {
				t.Fatalf("expected [%d] got [%d]", tC.distance, result)
			}
		})
	}
}
