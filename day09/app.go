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
	part1()
	part2()
}

func part1() {
	moves := loadInput()
	state := State{}
	state.Visited = VisitedCells{}
	setVisited(&state, int64(0), int64(0))
	state.Points = []XY{{X: 0, Y: 0}, {X: 0, Y: 0}}

	emulateMoves(moves, &state)

	sum := 0
	for _, v := range state.Visited {
		sum += len(v)
	}

	fmt.Println("P1:", sum) // 6269
}

func part2() {
	moves := loadInput()
	state := State{}
	state.Visited = VisitedCells{}
	setVisited(&state, int64(0), int64(0))
	state.Points = []XY{{X: 0, Y: 0}, {X: 0, Y: 0}, {X: 0, Y: 0}, {X: 0, Y: 0}, {X: 0, Y: 0}, {X: 0, Y: 0}, {X: 0, Y: 0}, {X: 0, Y: 0}, {X: 0, Y: 0}, {X: 0, Y: 0}}

	emulateMoves(moves, &state)

	sum := 0
	for _, v := range state.Visited {
		sum += len(v)
	}
	fmt.Println("P2:", sum)
}

func emulateMoves(moves []Move, state *State) {
	for _, m := range moves {
		// fmt.Println(m)
		if m.Dir == "U" {
			applyMove(0, m.Cnt, state, 0)
		}
		if m.Dir == "D" {
			applyMove(0, -m.Cnt, state, 0)
		}
		if m.Dir == "L" {
			applyMove(-m.Cnt, 0, state, 0)
		}
		if m.Dir == "R" {
			applyMove(m.Cnt, 0, state, 0)
		}
		// fmt.Println()
		// printState(state, false)
	}
	// printState(state, true)
}

func applyMove(dx, dy int64, state *State, headIndex int) {
	h := &state.Points[headIndex]

	// fmt.Println("Begin:", h, dx, dy)

	if headIndex+1 == len(state.Points) {
		if dx == 0 {
			signY := int64(1)
			if math.Signbit(float64(dy)) {
				signY = -1
			}

			for y := h.Y + signY; y != h.Y+dy; y += signY {
				setVisited(state, h.X, y)
			}
		} else if dy == 0 {
			signX := int64(1)
			if math.Signbit(float64(dx)) {
				signX = -1
			}

			for x := h.X + signX; x != h.X+dx; x += signX {
				setVisited(state, x, h.Y)
			}
		}

		h.X += dx
		h.Y += dy
		setVisited(state, h.X, h.Y)
		return
	}

	h.X += dx
	h.Y += dy

	t := &state.Points[headIndex+1]
	// fmt.Println("  Tail:", t)

	if math.Abs(float64(h.X-t.X)) <= 1 && math.Abs(float64(h.Y-t.Y)) <= 1 {
		return
	}

	diagX := int64(1)
	diagY := int64(1)

	if math.Signbit(float64(h.X - t.X)) {
		diagX = -1
	}

	if math.Signbit(float64(h.Y - t.Y)) {
		diagY = -1
	}

	// diagonal move
	if h.X != t.X && h.Y != t.Y {
		applyMove(diagX, diagY, state, headIndex+1)
	}

	if math.Abs(float64(h.X-t.X)) <= 1 && math.Abs(float64(h.Y-t.Y)) <= 1 {
		return
	}

	// catch head by line
	if h.X == t.X && h.Y != t.Y {
		applyMove(0, h.Y-t.Y-diagY, state, headIndex+1)
	} else if h.X != t.X && h.Y == t.Y {
		applyMove(h.X-t.X-diagX, 0, state, headIndex+1)
	} else {
		fmt.Println("PANIC !!!", "T:", h.X, h.Y, "H:", t.X, t.Y)
	}
}

func printState(state *State, showVisited bool) {
	maxX := int64(-1000)
	minX := int64(1000)
	maxY := int64(-1000)
	minY := int64(1000)

	for x, yy := range state.Visited {
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
		for y, _ := range yy {
			if y > maxY {
				maxY = y
			}
			if y < minY {
				minY = y
			}
		}
	}
	for _, p := range state.Points {
		if p.X > maxX {
			maxX = p.X
		}
		if p.X < minX {
			minX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
		if p.Y < minY {
			minY = p.Y
		}
	}

	fmt.Println("-----------------------------")

	if showVisited {
		for y := maxY; y >= minY; y-- {
			for x := minX; x <= maxX; x++ {
				_, ex := state.Visited[x][y]
				if x == 0 && y == 0 {
					fmt.Print("s")
				} else if ex {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}

		fmt.Println()
	}

	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			ex := false

			for ip, p := range state.Points {
				if p.X == x && p.Y == y {
					ex = true
					if ip == 0 {
						fmt.Print("H")
					} else {
						fmt.Print(ip)
					}
					break
				}
			}

			if !ex && x == 0 && y == 0 {
				fmt.Print("s")
			} else if !ex {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func setVisited(state *State, x, y int64) {
	// fmt.Println("V:", x, y)
	vv, ex := state.Visited[x]
	if !ex {
		vv = map[int64]int{}
	}
	vv[y] = 1
	state.Visited[x] = vv
}

type XY struct {
	X int64
	Y int64
}

type State struct {
	Points  []XY
	Visited VisitedCells
}

type Move struct {
	Dir string
	Cnt int64
}

type VisitedCells map[int64]map[int64]int

func loadInput() []Move {
	res := []Move{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		cnt, _ := strconv.ParseInt(parts[1], 10, 64)

		res = append(res, Move{Dir: parts[0], Cnt: cnt})
	}

	return res
}
