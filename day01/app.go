package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// part1()
	part2()
}

func part2() {
	max1 := int64(0)
	max2 := int64(0)
	max3 := int64(0)
	temp := int64(0)

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		x, _ := strconv.ParseInt(line, 10, 64)

		if x == 0 {
			if max1 < temp {
				max3 = max2
				max2 = max1
				max1 = temp
			} else if max2 < temp {
				max3 = max2
				max2 = temp
			} else if max3 < temp {
				max3 = temp
			}
			temp = 0
		} else {
			temp = temp + x
		}
	}

	if max1 < temp {
		max3 = max2
		max2 = max1
		max1 = temp
	} else if max2 < temp {
		max3 = max2
		max2 = temp
	} else if max3 < temp {
		max3 = temp
	}

	fmt.Println("MAX123:", max1+max2+max3)

}

func part1() {
	max := int64(0)
	temp := int64(0)

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		x, _ := strconv.ParseInt(line, 10, 64)

		if x == 0 {
			if max < temp {
				max = temp
			}
			temp = 0
		} else {
			temp = temp + x
		}
	}

	if max < temp {
		max = temp
	}

	fmt.Println("MAX:", max)
}
