package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	start := time.Now()

	part1()
	// part2()
	elapsed := time.Since(start)
	fmt.Println("time(ms):", elapsed.Milliseconds())
}

func part1() {
	elves := loadInput()
	s := State{Elves: elves, Dir: 0}

	// fmt.Println("Elves:", len(s.Elves))
	for i := 1; i <= 1221; i++ {
		cnt := move(&s)
		fmt.Println("Cnt:", cnt, "    End of Round:", i, "     Dir: ", s.Dir, "  Elves:", len(s.Elves))
		// fmt.Println()
	}

	fmt.Println("P1:", emptyGround(&s))
}

func part2() {
	state := loadInput()
	fmt.Println(state)
}

var NextDir = map[int]int{3: 1, 1: 2, 2: 0, 0: 3}

var CheckDeltas = map[int][]Point{
	0: {
		{Row: -1, Col: 1},
		{Row: 0, Col: 1},
		{Row: 1, Col: 1},
	},
	1: {
		{Row: 1, Col: -1},
		{Row: 1, Col: 0},
		{Row: 1, Col: 1},
	},
	2: {
		{Row: 1, Col: -1},
		{Row: 0, Col: -1},
		{Row: -1, Col: -1},
	},
	3: {
		{Row: -1, Col: 1},
		{Row: -1, Col: 0},
		{Row: -1, Col: -1},
	},
}

var DirDeltas = map[int]Point{
	0: {Row: 0, Col: 1},
	1: {Row: 1, Col: 0},
	2: {Row: 0, Col: -1},
	3: {Row: -1, Col: 0},
}

func move(s *State) int {
	s.Dir = NextDir[s.Dir]

	propositions := CoordsHash{}
	dirs := CoordsHash{}

	for point := range s.Elves {
		empty := isEmptyAround(s, point)
		if !empty {
			dir := s.Dir
			applied := applyProps(s, point, dir, propositions, dirs)
			// fmt.Println(point, applied, dir)

			if !applied {
				dir := NextDir[dir]
				applied = applyProps(s, point, dir, propositions, dirs)
				// fmt.Println(point, applied, dir)
				if !applied {
					dir := NextDir[dir]
					applied = applyProps(s, point, dir, propositions, dirs)
					// fmt.Println(point, applied, dir)
					if !applied {
						dir := NextDir[dir]
						applied = applyProps(s, point, dir, propositions, dirs)
						// fmt.Println(point, applied, dir)
					}
				}
			}
		}
	}

	// printElves(s, &dirs)

	newElves := CoordsHash{}

	// fmt.Println(propositions)

	res := 0
	for point := range s.Elves {
		dir, dirEx := dirs[point]
		if dirEx {
			dd := DirDeltas[dir]

			temp := point
			temp.Row = temp.Row + dd.Row
			temp.Col = temp.Col + dd.Col

			cnt, ex := propositions[temp]

			// fmt.Println(point, temp, ex, cnt)

			if ex && cnt == 1 {
				newElves[temp] = 1
				res++
			} else {
				newElves[point] = 1
			}
		} else {
			newElves[point] = 1
		}
	}

	// fmt.Println(len(s.Elves), "-->", len(newElves))

	s.Elves = newElves

	// printElves(s, nil)

	return res
}

func applyProps(s *State, point Point, dir int, propositions CoordsHash, dirs CoordsHash) bool {
	ed := emptyDirward(s, point, dir)
	if ed {
		dd := DirDeltas[dir]
		temp := point
		temp.Row = temp.Row + dd.Row
		temp.Col = temp.Col + dd.Col
		propositions[temp] = propositions[temp] + 1
		dirs[point] = dir
	}
	return ed
}

func emptyGround(s *State) int {
	maxRow := math.MinInt
	minRow := math.MaxInt
	maxCol := math.MinInt
	minCol := math.MaxInt

	for point := range s.Elves {
		if maxRow < point.Row {
			maxRow = point.Row
		}
		if maxCol < point.Col {
			maxCol = point.Col
		}
		if minRow > point.Row {
			minRow = point.Row
		}
		if minCol > point.Col {
			minCol = point.Col
		}
	}

	return (maxRow-minRow+1)*(maxCol-minCol+1) - len(s.Elves)
}

func printElves(s *State, dirs *CoordsHash) {
	maxRow := math.MinInt
	minRow := math.MaxInt
	maxCol := math.MinInt
	minCol := math.MaxInt

	for point := range s.Elves {
		if maxRow < point.Row {
			maxRow = point.Row
		}
		if maxCol < point.Col {
			maxCol = point.Col
		}
		if minRow > point.Row {
			minRow = point.Row
		}
		if minCol > point.Col {
			minCol = point.Col
		}
	}

	// maxRow = 9
	// minRow = -2
	// maxCol = 10
	// minCol = -3

	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			_, ex := s.Elves[Point{Row: row, Col: col}]
			if ex {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		if dirs != nil {
			fmt.Print("     ")
			for col := minCol; col <= maxCol; col++ {
				d, ex := (*dirs)[Point{Row: row, Col: col}]
				if ex && d == 0 {
					fmt.Print("→")
				} else if ex && d == 1 {
					fmt.Print("↓")
				} else if ex && d == 2 {
					fmt.Print("←")
				} else if ex && d == 3 {
					fmt.Print("↑")
				} else {
					fmt.Print(".")
				}
			}
		}

		fmt.Println()
	}
	fmt.Println()
}

func emptyDirward(s *State, p Point, dir int) bool {
	deltas := CheckDeltas[dir]
	for _, d := range deltas {
		temp := p
		temp.Row = temp.Row + d.Row
		temp.Col = temp.Col + d.Col
		_, ex := s.Elves[temp]
		if ex {
			return false
		}
	}

	return true
}

func isEmptyAround(s *State, p Point) bool {
	for dRow := -1; dRow <= 1; dRow++ {
		for dCol := -1; dCol <= 1; dCol++ {
			if dRow != 0 || dCol != 0 {
				temp := p
				temp.Row = temp.Row + dRow
				temp.Col = temp.Col + dCol
				_, ex := s.Elves[temp]
				if ex {
					return false
				}
			}
		}
	}
	return true
}

type Point struct {
	Row int
	Col int
}

type CoordsHash map[Point]int

type State struct {
	Elves CoordsHash
	Dir   int
}

func loadInput() CoordsHash {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	res := CoordsHash{}
	row := 0
	for scanner.Scan() {
		line := scanner.Text()

		for col, c := range line {
			if c == '#' {
				temp := Point{Row: row, Col: col}
				res[temp] = 1
			}
		}
		row++
	}

	return res
}
