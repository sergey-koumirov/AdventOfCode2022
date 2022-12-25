package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part 1:")
	part1()
	fmt.Println("Part 2:")
	part2()
}

func part1() {
	stacks, moves := loadInput()

	for _, m := range moves {
		for i := int64(0); i < m.Cnt; i++ {
			v := stacks[m.From].Remove(stacks[m.From].Front())
			stacks[m.To].PushFront(v)
		}
	}

	for i := 0; i <= 8; i++ {
		fmt.Print(stacks[i].Front().Value)
	}
	fmt.Println()
}

func part2() {
	stacks, moves := loadInput()

	for _, m := range moves {
		temp := []any{}
		for i := int64(0); i < m.Cnt; i++ {
			temp = append(temp, stacks[m.From].Remove(stacks[m.From].Front()))
		}
		for i := len(temp) - 1; i >= 0; i-- {
			stacks[m.To].PushFront(temp[i])
		}
	}

	for i := 0; i <= 8; i++ {
		fmt.Print(stacks[i].Front().Value)
	}
	fmt.Println()
}

type Move struct {
	From int64
	To   int64
	Cnt  int64
}

func loadInput() ([]*list.List, []Move) {
	stacks := []*list.List{}
	for i := 0; i <= 8; i++ {
		stacks = append(stacks, list.New())
	}
	moves := []Move{}

	firstPart := true

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			firstPart = false
		} else {
			if firstPart {
				for i := 0; i <= 8; i++ {
					sub := line[i*4+1 : i*4+2]
					if sub != " " && line[0:1] == "[" {
						stacks[i].PushBack(sub)
					}
				}
			} else {
				temp := Move{}
				s1 := strings.Split(line[5:], " from ")
				temp.Cnt, _ = strconv.ParseInt(s1[0], 10, 64)
				s2 := strings.Split(s1[1], " to ")
				temp.From, _ = strconv.ParseInt(s2[0], 10, 64)
				temp.From -= 1
				temp.To, _ = strconv.ParseInt(s2[1], 10, 64)
				temp.To -= 1
				moves = append(moves, temp)
			}
		}
	}

	return stacks, moves
}
