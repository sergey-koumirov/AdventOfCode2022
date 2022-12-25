package main

import (
	"bufio"
	"fmt"
	"os"
)

const L = 99

func main() {
	part1()
	part2()
}

func part1() {
	state := State{}
	initState(&state)

	topToBottom(&state)
	bottomToTop(&state)
	leftToRight(&state)
	rightToLeft(&state)

	fmt.Println("P1:", L*L-state.Invisible)
}

func part2() {
	state := State{}
	initState(&state)

	calcScore(&state)

	fmt.Println("P2:", state.BestScore)

}

func calcScore(state *State) {
	for row := 1; row < L-1; row++ {
		for col := 1; col < L-1; col++ {

			scoreUp := 0
			for i := row - 1; i >= 0; i-- {
				scoreUp += 1
				if state.F[row][col] <= state.F[i][col] {
					break
				}
			}

			scoreDown := 0
			for i := row + 1; i <= L-1; i++ {
				scoreDown += 1
				if state.F[row][col] <= state.F[i][col] {
					break
				}
			}

			scoreLeft := 0
			for i := col - 1; i >= 0; i-- {
				scoreLeft += 1
				if state.F[row][col] <= state.F[row][i] {
					break
				}
			}

			scoreRight := 0
			for i := col + 1; i <= L-1; i++ {
				scoreRight += 1
				if state.F[row][col] <= state.F[row][i] {
					break
				}
			}

			score := scoreUp * scoreDown * scoreLeft * scoreRight

			if state.BestScore < score {
				state.BestScore = score
			}
		}
	}
}

func rightToLeft(state *State) {
	for row := 1; row < L-1; row++ {
		maxIndex := L - 1
		for col := L - 2; col > state.RowMaxsI[row]; col-- {
			if state.F[row][col] > state.F[row][maxIndex] {
				if state.V[row][col] == 0 {
					state.V[row][col] = 1
					state.Invisible -= 1
				}
				maxIndex = col
			}
		}
	}
}

func leftToRight(state *State) {
	for row := 1; row < L-1; row++ {
		maxIndex := 0
		for col := 1; col < L-1; col++ {
			if state.F[row][col] > state.F[row][maxIndex] {
				if state.V[row][col] == 0 {
					state.V[row][col] = 1
					state.Invisible -= 1
				}
				maxIndex = col
			}
		}
		state.RowMaxsI[row] = maxIndex
	}
}

func bottomToTop(state *State) {
	for col := 1; col < L-1; col++ {
		maxIndex := L - 1
		for row := L - 2; row > state.ColMaxsI[col]; row-- {
			if state.F[row][col] > state.F[maxIndex][col] {
				if state.V[row][col] == 0 {
					state.V[row][col] = 1
					state.Invisible -= 1
				}
				maxIndex = row
			}
		}
	}
}

func topToBottom(state *State) {
	for col := 1; col < L-1; col++ {
		maxIndex := 0
		for row := 1; row < L-1; row++ {
			if state.F[row][col] > state.F[maxIndex][col] {
				if state.V[row][col] == 0 {
					state.V[row][col] = 1
					state.Invisible -= 1
				}
				maxIndex = row
			}
		}
		state.ColMaxsI[col] = maxIndex
	}
}

func initState(state *State) {
	state.F = loadInput()
	state.Invisible = (L - 2) * (L - 2)
	state.V = Field{}
	state.V = make([][]int, L)
	for i := 0; i < L; i++ {
		state.V[i] = make([]int, L)
	}
	state.RowMaxsI = make([]int, L)
	state.ColMaxsI = make([]int, L)
}

type State struct {
	F         Field
	V         Field
	RowMaxsI  []int
	ColMaxsI  []int
	Invisible int
	BestScore int
}

type Field [][]int

func loadInput() Field {
	res := Field{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		temp := make([]int, L)
		for i, c := range line {
			temp[i] = int(c) - 48
		}

		res = append(res, temp)
	}

	return res
}
