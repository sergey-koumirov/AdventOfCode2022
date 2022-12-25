package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	monkeys := loadInput()

	res := calcOp(monkeys, "root")

	fmt.Println("P1:", res)
}

func part2() {
	monkeys := loadInput()

	m := monkeys["root"]
	m.Sign = "="
	monkeys["root"] = m

	for {
		cnt := 0
		calcAndReplaceOp(monkeys, "root", &cnt)
		if cnt == 0 {
			break
		}
	}

	// for _, v := range monkeys {
	// 	if v.Sign != "->" {
	// 		v1 := v.Op1
	// 		if monkeys[v.Op1].Sign == "->" && v.Op1 != "humn" {
	// 			v1 = strconv.Itoa(monkeys[v.Op1].Value)
	// 		}
	// 		v2 := v.Op2
	// 		if monkeys[v.Op2].Sign == "->" && v.Op2 != "humn" {
	// 			v2 = strconv.Itoa(monkeys[v.Op2].Value)
	// 		}
	// 		fmt.Printf("%s   %s %s %s\n", v.Name, v1, v.Sign, v2)
	// 	}
	// }

	m = monkeys["root"]
	unknown := m.Op1
	known := m.Op2
	if monkeys[m.Op1].Sign == "->" {
		unknown = m.Op2
		known = m.Op1
	}

	// fmt.Println(monkeys[known].Value)

	k := calcK(monkeys, unknown, monkeys[known].Value)

	fmt.Println("P1:", k)
}

func printOp(mm map[string]Monkey, m Monkey) {
	if m.Sign != "->" {
		v1 := m.Op1
		if mm[m.Op1].Sign == "->" && m.Op1 != "humn" {
			v1 = strconv.Itoa(mm[m.Op1].Value)
		}
		v2 := m.Op2
		if mm[m.Op2].Sign == "->" && m.Op2 != "humn" {
			v2 = strconv.Itoa(mm[m.Op2].Value)
		}
		fmt.Printf("%s   %s %s %s\n", m.Name, v1, m.Sign, v2)
	}
}

func calcK(mm map[string]Monkey, current string, k int) int {
	m := mm[current]

	// printOp(mm, m)

	unknown := m.Op1
	known := m.Op2
	if mm[m.Op1].Sign == "->" && mm[m.Op1].Name != "humn" {
		unknown = m.Op2
		known = m.Op1
	}

	v := mm[known].Value

	var newK int

	if m.Sign == "+" {
		newK = k - v
	} else if m.Sign == "-" && known == m.Op2 {
		newK = k + v
	} else if m.Sign == "-" && known == m.Op1 {
		newK = v - k
	} else if m.Sign == "*" {
		newK = k / v
	} else if m.Sign == "/" {
		newK = k * v
	} else {
		fmt.Println("PANIC", m)
	}

	// fmt.Println(newK)

	if mm[unknown].Name == "humn" {
		return newK
	}
	return calcK(mm, unknown, newK)
}

func calcAndReplaceOp(mm map[string]Monkey, current string, cnt *int) {
	m := mm[current]
	if m.Sign == "->" {
		return
	}

	op1 := mm[m.Op1]
	op2 := mm[m.Op2]

	if op1.Name != "humn" && op2.Name != "humn" {
		if op1.Sign == "->" && op2.Sign == "->" {

			v := -1
			if m.Sign == "+" {
				v = op1.Value + op2.Value
			} else if m.Sign == "-" {
				v = op1.Value - op2.Value
			} else if m.Sign == "*" {
				v = op1.Value * op2.Value
			} else if m.Sign == "/" {
				v = op1.Value / op2.Value
			} else {
				fmt.Println("PANIC", m)
			}

			m.Sign = "->"
			m.Value = v
			mm[current] = m
			*cnt++
		} else {
			calcAndReplaceOp(mm, op1.Name, cnt)
			calcAndReplaceOp(mm, op2.Name, cnt)
		}
	} else if op1.Name != "humn" {
		calcAndReplaceOp(mm, op1.Name, cnt)
	} else if op2.Name != "humn" {
		calcAndReplaceOp(mm, op2.Name, cnt)
	}

}

func calcOp(mm map[string]Monkey, current string) int {

	m := mm[current]

	if m.Sign == "->" {
		return m.Value
	} else if m.Sign == "+" {
		return calcOp(mm, m.Op1) + calcOp(mm, m.Op2)
	} else if m.Sign == "-" {
		return calcOp(mm, m.Op1) - calcOp(mm, m.Op2)
	} else if m.Sign == "*" {
		return calcOp(mm, m.Op1) * calcOp(mm, m.Op2)
	} else if m.Sign == "/" {
		return calcOp(mm, m.Op1) / calcOp(mm, m.Op2)
	} else {
		fmt.Println("PANIC", m)
	}

	return 0
}

type Monkey struct {
	Name  string
	Sign  string
	Op1   string
	Op2   string
	Value int
}

func loadInput() map[string]Monkey {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	res := map[string]Monkey{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		temp := Monkey{Name: parts[0]}
		cuts := strings.Split(parts[1], " ")
		if len(cuts) == 1 {
			temp.Sign = "->"
			temp.Value, _ = strconv.Atoi(cuts[0])
		} else {
			temp.Op1 = cuts[0]
			temp.Sign = cuts[1]
			temp.Op2 = cuts[2]
		}
		res[temp.Name] = temp
	}

	return res
}
