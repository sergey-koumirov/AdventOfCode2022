package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part1() {
	state := loadInput()
	state.Map = compact(state.Map, state.Rates)

	fmt.Println("AA --> ", state.Map["AA"])
	fmt.Println()

	enabled := map[string]int{}
	for k := range state.Map {
		enabled[k] = 0
	}
	current := "AA"
	minutes := 0
	deepScan(&state, enabled, current, minutes, 0)

	fmt.Println("P1:", state.Best)
}

func part2() {
	state := loadInput()
	state.Map = compact(state.Map, state.Rates)

	state.Sorted = make(RateVals, len(state.Map))

	index := 0
	for k := range state.Map {
		state.Sorted[index].Valve = k
		state.Sorted[index].Rate = state.Rates[k]
		index++
	}

	sort.Sort(state.Sorted)

	fmt.Println("AA --> ", state.Map["AA"])
	fmt.Println()

	enabled := map[string]int{"AA": 1}
	for k := range state.Map {
		enabled[k] = 0
	}
	curHum := "AA"
	curEl := "AA"
	deepScan2(&state, enabled, curHum, curEl, 0, 0, 0)

	fmt.Println("P2:", state.Best)
}

func tooLate(st *State, enabled map[string]int, minutes int, rate int64) bool {
	rest := int64(0)
	for _, k := range st.Sorted {
		if enabled[k.Valve] == 0 {
			kRate := st.Rates[k.Valve]
			rest += kRate * int64(26-minutes-1)
		}
	}

	return rate+rest <= st.Best
}

func deepScan2(st *State, enabled map[string]int, curHum string, curEl string, minutesHum int, minutesEl int, rate int64) {
	if minutesHum <= minutesEl {
		if minutesHum <= 24 {
			if tooLate(st, enabled, minutesHum, rate) {
				return
			}

			for _, next := range st.Map[curHum] {
				if enabled[curHum] == 0 {
					enabled[curHum] = 1
					newRate := rate + int64(26-minutesHum-1)*st.Rates[curHum]
					if st.Best < newRate {
						st.Best = newRate
						fmt.Println(newRate)
					}
					deepScan2(st, enabled, next.Valve, curEl, minutesHum+1+next.Cost, minutesEl, newRate)
					enabled[curHum] = 0
				}

				deepScan2(st, enabled, next.Valve, curEl, minutesHum+next.Cost, minutesEl, rate)
			}
		}
	} else {
		if minutesEl <= 24 {
			if tooLate(st, enabled, minutesEl, rate) {
				return
			}
			for _, next := range st.Map[curEl] {
				if enabled[curEl] == 0 {
					enabled[curEl] = 1
					newRate := rate + int64(26-minutesEl-1)*st.Rates[curEl]
					if st.Best < newRate {
						st.Best = newRate
						fmt.Println(newRate)
					}
					deepScan2(st, enabled, curHum, next.Valve, minutesHum, minutesEl+1+next.Cost, newRate)
					enabled[curEl] = 0
				}

				deepScan2(st, enabled, curHum, next.Valve, minutesHum, minutesEl+next.Cost, rate)
			}
		}
	}

}

func compact(m map[string]Destinations, r map[string]int64) map[string]Destinations {
	res := map[string]Destinations{}

	allDist := map[string]map[string]int{}

	for k := range m {
		dist := map[string]int{k: 0}
		calcDistances(k, m, dist)
		allDist[k] = dist
		// fmt.Println(k, dist)
	}

	significant := map[string]int{"AA": 1}
	for k, v := range r {
		if v > 0 {
			significant[k] = 1
		}
	}

	for fromKey := range significant {
		res[fromKey] = Destinations{}
		for toKey := range significant {
			if fromKey != toKey {
				res[fromKey] = append(res[fromKey], Destination{Valve: toKey, Cost: allDist[fromKey][toKey]})
			}
		}
		sort.Sort(res[fromKey])
	}

	return res
}

func calcDistances(current string, m map[string]Destinations, res map[string]int) {
	currentCost := res[current]

	for _, d := range m[current] {
		cost, ex := res[d.Valve]
		if !ex || cost > currentCost+1 {
			res[d.Valve] = currentCost + 1
			calcDistances(d.Valve, m, res)
		}
	}
}

func deepScan(st *State, enabled map[string]int, current string, minutes int, rate int64) {

	// fmt.Println(enabled, current, minutes, rate, path)

	if minutes <= 28 {

		rest := int64(0)
		for k := range st.Map {
			kRate := st.Rates[k]
			if enabled[k] == 0 && kRate > 0 {
				rest += kRate * int64(30-minutes-1)
			}
		}

		if rate+rest <= st.Best {
			return
		}

		for _, next := range st.Map[current] {
			curRate := st.Rates[current]
			if enabled[current] == 0 && curRate > 0 {
				enabled[current] = 1
				newRate := rate + int64(30-minutes-1)*curRate
				if st.Best < newRate {
					st.Best = newRate
					fmt.Println(newRate)
				}
				deepScan(st, enabled, next.Valve, minutes+1+next.Cost, newRate)
				enabled[current] = 0
			}

			deepScan(st, enabled, next.Valve, minutes+next.Cost, rate)
		}
	}
}

type Destinations []Destination

type Destination struct {
	Valve string
	Cost  int
}

func (a Destinations) Len() int           { return len(a) }
func (a Destinations) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Destinations) Less(i, j int) bool { return a[i].Cost < a[j].Cost }

type RateVals []RateVal

type RateVal struct {
	Valve string
	Rate  int64
}

func (a RateVals) Len() int           { return len(a) }
func (a RateVals) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a RateVals) Less(i, j int) bool { return a[i].Rate > a[j].Rate }

type State struct {
	Map    map[string]Destinations
	Sorted RateVals
	Rates  map[string]int64
	Best   int64
}

func loadInput() State {
	res := State{
		Map:   map[string]Destinations{},
		Rates: map[string]int64{},
	}

	file, _ := os.Open("input-1.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			// Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
			parts0 := strings.Split(line, " has flow rate=")
			parts1 := strings.Split(parts0[0], " ")
			name := parts1[1]

			parts2 := strings.Split(parts0[1], "; tunnels lead to valves ")
			parts22 := strings.Split(parts0[1], "; tunnel leads to valve ")
			if len(parts22) > len(parts2) {
				parts2 = parts22
			}

			rate, _ := strconv.ParseInt(parts2[0], 10, 64)
			tunnels := strings.Split(parts2[1], ", ")

			temp := []Destination{}

			for _, tunnel := range tunnels {
				temp = append(temp, Destination{Valve: tunnel, Cost: 1})
			}

			res.Map[name] = temp
			res.Rates[name] = rate
		}
	}

	return res
}
