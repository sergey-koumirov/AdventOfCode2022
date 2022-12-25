package main

import (
	"bufio"
	"fmt"
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
	printField(&state)

	state.Current.X = state.Init.X
	state.Current.Y = state.Init.Y
	for {
		void, rest := emulateStep(&state)
		// printField(&state)
		// fmt.Println()
		if void {
			break
		}
		if rest {
			state.Sand += 1
			state.Field[state.Current.Y][state.Current.X] = 2
			state.Current.X = state.Init.X
			state.Current.Y = state.Init.Y
		}
	}

	fmt.Println("P1:", state.Sand)
	// fmt.Printf("%+v\n", state)
}

func part2() {
	state := loadInput2()
	// printField(&state)
	// fmt.Println()

	state.Current.X = state.Init.X
	state.Current.Y = state.Init.Y
	for {
		rest := emulateStep2(&state)
		if state.Current.Y == state.Init.Y && state.Current.X == state.Init.X {
			state.Sand += 1
			// printField(&state)
			break
		}
		if rest {
			state.Sand += 1
			state.Field[state.Current.Y][state.Current.X] = 2
			state.Current.X = state.Init.X
			state.Current.Y = state.Init.Y
		}
	}

	fmt.Println("P2:", state.Sand)
}

func emulateStep2(state *State) bool {
	if state.Current.Y == state.YSize-1 {
		return true
	}

	if state.Field[state.Current.Y+1][state.Current.X] == 0 {
		state.Current.Y += 1
		return false
	}

	if state.Field[state.Current.Y+1][state.Current.X-1] == 0 {
		state.Current.X -= 1
		state.Current.Y += 1
		return false
	}

	if state.Field[state.Current.Y+1][state.Current.X+1] == 0 {
		state.Current.X += 1
		state.Current.Y += 1
		return false
	}
	return true
}

func emulateStep(state *State) (bool, bool) {
	if state.Current.Y == state.YSize-1 {
		return true, false
	}

	if state.Field[state.Current.Y+1][state.Current.X] == 0 {
		state.Current.Y += 1
		return false, false
	}

	if state.Current.X == 0 {
		return true, false
	}

	if state.Field[state.Current.Y+1][state.Current.X-1] == 0 {
		state.Current.X -= 1
		state.Current.Y += 1
		return false, false
	}

	if state.Current.X == state.XSize-1 {
		return true, false
	}

	if state.Field[state.Current.Y+1][state.Current.X+1] == 0 {
		state.Current.X += 1
		state.Current.Y += 1
		return false, false
	}
	return false, true
}

type Point struct {
	X int64
	Y int64
}

type State struct {
	Init    Point
	Current Point
	Field   [][]int64
	XSize   int64
	YSize   int64
	Sand    int64
}

func loadInput() State {
	res := State{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	temp := [][]Point{}

	minX := int64(9999)
	maxX := int64(0)

	minY := int64(9999)
	maxY := int64(0)

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			temp = append(temp, []Point{})
			parts := strings.Split(line, " -> ")
			for _, part := range parts {
				coords := strings.Split(part, ",")

				x, _ := strconv.ParseInt(coords[0], 10, 64)
				y, _ := strconv.ParseInt(coords[1], 10, 64)

				temp[len(temp)-1] = append(temp[len(temp)-1], Point{X: x, Y: y})
				applyMinMax(&minX, &maxX, &minY, &maxY, x, y)
			}
		}
	}

	applyMinMax(&minX, &maxX, &minY, &maxY, 500, 0)

	for i, vv := range temp {
		for j := range vv {
			temp[i][j].X -= minX
			temp[i][j].Y -= minY
		}
	}

	res.XSize = maxX - minX + 1
	res.YSize = maxY - minY + 1

	res.Init.X = 500 - minX
	res.Init.Y = 0 - minY

	res.Current.X = -1
	res.Current.Y = -1

	res.Field = make([][]int64, res.YSize)
	for i := int64(0); i < res.YSize; i++ {
		res.Field[i] = make([]int64, res.XSize)
	}

	for _, vv := range temp {
		for j := 1; j < len(vv); j++ {
			p1 := vv[j-1]
			p2 := vv[j]
			if p1.X == p2.X {
				dy := sign(p2.Y - p1.Y)
				for row := p1.Y; row != p2.Y+dy; row += dy {
					res.Field[row][p1.X] = 1
				}
			}
			if p1.Y == p2.Y {
				dx := sign(p2.X - p1.X)
				for col := p1.X; col != p2.X+dx; col += dx {
					res.Field[p1.Y][col] = 1
				}
			}
		}
	}

	return res
}

func loadInput2() State {
	res := State{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	temp := [][]Point{}

	minX := int64(99999)
	maxX := int64(0)

	minY := int64(99999)
	maxY := int64(0)

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			temp = append(temp, []Point{})
			parts := strings.Split(line, " -> ")
			for _, part := range parts {
				coords := strings.Split(part, ",")

				x, _ := strconv.ParseInt(coords[0], 10, 64)
				y, _ := strconv.ParseInt(coords[1], 10, 64)

				temp[len(temp)-1] = append(temp[len(temp)-1], Point{X: x, Y: y})
				applyMinMax(&minX, &maxX, &minY, &maxY, x, y)
			}
		}
	}

	applyMinMax(&minX, &maxX, &minY, &maxY, 500, 0)
	maxY += 1
	minX -= maxY + 10
	maxX += maxY + 10

	for i, vv := range temp {
		for j := range vv {
			temp[i][j].X -= minX
			temp[i][j].Y -= minY
		}
	}

	res.XSize = maxX - minX + 1
	res.YSize = maxY - minY + 1

	res.Init.X = 500 - minX
	res.Init.Y = 0 - minY

	res.Current.X = -1
	res.Current.Y = -1

	res.Field = make([][]int64, res.YSize)
	for i := int64(0); i < res.YSize; i++ {
		res.Field[i] = make([]int64, res.XSize)
	}

	for _, vv := range temp {
		for j := 1; j < len(vv); j++ {
			p1 := vv[j-1]
			p2 := vv[j]
			if p1.X == p2.X {
				dy := sign(p2.Y - p1.Y)
				for row := p1.Y; row != p2.Y+dy; row += dy {
					res.Field[row][p1.X] = 1
				}
			}
			if p1.Y == p2.Y {
				dx := sign(p2.X - p1.X)
				for col := p1.X; col != p2.X+dx; col += dx {
					res.Field[p1.Y][col] = 1
				}
			}
		}
	}

	return res
}

func sign(v int64) int64 {
	if v == 0 {
		return 0
	}
	if v > 0 {
		return 1
	}
	return -1
}

func printField(state *State) {
	for row, vv := range state.Field {
		for col, v := range vv {
			if row == int(state.Current.Y) && col == int(state.Current.X) {
				fmt.Print("O")
			} else if row == int(state.Init.Y) && col == int(state.Init.X) {
				fmt.Print("+")
			} else if v == 2 {
				fmt.Print("O")
			} else if v == 0 {
				fmt.Print(".")
			} else if v == 1 {
				fmt.Print("#")
			}

		}
		fmt.Println()
	}
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
