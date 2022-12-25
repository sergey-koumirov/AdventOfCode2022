package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	s := loadInput()
	recalcWalls(&s)
	s.ExitPoint = Point{Row: s.Rows - 1, Col: s.Cols - 2}
	s.Current = map[Point]int{{Row: 0, Col: 1}: 1}

	for i := 0; i < 1000; i++ {
		moveBlizzards(&s)
		recalcWalls(&s)
		movePoints(&s)
		_, exit := s.Current[s.ExitPoint]
		if exit {
			fmt.Println("FOUND EXIT-1 !!!", s.Minute)
			break
		}
	}

	s.Current = map[Point]int{{Row: s.ExitPoint.Row, Col: s.ExitPoint.Col}: 1}
	s.ExitPoint = Point{Row: 0, Col: 1}

	for i := 0; i < 1000; i++ {
		moveBlizzards(&s)
		recalcWalls(&s)
		movePoints(&s)
		_, exit := s.Current[s.ExitPoint]
		if exit {
			fmt.Println("RETURN !!!", s.Minute)
			break
		}
	}

	s.ExitPoint = Point{Row: s.Rows - 1, Col: s.Cols - 2}
	s.Current = map[Point]int{{Row: 0, Col: 1}: 1}

	for i := 0; i < 1000; i++ {
		moveBlizzards(&s)
		recalcWalls(&s)
		movePoints(&s)
		_, exit := s.Current[s.ExitPoint]
		if exit {
			fmt.Println("FOUND EXIT-2 !!!", s.Minute)
			break
		}
	}
}

func part2() {
	loadInput()
	fmt.Println("P2:")
}

func recalcWalls(s *State) {
	L := len(s.Blizzards)

	newWalls := map[int]map[int]int{}
	for row := 0; row < s.Rows; row++ {
		newWalls[row] = map[int]int{}
		newWalls[row][0] = 1
		newWalls[row][s.Cols-1] = 1
	}
	for col := 0; col < s.Cols; col++ {
		if col != 1 {
			newWalls[0][col] = 1
		}
		if col != s.Cols-2 {
			newWalls[s.Rows-1][col] = 1
		}
	}

	for i := 0; i < L; i++ {
		b := s.Blizzards[i]
		if b.Dir == '>' || b.Dir == '<' || b.Dir == 'v' || b.Dir == '^' {
			newWalls[s.Blizzards[i].Row][s.Blizzards[i].Col] = 1
		}

	}
	s.Walls = newWalls
}

func movePoints(s *State) {
	newCurrent := map[Point]int{}
	for p := range s.Current {
		// fmt.Println("Point", p)
		if empty(s, p.Row+1, p.Col) {
			// fmt.Println("P  +1   0")
			newCurrent[Point{Row: p.Row + 1, Col: p.Col}] = 1
		}
		if empty(s, p.Row-1, p.Col) {
			// fmt.Println("P  -1   0")
			newCurrent[Point{Row: p.Row - 1, Col: p.Col}] = 1
		}
		if empty(s, p.Row, p.Col+1) {
			// fmt.Println("P   0  +1")
			newCurrent[Point{Row: p.Row, Col: p.Col + 1}] = 1
		}
		if empty(s, p.Row, p.Col-1) {
			// fmt.Println("P   0  -1")
			newCurrent[Point{Row: p.Row, Col: p.Col - 1}] = 1
		}
		if empty(s, p.Row, p.Col) {
			// fmt.Println("P   0   0")
			newCurrent[Point{Row: p.Row, Col: p.Col}] = 1
		}
	}
	s.Current = newCurrent
}

func empty(s *State, row int, col int) bool {
	if row < 0 || row > s.Rows-1 || col < 0 || col > s.Cols-1 {
		return false
	}

	cols, exRow := s.Walls[row]
	if exRow {
		_, exCol := cols[col]
		return !exCol
	}

	return true
}

func moveBlizzards(s *State) {
	s.Minute++
	L := len(s.Blizzards)

	FRow := s.Rows - 2
	FCol := s.Cols - 2

	for i := 0; i < L; i++ {
		b := s.Blizzards[i]
		if b.Dir == '>' || b.Dir == '<' || b.Dir == 'v' || b.Dir == '^' {
			if b.Dir == '>' {
				s.Blizzards[i].Col = 1 + b.Col%FCol
			}
			if b.Dir == '<' {
				s.Blizzards[i].Col = 1 + (FCol+(b.Col-2))%FCol
			}
			if b.Dir == 'v' {
				s.Blizzards[i].Row = 1 + b.Row%FRow
			}
			if b.Dir == '^' {
				s.Blizzards[i].Row = 1 + (FRow+(b.Row-2))%FRow
			}
		}
	}
}

type Point struct {
	Row int
	Col int
}

type Blizzard struct {
	Row int
	Col int
	Dir rune
}

type State struct {
	Rows      int
	Cols      int
	Blizzards []Blizzard
	Walls     map[int]map[int]int
	Minute    int
	Current   map[Point]int
	ExitPoint Point
}

func printState(s *State) {
	fmt.Println("M:", s.Minute)
	for row := 0; row < s.Rows; row++ {
		for col := 0; col < s.Cols; col++ {
			if row == 0 && col == 1 {
				fmt.Print(".")
			} else if row == s.Rows-1 && col == s.Cols-2 {
				fmt.Print(".")
			} else if row == 0 || row == s.Rows-1 || col == 0 || col == s.Cols-1 {
				fmt.Print("#")
			} else {
				ex := 0
				sym := ""
				for _, b := range s.Blizzards {
					if b.Row == row && b.Col == col {
						ex++
						sym = string(b.Dir)
					}
				}
				if ex == 0 {
					fmt.Print(".")
				} else if ex == 1 {
					fmt.Print(sym)
				} else if ex < 10 {
					fmt.Print(strconv.Itoa(ex))
				} else {
					fmt.Print("@")
				}
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func loadInput() State {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	res := State{
		Blizzards: []Blizzard{},
	}

	rows := 0

	for scanner.Scan() {
		line := scanner.Text()
		if res.Cols == 0 {
			res.Cols = len(line)
		}
		if line != "" {
			for i, c := range line {
				if c == '>' || c == '<' || c == 'v' || c == '^' {
					res.Blizzards = append(res.Blizzards, Blizzard{Row: rows, Col: i, Dir: c})
				}
			}
			rows++
		}
	}
	res.Rows = rows

	return res
}
