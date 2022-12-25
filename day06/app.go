package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1()
	part2()
}

func part1() {
	line := loadInput()

	lineLen := len(line)

	res := -1

	for i := 3; i < lineLen; i++ {
		if uniq4(line[i-3], line[i-2], line[i-1], line[i]) {
			res = i + 1
			break
		}
	}

	fmt.Println("P1", res)
}

func uniq4(c1, c2, c3, c4 byte) bool {
	hasSame := c1 == c2 || c1 == c3 || c1 == c4 || c2 == c3 || c2 == c4 || c3 == c4
	return !hasSame
}

func uniqN(s string) bool {
	L := len(s)

	for i := 1; i < L; i++ {
		for j := i + 1; j < L; j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}

	return true
}

func part2() {
	line := loadInput()

	lineLen := len(line)

	res := -1

	for i := 14; i < lineLen; i++ {
		if uniqN(line[i-14 : i+1]) {
			res = i + 1
			break
		}
	}

	fmt.Println("P2", res)
}

func loadInput() string {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	l := scanner.Text()
	return l
}
