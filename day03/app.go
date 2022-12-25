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

func part2() {
	rucksaks := loadInput()

	sum := 0

	for i := 0; i < len(rucksaks); i += 3 {
		u0 := unitedKeys(rucksaks[i].First, rucksaks[i].Second)
		u1 := unitedKeys(rucksaks[i+1].First, rucksaks[i+1].Second)
		u2 := unitedKeys(rucksaks[i+2].First, rucksaks[i+2].Second)
		keys := commonKeys3(u0, u1, u2)
		for _, k := range keys {
			sum += k
		}
	}

	fmt.Println("L:", len(rucksaks), "sum:", sum)
}

func unitedKeys(f map[int]int, s map[int]int) map[int]int {
	res := map[int]int{}

	for k, v := range f {
		res[k] = v
	}

	for k, v := range s {
		res[k] += v
	}

	return res
}

func part1() {
	rucksaks := loadInput()

	sum := 0
	for _, r := range rucksaks {
		keys := commonKeys(r.First, r.Second)
		for _, k := range keys {
			sum += k
		}
	}

	fmt.Println("L:", len(rucksaks), "sum:", sum)
}

func commonKeys3(f map[int]int, s map[int]int, t map[int]int) []int {
	res := []int{}
	temp := map[int]int{}

	for k := range f {
		temp[k] += 1
	}

	for k := range s {
		temp[k] += 1
	}

	for k := range t {
		temp[k] += 1
	}

	for k, v := range temp {
		if v == 3 {
			res = append(res, k)
		}
	}

	return res
}

func commonKeys(f map[int]int, s map[int]int) []int {
	res := []int{}
	temp := map[int]int{}

	for k := range f {
		temp[k] = 1
	}

	for k := range s {
		_, ex := temp[k]
		if ex {
			res = append(res, k)
		}
	}

	return res
}

type Rucksack struct {
	First  map[int]int
	Second map[int]int
}

func loadInput() []Rucksack {
	result := []Rucksack{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		half := len(line) / 2

		first := map[int]int{}
		second := map[int]int{}

		for i, c := range line {
			if c >= 97 {
				c -= 96
			}
			if c >= 65 {
				c -= 38
			}

			if i < half {
				first[int(c)] += 1
			} else {
				second[int(c)] += 1
			}
		}

		result = append(result, Rucksack{First: first, Second: second})
	}

	return result
}
