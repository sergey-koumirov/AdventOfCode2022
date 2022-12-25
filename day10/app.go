package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	// part2()
}

func part1() {
	cmds := loadInput()

	sum := emulateCycles(cmds)

	fmt.Println("P1:", sum)
}

func part2() {
	cmds := loadInput()
	fmt.Println("P2:", cmds)
}

func emulateCycles(cmds []Command) int {
	cycle := 0
	x := 1
	sum := 0

	screen := [6][40]int{}

	for _, cmd := range cmds {
		if cmd.Code == "noop" {
			cycle++
			calcSum(cycle, x, &sum, &screen)
		} else if cmd.Code == "addx" {
			cycle++
			calcSum(cycle, x, &sum, &screen)

			cycle++
			calcSum(cycle, x, &sum, &screen)

			x += int(cmd.Value)
		}
	}
	cycle++
	calcSum(cycle, x, &sum, &screen)

	for y := 1; y <= 6; y++ {
		for x := 1; x <= 40; x++ {
			if screen[y-1][x-1] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	return sum
}

func calcSum(cycle int, x int, s *int, screen *[6][40]int) {
	screenX := cycle % 40
	if screenX == 0 {
		screenX = 40
	}

	if x <= screenX && screenX <= x+2 {
		screen[cycle/40][(cycle-1)%40] = 1
	}

	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		*s += cycle * x
	}
}

type Command struct {
	Code  string
	Value int64
}

func loadInput() []Command {
	res := []Command{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		if len(parts) == 1 {
			res = append(res, Command{Code: parts[0]})
		} else {
			v, _ := strconv.ParseInt(parts[1], 10, 64)
			res = append(res, Command{Code: parts[0], Value: v})
		}
	}

	return res
}
