package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part2()
}

func part1() {
	gg := loadInput()

	res := 0

	for _, g := range gg {
		res += calcScore(g)
	}

	fmt.Println(res)
}

func part2() {
	gg := loadInput()

	res := 0

	for _, g := range gg {
		res += calcScore2(g)
	}

	fmt.Println(res)
}

// A for Rock,
// B for Paper, and
// C for Scissors.

// X for Rock,
// Y for Paper, and
// Z for Scissors

// Anyway, the second column says how the round needs to end:
// X means you need to lose,
// Y means you need to end the round in a draw, and
// Z means you need to win. Good luck!
func calcScore2(g Game) int {
	you := ""

	// lose
	if g.You == "X" {
		if g.Elf == "A" {
			you = "Z"
		}
		if g.Elf == "B" {
			you = "X"
		}
		if g.Elf == "C" {
			you = "Y"
		}
	}

	// draw
	if g.You == "Y" {
		if g.Elf == "A" {
			you = "X"
		}
		if g.Elf == "B" {
			you = "Y"
		}
		if g.Elf == "C" {
			you = "Z"
		}
	}

	// win
	if g.You == "Z" {
		if g.Elf == "A" {
			you = "Y"
		}
		if g.Elf == "B" {
			you = "Z"
		}
		if g.Elf == "C" {
			you = "X"
		}
	}

	return calcScore(Game{Elf: g.Elf, You: you})
}

func calcScore(g Game) int {
	score := 0

	if g.You == "X" {
		score += 1
	} else if g.You == "Y" {
		score += 2
	} else if g.You == "Z" {
		score += 3
	}

	if g.Elf == "A" && g.You == "X" || g.Elf == "B" && g.You == "Y" || g.Elf == "C" && g.You == "Z" {
		score += 3
	}

	if g.Elf == "A" && g.You == "Y" || g.Elf == "B" && g.You == "Z" || g.Elf == "C" && g.You == "X" {
		score += 6
	}

	return score
}

type Game struct {
	Elf string
	You string
}

func loadInput() []Game {
	result := []Game{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		result = append(result, Game{Elf: parts[0], You: parts[1]})
	}

	return result
}
