package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// vv := []int64{11, 5, 2, 4, 11, 5, 2, 4, 11, 5, 2, 4}

	// x := int64(5)
	// y := x
	// d := int64(7)

	// for _, v := range vv {
	// 	y = y % d

	// 	fmt.Println(x, x%d, "?", y, y%d)

	// 	x = x * v
	// 	y = y * v

	// 	fmt.Println(x, x%d, "?", y, y%d)
	// 	fmt.Println()
	// }

	// part1()
	part2()
}

func part1() {
	monkeys := loadInput()

	for i := 1; i <= 20; i++ {
		playRound(&monkeys)
	}

	mb := find2Max(&monkeys)

	// fmt.Printf("%+v\n", monkeys)
	fmt.Println("P1:", mb)
}

func part2() {
	monkeys2 := loadInput()
	// fmt.Println(monkeys2)

	for i := 1; i <= 10000; i++ {
		playRound2(&monkeys2)

		if i == 1 || i == 20 || i == 1000 || i == 2000 || i == 3000 || i == 4000 || i == 5000 || i == 6000 || i == 7000 || i == 8000 || i == 9000 || i == 10000 {
			fmt.Println("R", i)
			for mi, m := range monkeys2 {
				fmt.Println("    ", mi, m.Bussiness)
			}
			fmt.Println()
		}

	}

	mb := find2Max(&monkeys2)

	fmt.Println("P2:", mb)
}

func printMonkeys(mm []Monkey) {
	for _, m := range mm {
		fmt.Println(m.Test, m.Items)
	}
	fmt.Println()
}

func find2Max(monkeys *[]Monkey) int {
	max1 := 0
	max2 := 0

	for _, m := range *monkeys {
		if max1 < m.Bussiness {
			max2 = max1
			max1 = m.Bussiness
		} else if max2 < m.Bussiness {
			max2 = m.Bussiness
		}
	}

	return max1 * max2
}

func playRound2(monkeys *[]Monkey) {

	for mi, m := range *monkeys {
		for _, items := range m.ItemsMod {
			if m.OpType == "square" {
				for i := 0; i < len(items); i++ {
					items[i] = (items[i] % (*monkeys)[i].Test) * items[i]
				}
			} else if m.OpType == "mult" {
				for i := 0; i < len(items); i++ {
					items[i] = (items[i] % (*monkeys)[i].Test) * m.OpValue
				}
			} else if m.OpType == "plus" {
				for i := 0; i < len(items); i++ {
					items[i] = items[i] + m.OpValue
				}
			}

			// fmt.Print(item, item%m.Test == 0, " | ")

			if items[mi]%m.Test == 0 {
				(*monkeys)[m.TestTrue].ItemsMod = append((*monkeys)[m.TestTrue].ItemsMod, items)
			} else {
				(*monkeys)[m.TestFalse].ItemsMod = append((*monkeys)[m.TestFalse].ItemsMod, items)
			}

		}
		(*monkeys)[mi].Bussiness += len((*monkeys)[mi].ItemsMod)
		(*monkeys)[mi].ItemsMod = [][]int64{}
	}
}

func playRound(monkeys *[]Monkey) {

	for mi, m := range *monkeys {
		for _, item := range m.Items {
			if m.OpType == "square" {
				item = item * item
			} else if m.OpType == "mult" {
				item = item * m.OpValue
			} else if m.OpType == "plus" {
				item = item + m.OpValue
			}
			item = item / 3
			if item%m.Test == 0 {
				(*monkeys)[m.TestTrue].Items = append((*monkeys)[m.TestTrue].Items, item)
			} else {
				(*monkeys)[m.TestFalse].Items = append((*monkeys)[m.TestFalse].Items, item)
			}
		}
		(*monkeys)[mi].Bussiness += len((*monkeys)[mi].Items)
		(*monkeys)[mi].Items = []int64{}
	}
}

type Monkey struct {
	Test      int64
	Items     []int64
	ItemsMod  [][]int64
	OpType    string //plus, mult, square
	OpValue   int64
	TestTrue  int64
	TestFalse int64
	Bussiness int
}

func loadInput() []Monkey {
	res := []Monkey{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		scanner.Scan()
		line2 := scanner.Text()
		scanner.Scan()
		line3 := scanner.Text()
		scanner.Scan()
		line4 := scanner.Text()
		scanner.Scan()
		line5 := scanner.Text()
		scanner.Scan()
		line6 := scanner.Text()
		scanner.Scan()

		temp := Monkey{}

		parts20 := strings.Split(line2, ": ")
		parts21 := strings.Split(parts20[1], ", ")
		for _, p := range parts21 {
			v, _ := strconv.ParseInt(p, 10, 64)
			temp.Items = append(temp.Items, v)
		}

		parts30 := strings.Split(line3, " = ")
		parts31 := strings.Split(parts30[1], " ")
		if parts31[0] == "old" && parts31[2] == "old" {
			temp.OpType = "square"
		} else {
			v, _ := strconv.ParseInt(parts31[2], 10, 64)
			temp.OpValue = v
			if parts31[1] == "+" {
				temp.OpType = "plus"
			} else if parts31[1] == "*" {
				temp.OpType = "mult"
			}
		}

		parts40 := strings.Split(line4, " by ")
		v4, _ := strconv.ParseInt(parts40[1], 10, 64)
		temp.Test = v4

		parts50 := strings.Split(line5, " monkey ")
		v5, _ := strconv.ParseInt(parts50[1], 10, 64)
		temp.TestTrue = v5

		parts60 := strings.Split(line6, " monkey ")
		v6, _ := strconv.ParseInt(parts60[1], 10, 64)
		temp.TestFalse = v6

		res = append(res, temp)
	}

	for i := range res {
		res[i].ItemsMod = make([][]int64, len(res[i].Items))
		for j := 0; j < len(res[i].ItemsMod); j++ {
			res[i].ItemsMod[j] = make([]int64, len(res))
			for k := 0; k < len(res); k++ {
				res[i].ItemsMod[j][k] = res[i].Items[j]
			}
		}
	}

	return res
}
