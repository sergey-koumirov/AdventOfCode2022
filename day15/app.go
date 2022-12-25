package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part1() {
	state := loadInput()

	targetY := int64(10)

	ranges := []Range{}
	for i, s := range state.Sensors {
		deltaX := state.Dist[i] - abs(targetY-s.Y)
		if deltaX >= 0 {
			ranges = append(ranges, Range{V1: s.X - deltaX, V2: s.X + deltaX})
		}
	}

	compacted := compact(ranges)
	sum := calcSum(compacted, state, targetY)

	fmt.Println("P1:", sum)
	printField(state)
}

func part2() {
	state := loadInput()

	for targetY := int64(0); targetY <= 4_000_000; targetY++ {
		ranges := []Range{}
		for i, s := range state.Sensors {
			deltaX := state.Dist[i] - abs(targetY-s.Y)
			if deltaX >= 0 {
				ranges = append(ranges, Range{V1: s.X - deltaX, V2: s.X + deltaX})
			}
		}
		compacted := compact(ranges)
		if len(compacted) > 1 {
			fmt.Println(targetY, compacted)
			x := int64(0)
			if compacted[0].V1 < compacted[1].V1 {
				x = compacted[0].V2 + 1
			} else {
				x = compacted[0].V1 - 1
			}
			fmt.Println("P2:", x*4_000_000+targetY)
		}
	}
}

func calcSum(compacted []Range, state State, targetY int64) int64 {
	sum := int64(0)
	for _, r := range compacted {
		sum += r.V2 - r.V1 + 1
	}
	for _, b := range state.Beacons {
		if b.Y == targetY {
			sum--
		}
	}
	for _, s := range state.Sensors {
		if s.Y == targetY {
			sum--
		}
	}
	return sum
}

func printField(state State) {

	for row := state.MinY; row <= state.MaxY; row++ {
		for col := state.MinX; col <= state.MaxX; col++ {
			empty := true
			for _, b := range state.Beacons {
				if b.X == col && b.Y == row {
					empty = false
					fmt.Print("B")
				}
			}
			for _, s := range state.Sensors {
				if s.X == col && s.Y == row {
					empty = false
					fmt.Print("S")
				}
			}
			if empty {
				for i, s := range state.Sensors {
					if empty && abs(s.X-col)+abs(s.Y-row) <= state.Dist[i] {
						empty = false
						if row >= 0 && row <= 20 && col >= 0 && col <= 20 {
							fmt.Print("*")
						} else {
							fmt.Print("#")
						}
					}
				}
			}

			if empty {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}

func compact(ranges []Range) []Range {
	L := len(ranges)
	res := []Range{}
	cnt := 0
	for i := 0; i < L; i++ {
		if !ranges[i].Used {
			ranges[i].Used = true
			test := Range{V1: ranges[i].V1, V2: ranges[i].V2}
			for j := i + 1; j < L; j++ {
				if ranges[j].V1 <= test.V2 && test.V1 <= ranges[j].V2 || ranges[j].V2+1 == test.V1 || test.V2+1 == ranges[j].V1 {
					test.V1 = min(ranges[j].V1, test.V1)
					test.V2 = max(ranges[j].V2, test.V2)
					ranges[j].Used = true
					cnt++
				}
			}
			res = append(res, test)
		}
	}

	if cnt > 0 {
		res = compact(res)
	}

	return res
}

type Point struct {
	X int64
	Y int64
}

type Range struct {
	V1   int64
	V2   int64
	Used bool
}

type State struct {
	Sensors []Point
	Beacons []Point
	Dist    []int64
	MinX    int64
	MaxX    int64
	MinY    int64
	MaxY    int64
}

func loadInput() State {
	res := State{
		Sensors: []Point{},
		Beacons: []Point{},
		Dist:    []int64{},
	}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	minX := int64(math.MaxInt64)
	maxX := int64(math.MinInt64)
	minY := int64(math.MaxInt64)
	maxY := int64(math.MinInt64)

	for scanner.Scan() {
		line := scanner.Text()

		// Sensor at x=44398, y=534916: closest beacon is at x=278652, y=-182407
		parts := strings.Split(line, ":")

		parts0 := strings.Split(parts[0], ",")
		parts01 := strings.Split(parts0[0], "=")
		parts02 := strings.Split(parts0[1], "=")

		parts1 := strings.Split(parts[1], ",")
		parts11 := strings.Split(parts1[0], "=")
		parts12 := strings.Split(parts1[1], "=")

		sx, _ := strconv.ParseInt(parts01[1], 10, 64)
		sy, _ := strconv.ParseInt(parts02[1], 10, 64)
		bx, _ := strconv.ParseInt(parts11[1], 10, 64)
		by, _ := strconv.ParseInt(parts12[1], 10, 64)

		res.Sensors = append(res.Sensors, Point{X: sx, Y: sy})
		dist := abs(sx-bx) + abs(sy-by)
		res.Dist = append(res.Dist, dist)

		ex := false
		for _, b := range res.Beacons {
			if b.X == bx && b.Y == by {
				ex = true
			}
		}
		if !ex {
			res.Beacons = append(res.Beacons, Point{X: bx, Y: by})
		}

		applyMinMax(&minX, &maxX, &minY, &maxY, sx-dist, sy)
		applyMinMax(&minX, &maxX, &minY, &maxY, sx+dist, sy)
		applyMinMax(&minX, &maxX, &minY, &maxY, sx, sy-dist)
		applyMinMax(&minX, &maxX, &minY, &maxY, sx, sy+dist)
		applyMinMax(&minX, &maxX, &minY, &maxY, sx, sy)
		applyMinMax(&minX, &maxX, &minY, &maxY, bx, by)
	}

	res.MaxX = maxX
	res.MinX = minX
	res.MaxY = maxY
	res.MinY = minY

	return res
}

func abs(v int64) int64 {
	if v >= 0 {
		return v
	} else {
		return v * -1
	}
}

func min(v1, v2 int64) int64 {
	if v1 < v2 {
		return v1
	}
	return v2
}

func max(v1, v2 int64) int64 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func applyMinMax(minX *int64, maxX *int64, minY *int64, maxY *int64, x, y int64) {
	if x > *maxX {
		*maxX = x
	}
	if x < *minX {
		*minX = x
	}

	if y > *maxY {
		*maxY = y
	}
	if y < *minY {
		*minY = y
	}
}
