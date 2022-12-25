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

	// part1()
	part2()

	elapsed := time.Since(start)
	fmt.Println("time(ms):", elapsed.Milliseconds())
}

func part1() {
	state := loadInput()

	fmt.Printf("P1: Row=%d Col=%d Face=%d\n", state.Row, state.Col, state.Dir)

	// printState(state)
	move(&state)

	fmt.Printf("P1: Row=%d Col=%d Face=%d Pass=%d\n", state.Row, state.Col, state.Dir, 1000*(state.Row+1)+4*(state.Col+1)+state.Dir)
}

func part2() {
	state := loadInput()
	move2(&state)
	fmt.Printf("P1: Row=%d Col=%d Face=%d Pass=%d\n", state.Row, state.Col, state.Dir, 1000*(state.Row+1)+4*(state.Col+1)+state.Dir)
}

var dirToDeltaRowCol = map[int][]int{
	0: {0, 1},
	1: {1, 0},
	2: {0, -1},
	3: {-1, 0},
}

func move2(s *State) {
	for _, m := range s.Moves {
		if m.Kind == "F" {
			moveForward2(s, m)
		}
		if m.Kind == "R" {
			s.Dir = (s.Dir + 1) % 4
		}
		if m.Kind == "L" {
			s.Dir = (s.Dir + 3) % 4
		}
		// fmt.Printf("P1: Row=%d Col=%d Face=%d\n", s.Row, s.Col, s.Dir)
	}
}

func move(s *State) {
	for _, m := range s.Moves {
		if m.Kind == "F" {
			moveForward(s, m)
		}
		if m.Kind == "R" {
			s.Dir = (s.Dir + 1) % 4
		}
		if m.Kind == "L" {
			s.Dir = (s.Dir + 3) % 4
		}
		// fmt.Printf("P1: Row=%d Col=%d Face=%d\n", s.Row, s.Col, s.Dir)
	}
}

func moveForward2(s *State, m Move) {
	for cnt := 0; cnt < m.Value; cnt++ {

		newRow, newCol, newDir := moveOnCube(s.Dir, s.Row, s.Col)
		fmt.Println(s.Dir, s.Row, s.Col, "|", newRow, newCol, newDir)

		if s.Field[newRow][newCol] == '.' {
			// fmt.Println(newRow, newCol, newDir)
			s.Row = newRow
			s.Col = newCol
			s.Dir = newDir
		} else {
			break
		}
	}
}

func moveOnCube(dir, row, col int) (int, int, int) {

	if dir == 3 && row == 0 && col >= 50 && col <= 99 {
		return col + 100, 0, 0
	}
	if dir == 2 && row >= 150 && row <= 199 && col == 0 {
		return 0, row - 100, 1
	}

	if dir == 3 && row == 0 && col >= 100 && col <= 149 {
		return 199, col - 100, 3
	}
	if dir == 1 && row == 199 && col >= 0 && col <= 49 {
		return 0, col + 100, 1
	}

	if dir == 0 && row >= 0 && row <= 49 && col == 149 {
		return 149 - row, 99, 2
	}
	if dir == 0 && row >= 100 && row <= 149 && col == 99 {
		return 149 - row, 149, 2
	}

	if dir == 1 && row == 49 && col >= 100 && col <= 149 {
		return col - 50, 99, 2
	}
	if dir == 0 && row >= 50 && row <= 99 && col == 99 {
		return 49, row + 50, 3
	}

	if dir == 1 && row == 149 && col >= 50 && col <= 99 {
		return col + 100, 49, 2
	}
	if dir == 0 && row >= 150 && row <= 199 && col == 49 {
		return 149, row - 100, 3
	}

	if dir == 2 && row >= 0 && row <= 49 && col == 50 {
		return 149 - row, 0, 0
	}
	if dir == 2 && row >= 100 && row <= 149 && col == 0 {
		return 149 - row, 50, 0
	}

	if dir == 2 && row >= 50 && row <= 99 && col == 50 {
		return 100, row - 50, 1
	}
	if dir == 3 && row == 100 && col >= 0 && col <= 49 {
		return col + 50, 50, 0
	}

	delta := dirToDeltaRowCol[dir]
	return row + delta[0], col + delta[1], dir
}

