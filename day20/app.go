package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	// part1()
	part2()

	elapsed := time.Since(start)
	fmt.Println("time(ms):", elapsed.Milliseconds())
}

func part1() {
	nn := loadInput()

	numbers := findNumbers(nn)

	pos := -1
	for i, n := range numbers {
		if n.V == 0 {
			pos = i
			break
		}
	}

	L := len(numbers)
	fmt.Println("Pos:", pos)

	n1 := numbers[(pos+1000)%L]
	fmt.Println("1:", (pos+1000)%L, n1)

	n2 := numbers[(pos+2000)%L]
	fmt.Println("2:", (pos+2000)%L, n2)

	n3 := numbers[(pos+3000)%L]
	fmt.Println("3:", (pos+3000)%L, n3)

	fmt.Println("Sum:", n1.V+n2.V+n3.V)

	fmt.Println("P1:", L)
}

func part2() {
	nn := loadInput()

	numbers := findNumbers2(nn)

	pos := -1
	for i, n := range numbers {
		if n.V == 0 {
			pos = i
			break
		}
	}

	L := len(numbers)
	fmt.Println("Pos:", pos)

	n1 := numbers[(pos+1000)%L]
	fmt.Println("1:", (pos+1000)%L, n1)

	n2 := numbers[(pos+2000)%L]
	fmt.Println("2:", (pos+2000)%L, n2)

	n3 := numbers[(pos+3000)%L]
	fmt.Println("3:", (pos+3000)%L, n3)

	fmt.Println("Sum:", n1.V+n2.V+n3.V)

	fmt.Println("P2:", L)
}

func findNumbers2(nn []Number) []Number {
	L := len(nn)
	temp := make([]Number, L)
	for i := 0; i < L; i++ {
		nn[i].V *= 811589153
		temp[i] = nn[i]
	}

	// fmt.Println(temp)

	for x := 0; x < 10; x++ {
		for index, n := range nn {
			if n.V != 0 {
				pos := findIndex(temp, index)
				newPos := calcNewPos(pos, L, n.V)
				if pos < newPos {
					for i := pos; i < newPos; i++ {
						temp[i] = temp[i+1]
					}
					temp[newPos] = n
				} else if pos > newPos {
					for i := pos; i > newPos; i-- {
						temp[i] = temp[i-1]
					}
					temp[newPos] = n
				} else {
					fmt.Println("FUCK", newPos, pos, L, n.V)
				}
			}
		}
	}

	// fmt.Println(indexes)
	return temp
}

func findNumbers(nn []Number) []Number {
	L := len(nn)
	temp := make([]Number, L)
	for i := 0; i < L; i++ {
		temp[i] = nn[i]
	}

	// fmt.Println(temp)

	for index, n := range nn {
		if n.V != 0 {
			pos := findIndex(temp, index)
			newPos := calcNewPos(pos, L, n.V)

			// fmt.Printf("%5d   %4d -> %4d\n", n.V, pos, newPos)

			if pos < newPos {
				for i := pos; i < newPos; i++ {
					temp[i] = temp[i+1]
				}
				temp[newPos] = n
			} else if pos > newPos {
				for i := pos; i > newPos; i-- {
					temp[i] = temp[i-1]
				}
				temp[newPos] = n
			} else {
				fmt.Println("FUCK", newPos, pos, L, n.V)
			}
		}

		// fmt.Println(temp)
		// fmt.Println(indexes)
	}

	// fmt.Println(indexes)
	return temp
}

func calcNewPos(pos, L, v int) int {
	res := 0

	d := v % (L - 1)

	if pos+d == L-1 {
		fmt.Println("ACHTUNG")
	}

	if pos+d > 0 && pos+d <= L-1 {
		res = pos + d
	}

	if d < 0 && pos+d == 0 {
		res = L - 1 + pos + d
	}

	if d < 0 && pos+d < 0 {
		res = L - 1 + pos + d
	}

	if d > 0 && pos+d > L-1 {
		res = (pos + d) % (L - 1)
	}

	return res
}

func sign(v int) int {
	if v >= 0 {
		return 1
	}
	return -1
}

func findIndex(nn []Number, pos int) int {
	L := len(nn)
	for p := 0; p < L; p++ {
		if nn[p].Pos == pos {
			return p
		}
	}

	fmt.Println("PANIC findIndex")
	return -1
}

type Number struct {
	V   int
	Pos int
}

// 11123 input-2.txt
func loadInput() []Number {
	file, _ := os.Open("input-2.txt")
	scanner := bufio.NewScanner(file)

	res := []Number{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		v, _ := strconv.Atoi(line)

		res = append(res, Number{V: v, Pos: i})
		i++
	}

	return res
}
