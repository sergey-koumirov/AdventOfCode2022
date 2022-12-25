package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// part1()
	part2()
}

func part1() {
	records := loadInput()

	sum := 0
	for i, r := range records {
		order := compare(r.Left, r.Right)
		// fmt.Println(i+1, order)

		if order == -1 {
			sum += (i + 1)
		}
	}

	fmt.Println("P1:", sum)
}

func part2() {
	records := loadInput2()

	sort.Sort(records)

	mult := 1
	for i, r := range records {
		// fmt.Println(i+1, r)
		if r.Special {
			mult *= (i + 1)
		}
	}

	fmt.Println("P2:", mult)
}

func compare(left Values, right Values) int {

	if left.Kind == "N" && right.Kind == "N" {
		if left.V < right.V {
			return -1
		}
		if left.V > right.V {
			return 1
		}
		return 0
	}

	if left.Kind == "A" && right.Kind == "N" {
		return compare(left, Values{Kind: "A", Children: []Values{{Kind: "N", V: right.V}}})
	}

	if left.Kind == "N" && right.Kind == "A" {
		return compare(Values{Kind: "A", Children: []Values{{Kind: "N", V: left.V}}}, right)
	}

	if left.Kind == "A" && right.Kind == "A" {
		min := len(left.Children)
		if min > len(right.Children) {
			min = len(right.Children)
		}
		for i := 0; i < min; i++ {
			eqq := compare(left.Children[i], right.Children[i])
			if eqq != 0 {
				return eqq
			}
		}

		if len(left.Children) < len(right.Children) {
			return -1
		}
		if len(left.Children) > len(right.Children) {
			return 1
		}
		return 0
	}

	fmt.Print("PANIC !!! < ", left, "|", right, " >")
	return 0
}

type ByCustomRule []Values

func (a ByCustomRule) Len() int           { return len(a) }
func (a ByCustomRule) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCustomRule) Less(i, j int) bool { return compare(a[i], a[j]) == -1 }

type Values struct {
	Special  bool
	Kind     string
	V        int64
	Children []Values
}

type Pair struct {
	Left  Values
	Right Values
}

func loadInput() []Pair {
	res := []Pair{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line1 := scanner.Text()
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()

		temp := Pair{}

		temp.Left = parseInput(line1)
		temp.Right = parseInput(line2)

		res = append(res, temp)
	}

	return res
}

func loadInput2() ByCustomRule {
	res := ByCustomRule{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			temp := parseInput(line)
			res = append(res, temp)
		}
	}

	temp1 := parseInput("[[2]]")
	temp1.Special = true
	res = append(res, temp1)

	temp2 := parseInput("[[6]]")
	temp2.Special = true
	res = append(res, temp2)

	return res
}

// [[4,[4]],4,4,4]
func parseInput(s string) Values {
	// fmt.Println(s)
	res := Values{
		Children: []Values{},
	}

	if s == "[]" {
		res.Kind = "A"
	} else if s[0] == '[' {
		res.Kind = "A"

		deep := 0
		position := 0
		ss := s[1 : len(s)-1]
		for i, c := range ss {
			if c == '[' {
				deep++
			} else if c == ']' {
				deep--
			} else if deep == 0 && c == ',' {
				test := ss[position:i]
				res.Children = append(res.Children, parseInput(test))

				position = i + 1
			}
		}
		// fmt.Println(ss[position:])
		test := ss[position:]
		res.Children = append(res.Children, parseInput(test))

	} else {
		res.Kind = "N"
		v, _ := strconv.ParseInt(s, 10, 64)
		res.V = v
	}

	return res
}