func moveForward(s *State, m Move) {
	delta := dirToDeltaRowCol[s.Dir]

	for cnt := 0; cnt < m.Value; cnt++ {
		var (
			newRow int
			newCol int
		)

		if s.Dir == 0 && s.Col == s.RowLR[s.Row].V2 {
			fmt.Println("T", s.Dir, s.Row, s.Col, s.RowLR[s.Row])
			newCol = s.RowLR[s.Row].V1
		} else if s.Dir == 2 && s.Col == s.RowLR[s.Row].V1 {
			newCol = s.RowLR[s.Row].V2
		} else {
			newCol = s.Col + delta[1]
		}

		if s.Dir == 1 && s.Row == s.ColTB[s.Col].V2 {
			newRow = s.ColTB[s.Col].V1
		} else if s.Dir == 3 && s.Row == s.ColTB[s.Col].V1 {
			newRow = s.ColTB[s.Col].V2
		} else {
			newRow = s.Row + delta[0]
		}

		if s.Field[newRow][newCol] == '.' {
			fmt.Println(newRow, newCol)
			s.Row = newRow
			s.Col = newCol
		} else {
			break
		}
	}
}

func printState(s State) {
	for row := 0; row < len(s.Field); row++ {
		for col := 0; col < len(s.Field[row]); col++ {
			if s.Field[row][col] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(string(s.Field[row][col]))
			}
		}
		fmt.Println()
	}
	fmt.Println(s.Moves)
	fmt.Println(s.RowLR)
	fmt.Println(s.ColTB)
}

type Move struct {
	Kind  string
	Value int
}

type V1V2 struct {
	V1 int
	V2 int
}

type State struct {
	Field [][]byte
	RowLR []V1V2
	ColTB []V1V2
	Row   int
	Col   int
	Dir   int
	Moves []Move
}

func loadInput() State {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	temp := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		temp = append(temp, line)
	}

	LL := len(temp) - 2

	maxCol := 0
	for i := 0; i < LL; i++ {
		if maxCol < len(temp[i]) {
			maxCol = len(temp[i])
		}
	}

	res := State{}
	res.Field = make([][]byte, LL)
	for i := 0; i < LL; i++ {
		res.Field[i] = make([]byte, maxCol)
		for j := 0; j < len(temp[i]); j++ {
			if temp[i][j] != ' ' {
				res.Field[i][j] = temp[i][j]
			}
		}
	}

	moves := temp[LL+1]
	test := ""
	res.Moves = []Move{}
	for _, part := range moves {
		if (part == 'R' || part == 'L') && len(test) > 0 {
			v, _ := strconv.Atoi(test)
			res.Moves = append(res.Moves, Move{Kind: "F", Value: v})
			test = ""
		}

		test = test + string(part)

		if test == "R" {
			res.Moves = append(res.Moves, Move{Kind: "R", Value: 0})
			test = ""
		} else if test == "L" {
			res.Moves = append(res.Moves, Move{Kind: "L", Value: 0})
			test = ""
		}
	}

	if len(test) > 0 {
		v, _ := strconv.Atoi(test)
		res.Moves = append(res.Moves, Move{Kind: "F", Value: v})
	}

	res.RowLR = make([]V1V2, LL)
	for row := 0; row < LL; row++ {
		v1 := -1
		v2 := -1
		for col := 0; col < maxCol; col++ {
			if res.Field[row][col] != 0 && v1 == -1 {
				v1 = col
			}
			if (res.Field[row][col] == 0) && v1 != -1 && v2 == -1 {
				v2 = col - 1
			}
			if col == maxCol-1 && v1 != -1 && v2 == -1 {
				v2 = col
			}
		}
		res.RowLR[row] = V1V2{V1: v1, V2: v2}
	}

	res.ColTB = make([]V1V2, maxCol)
	for col := 0; col < maxCol; col++ {
		v1 := -1
		v2 := -1
		for row := 0; row < LL; row++ {
			if res.Field[row][col] != 0 && v1 == -1 {
				v1 = row
			}
			if res.Field[row][col] == 0 && v1 != -1 && v2 == -1 {
				v2 = row - 1
			}
			if row == LL-1 && v1 != -1 && v2 == -1 {
				v2 = row
			}
		}
		res.ColTB[col] = V1V2{V1: v1, V2: v2}
	}

	res.Row = 0
	res.Col = res.RowLR[0].V1
	res.Dir = 0 //right-0 down-1 left-2 up-3

	fmt.Println(LL, maxCol)

	return res
}
