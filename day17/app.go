package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var AllFigures = [5]Figure{
	{Coords: []Delta{{DX: 0, DY: 0}, {DX: 1, DY: 0}, {DX: 2, DY: 0}, {DX: 3, DY: 0}}, Height: 1, Width: 4},
	{Coords: []Delta{{DX: 1, DY: 0}, {DX: 0, DY: 1}, {DX: 1, DY: 1}, {DX: 2, DY: 1}, {DX: 1, DY: 2}}, Height: 3, Width: 3},
	{Coords: []Delta{{DX: 0, DY: 0}, {DX: 1, DY: 0}, {DX: 2, DY: 0}, {DX: 2, DY: 1}, {DX: 2, DY: 2}}, Height: 3, Width: 3},
	{Coords: []Delta{{DX: 0, DY: 0}, {DX: 0, DY: 1}, {DX: 0, DY: 2}, {DX: 0, DY: 3}}, Height: 4, Width: 1},
	{Coords: []Delta{{DX: 0, DY: 0}, {DX: 0, DY: 1}, {DX: 1, DY: 0}, {DX: 1, DY: 1}}, Height: 2, Width: 2},
}

func main() {
	start := time.Now()
	// part1()
	part2()
	elapsed := time.Since(start)
	fmt.Println("time(ms):", elapsed.Milliseconds())
}

func part1() {
	state := loadInput()
	for r := 1; r <= 2022; r++ {
		genereateRock(&state)
		emulateFalling(&state)
	}
	fmt.Println("P1:", state.RockHigh)
}

func part2() {
	state := loadInput()

	// 1000_000_000_000

	first := 1768
	cycle := 1715
	cyclePlus := 2690

	upper := (1000_000_000_000 - first) % cycle
	cycleCnt := (1000_000_000_000 - first) / cycle

	test := map[string]int{}

	max := 0
	for r := 1; r <= first+upper; r++ {
		genereateRock(&state)

		key := strconv.Itoa(state.FigIndex) + "-" + strconv.Itoa(state.Tick%state.WindsLen) + "-" + toKey(&state)

		_, ex := test[key]
		if ex {
			test[key]++
		} else {
			test[key] = 1
		}

		if key == "2-150-7d-78-7d-61-59-59-40-48-8-6a" {
			fmt.Println(r, state.RockHigh, "key", key)
			printMap(&state, 30)
		}

		moves := emulateFalling(&state)

		if moves > max {
			max = moves
		}
	}
	// fmt.Println(test)
	// printMap(&state, 30)
	// fmt.Println("P2:", state.RockHigh, max, state.WindsLen)

	// for k, cnt := range test {
	// 	if cnt > 1 {
	// 		fmt.Println(k)
	// 	}
	// }

	// 1_514_285_714_288
	// 1_514_285_714_288
	fmt.Println("Res", cycleCnt*cyclePlus+state.RockHigh)

}

func emulateFalling(s *State) int {
	moves := 0
	for {
		w := s.Winds[s.Tick%s.WindsLen]
		s.Tick++
		applyWind(s, w)
		halt := checkDownward(s)
		if halt {
			applyFigure(s)
			break
		}
		moves++
	}
	return moves
}

func applyFigure(s *State) {
	f := AllFigures[s.FigIndex]
	for i := 0; i < len(f.Coords); i++ {
		p := f.Coords[i]
		x := s.FigX + p.DX
		y := s.FigY + p.DY
		s.Map[y][x] = '#'
	}

	if s.RockHigh < s.FigY+f.Height {
		s.RockHigh = s.FigY + f.Height
	}

	// fmt.Println("New Rock", s.RockHigh)
}

func checkDownward(s *State) bool {
	f := AllFigures[s.FigIndex]

	moved := true

	for i := 0; i < len(f.Coords); i++ {
		p := f.Coords[i]
		x := s.FigX + p.DX
		newY := s.FigY + p.DY - 1
		if newY < 0 || s.Map[newY][x] == '#' {
			moved = false
			break
		}
	}

	if moved {
		s.FigY--
	}

	return !moved
}

func applyWind(s *State, w byte) {
	dx := 1
	if w == '<' {
		dx = -1
	}

	f := AllFigures[s.FigIndex]

	moved := true

	for i := 0; i < len(f.Coords); i++ {
		p := f.Coords[i]
		newX := s.FigX + p.DX + dx
		y := s.FigY + p.DY
		if newX < 0 || newX > 6 || s.Map[y][newX] == '#' {
			moved = false
			break
		}
	}

	if moved {
		s.FigX += dx
	}
}

func printMap(s *State, top int) {
	f := AllFigures[s.FigIndex]

	L := len(s.Map)
	for row := L - 1; row > L-top && row >= 0; row-- {
		fmt.Print("+")
		for x := 0; x < 7; x++ {
			y := row
			fEx := false

			for i := 0; i < len(f.Coords); i++ {
				if s.FigX+f.Coords[i].DX == x && s.FigY+f.Coords[i].DY == y {
					fEx = true
				}
			}

			if fEx {
				fmt.Print("@")
			} else if s.Map[y][x] == '#' {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println("+")
	}
	fmt.Println("+-------+")
	fmt.Println()
}

func genereateRock(s *State) {
	s.FigIndex = (s.FigIndex + 1) % 5
	needHigh := s.RockHigh + AllFigures[s.FigIndex].Height + 3

	if len(s.Map) < needHigh {
		for i := 0; i < 10; i++ {
			s.Map = append(s.Map, Line{})
		}
	}

	s.FigX = 2
	s.FigY = s.RockHigh + 3
}

func toKey(state *State) string {
	bitMask := [10]int64{}

	for i := 0; i < 10; i++ {
		if state.RockHigh-i-1 >= 0 {
			if state.Map[state.RockHigh-i-1][0] == 0 {
				bitMask[i] = bitMask[i] | 0b1000000
			}
			if state.Map[state.RockHigh-i-1][1] == 0 {
				bitMask[i] = bitMask[i] | 0b100000
			}
			if state.Map[state.RockHigh-i-1][2] == 0 {
				bitMask[i] = bitMask[i] | 0b10000
			}
			if state.Map[state.RockHigh-i-1][3] == 0 {
				bitMask[i] = bitMask[i] | 0b1000
			}
			if state.Map[state.RockHigh-i-1][4] == 0 {
				bitMask[i] = bitMask[i] | 0b100
			}
			if state.Map[state.RockHigh-i-1][5] == 0 {
				bitMask[i] = bitMask[i] | 0b10
			}
			if state.Map[state.RockHigh-i-1][6] == 0 {
				bitMask[i] = bitMask[i] | 0b1
			}
		}
	}

	temp := make([]string, 10)
	for i := 0; i < 10; i++ {
		temp[i] = strconv.FormatInt(bitMask[i], 16)
	}
	return strings.Join(temp, "-")
}

type Line [7]byte

type State struct {
	Map      []Line
	Winds    string
	WindsLen int
	Tick     int
	FigIndex int
	FigX     int
	FigY     int
	RockHigh int
}

type Delta struct {
	DX int
	DY int
}

type Figure struct {
	Coords []Delta
	Width  int
	Height int
}

func loadInput() State {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()

	state := State{}
	state.Winds = line
	state.WindsLen = len(line)
	state.Map = []Line{}
	state.FigIndex = -1
	for i := 0; i < 10; i++ {
		state.Map = append(state.Map, Line{})
	}

	return state
}
