package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// part1()
	part2()
}

func part1() {
	field := loadInput()

	// printVisited(&field)
	// fmt.Println(field)

	step := runWave(&field)
	fmt.Println("P1:", step)

}

func part2() {
	field := loadInput()
	step := runBackWave(&field)
	fmt.Println("P2:", step)
}

func runBackWave(f *Field) int {

	f.Visited[f.End.Row][f.End.Col] = 1
	current := []Point{f.End}

	step := 0

	deltas := [4][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	for {
		nextWave := []Point{}

		for _, c := range current {

			for _, dd := range deltas {
				newRow := c.Row + dd[0]
				newCol := c.Col + dd[1]
				inside := newRow >= 0 && newRow <= f.RowSize-1 && newCol >= 0 && newCol <= f.ColSize-1

				if inside {
					notVisted := f.Visited[newRow][newCol] == 0

					dh := f.Cells[newRow][newCol] - f.Cells[c.Row][c.Col]

					if notVisted && dh >= -1 {
						if f.Cells[newRow][newCol] == 0 {
							return step + 1
						}
						f.Visited[newRow][newCol] = 1
						nextWave = append(nextWave, Point{Row: newRow, Col: newCol})

					}
				}
			}
		}

		// printVisited(f)
		// fmt.Scanln()

		if len(nextWave) == 0 {
			break
		}

		step++

		current = nextWave
	}
	// fmt.Println(step)
	return -1
}

func runWave(f *Field) int {

	f.Visited[f.Start.Row][f.Start.Col] = 1
	current := []Point{f.Start}

	step := 0

	deltas := [4][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	for {
		nextWave := []Point{}

		for _, c := range current {

			for _, dd := range deltas {
				newRow := c.Row + dd[0]
				newCol := c.Col + dd[1]
				inside := newRow >= 0 && newRow <= f.RowSize-1 && newCol >= 0 && newCol <= f.ColSize-1

				if inside {
					notVisted := f.Visited[newRow][newCol] == 0
					dh := f.Cells[newRow][newCol] - f.Cells[c.Row][c.Col]

					if notVisted && dh <= 1 {
						if f.End.Row == newRow && f.End.Col == newCol {
							return step + 1
						}
						f.Visited[newRow][newCol] = 1
						nextWave = append(nextWave, Point{Row: newRow, Col: newCol})

					}
				}
			}
		}

		// printVisited(f)
		// fmt.Scanln()

		if len(nextWave) == 0 {
			break
		}

		step++

		current = nextWave
	}
	// fmt.Println(step)
	return -1
}

func printVisited(f *Field) {
	for ri, row := range f.Visited {
		for ci, v := range row {

			if ri == f.Start.Row && ci == f.Start.Col {
				fmt.Print("S")
			} else if ri == f.End.Row && ci == f.End.Col {
				fmt.Print("E")
			} else if v == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(string(f.Cells[ri][ci] + 97))
			}
		}
		fmt.Println()
	}
}

type Point struct {
	Col int
	Row int
}

type Field struct {
	ColSize int
	RowSize int
	Cells   [][]int
	Visited [][]int
	Start   Point
	End     Point
}

func loadInput() Field {
	res := Field{
		Cells:   [][]int{},
		Visited: [][]int{},
	}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if res.ColSize == 0 {
			res.ColSize = len(line)
		}

		temp := make([]int, res.ColSize)

		for i, c := range line {

			if c == 'S' {
				res.Start = Point{
					Col: i,
					Row: len(res.Cells),
				}
				c = 'a'
			}

			if c == 'E' {
				res.End = Point{
					Col: i,
					Row: len(res.Cells),
				}
				c = 'z'
			}

			temp[i] = int(c) - 97
		}

		res.Cells = append(res.Cells, temp)

		tempVis := make([]int, res.ColSize)
		res.Visited = append(res.Visited, tempVis)
	}

	res.RowSize = len(res.Cells)

	return res
}
