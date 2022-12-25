package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part2() {
	pp := loadInput()
	sum := 0
	for _, p := range pp {
		if overlap(p) {
			sum += 1
		}
	}
	fmt.Println(sum)

}

func part1() {
	pp := loadInput()
	sum := 0
	for _, p := range pp {
		if contain(p) {
			sum += 1
		}
	}
	fmt.Println(sum)
}

func overlap(p Pairs) bool {
	return p.Second1 <= p.First2 && p.First1 <= p.Second2
}

func contain(p Pairs) bool {
	return p.First1 <= p.Second1 && p.Second2 <= p.First2 || p.Second1 <= p.First1 && p.First2 <= p.Second2
}

type Pairs struct {
	First1  int64
	First2  int64
	Second1 int64
	Second2 int64
}

func loadInput() []Pairs {

	res := []Pairs{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		ff := strings.Split(parts[0], "-")
		f1, _ := strconv.ParseInt(ff[0], 10, 64)
		f2, _ := strconv.ParseInt(ff[1], 10, 64)

		ss := strings.Split(parts[1], "-")
		s1, _ := strconv.ParseInt(ss[0], 10, 64)
		s2, _ := strconv.ParseInt(ss[1], 10, 64)

		res = append(res, Pairs{First1: f1, First2: f2, Second1: s1, Second2: s2})
	}

	return res
}
