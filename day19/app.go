package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	bps := loadInput()
	sum := 0
	for i, bp := range bps {
		st := State{Minute: 0, OreBots: 1}
		max := 0
		calcGeodes(st, bp, &max, 24)
		sum += (i + 1) * max
		fmt.Println(i+1, max)
	}
	fmt.Println("P1:", sum)
}

func part2() {
	bps := loadInput()
	mult := 1
	for i, bp := range bps[0:3] {
		st := State{Minute: 0, OreBots: 1}
		max := 0
		calcGeodes(st, bp, &max, 32)
		mult *= max
		fmt.Println(i+1, max, (i+1)*max)
	}

	fmt.Println("P2:", mult)
}

func calcGeodes(st State, bp Blueprint, max *int, stop int) {
	possibleMax := st.Geode + (stop-st.Minute)*st.GeodeBots + (stop-1-st.Minute)*(stop-st.Minute)/2
	if *max >= possibleMax {
		return
	}

	nextGeode := st.Minute < stop-1 &&
		st.Ore >= bp.Geode_Ore &&
		st.Obs >= bp.Geode_Obs

	nextObs := st.Minute < stop-1 &&
		st.Ore >= bp.Obs_Ore &&
		st.Clay >= bp.Obs_Clay &&
		bp.Geode_Obs > st.ObsBots &&
		((stop-1-st.Minute)*bp.Geode_Obs > st.Obs+(stop-2-st.Minute)*st.ObsBots)

	nextClay := st.Minute < stop-1 &&
		st.Ore >= bp.Clay_Ore &&
		bp.Obs_Clay > st.ClayBots &&
		((stop-1-st.Minute)*bp.Obs_Clay > st.Clay+(stop-2-st.Minute)*st.ClayBots)

	nextOre := st.Minute < stop-1 &&
		st.Ore >= bp.Ore_Ore &&
		bp.Max_Ore > st.OreBots &&
		((stop-1-st.Minute)*bp.Max_Ore > st.Ore+(stop-2-st.Minute)*st.OreBots)

	st.Ore += st.OreBots
	st.Clay += st.ClayBots
	st.Obs += st.ObsBots
	st.Geode += st.GeodeBots

	st.Minute++

	if nextGeode {
		tempSt := st
		tempSt.Ore -= bp.Geode_Ore
		tempSt.Obs -= bp.Geode_Obs
		tempSt.GeodeBots++
		calcGeodes(tempSt, bp, max, stop)
	}

	if nextObs {
		tempSt := st
		tempSt.Ore -= bp.Obs_Ore
		tempSt.Clay -= bp.Obs_Clay
		tempSt.ObsBots++

		calcGeodes(tempSt, bp, max, stop)
	}

	if nextClay {
		tempSt := st
		tempSt.Ore -= bp.Clay_Ore
		tempSt.ClayBots++

		calcGeodes(tempSt, bp, max, stop)
	}

	if nextOre {
		tempSt := st
		tempSt.Ore -= bp.Ore_Ore
		tempSt.OreBots++

		calcGeodes(tempSt, bp, max, stop)
	}

	if st.Minute < stop {
		calcGeodes(st, bp, max, stop)
	} else {
		if st.Geode > *max {
			*max = st.Geode
		}
	}
}

func printState(st State) {
	fmt.Printf("     Ore: %d  Clay: %d  Obs: %d  G: %d   Bots: %d %d %d %d\n", st.Ore, st.Clay, st.Obs, st.Geode, st.OreBots, st.ClayBots, st.ObsBots, st.GeodeBots)
	fmt.Println()
}

type State struct {
	Minute    int
	OreBots   int
	Ore       int
	ClayBots  int
	Clay      int
	ObsBots   int
	Obs       int
	GeodeBots int
	Geode     int
}

type Blueprint struct {
	Ore_Ore   int
	Clay_Ore  int
	Obs_Ore   int
	Obs_Clay  int
	Geode_Ore int
	Geode_Obs int
	Max_Ore   int
}

func loadInput() []Blueprint {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	res := []Blueprint{}

	for scanner.Scan() {
		line := scanner.Text()

		re := regexp.MustCompile("[0-9]+")
		parts := re.FindAllString(line, -1)

		temp := Blueprint{}
		temp.Ore_Ore, _ = strconv.Atoi(parts[1])
		temp.Clay_Ore, _ = strconv.Atoi(parts[2])
		temp.Obs_Ore, _ = strconv.Atoi(parts[3])
		temp.Obs_Clay, _ = strconv.Atoi(parts[4])
		temp.Geode_Ore, _ = strconv.Atoi(parts[5])
		temp.Geode_Obs, _ = strconv.Atoi(parts[6])

		if temp.Ore_Ore >= temp.Clay_Ore && temp.Ore_Ore >= temp.Obs_Ore {
			temp.Max_Ore = temp.Ore_Ore
		} else if temp.Clay_Ore >= temp.Ore_Ore && temp.Clay_Ore >= temp.Obs_Ore {
			temp.Max_Ore = temp.Clay_Ore
		} else {
			temp.Max_Ore = temp.Obs_Ore
		}

		res = append(res, temp)
	}

	return res
}
