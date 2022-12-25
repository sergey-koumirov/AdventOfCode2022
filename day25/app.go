package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	part1()
	// part2()
	elapsed := time.Since(start)
	fmt.Println("time(ms):", elapsed.Milliseconds())
}

func part1() {
	numbers := loadInput()

	sum := 0
	for _, n5 := range numbers {
		sum += b5b10(n5)
	}

	fmt.Println("P1:", sum, b10b5s(sum))
}

func part2() {
	loadInput()
	fmt.Println("P2:")
}

func b10b5s(n10 int) string {

	test := n10
	p := 5
	b5 := make([]int, 0)

	for {
		b5 = append(b5, test%p)
		test = test / p
		if test == 0 {
			break
		}
	}

	// fmt.Println("B5:", b5, strconv.FormatInt(int64(n10), 5))
	// [0 4 3 3 4 3]

	index := 0
	for {
		// fmt.Println("B5:", index, b5)

		if b5[index] > 2 {
			if b5[index] == 3 {
				b5[index] = -2
			}
			if b5[index] == 4 {
				b5[index] = -1
			}

			overIndex := index + 1
			for {
				if overIndex > len(b5)-1 {
					b5 = append(b5, 0)
				}
				if b5[overIndex-1] == 5 {
					b5[overIndex-1] = 0
				}
				b5[overIndex] = b5[overIndex] + 1
				if b5[overIndex] != 5 {
					break
				}

				overIndex++
			}
		}

		// fmt.Println("B5:", index, b5)
		// fmt.Println()

		if index == len(b5)-1 {
			break
		}
		index++
	}

	res := ""

	for i := 0; i < len(b5); i++ {
		if b5[i] == -2 {
			res = "=" + res
		}
		if b5[i] == -1 {
			res = "-" + res
		}
		if b5[i] == 0 {
			res = "0" + res
		}
		if b5[i] == 1 {
			res = "1" + res
		}
		if b5[i] == 2 {
			res = "2" + res
		}
	}

	return res
}

func b5b10(n5 []int) int {
	n10 := 0

	p := 1

	L := len(n5)

	for i := 0; i < L; i++ {
		n10 = n10 + p*n5[L-i-1]
		p = p * 5
	}

	return n10
}

func loadInput() [][]int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	res := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		temp := make([]int, len(line))

		for i, c := range line {
			if c == '=' {
				temp[i] = -2
			} else if c == '-' {
				temp[i] = -1
			} else if c == '0' {
				temp[i] = 0
			} else if c == '1' {
				temp[i] = 1
			} else if c == '2' {
				temp[i] = 2
			}
		}

		res = append(res, temp)
	}

	return res
}
