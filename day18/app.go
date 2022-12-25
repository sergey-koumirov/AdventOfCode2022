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
	mmm := loadInput()
	r := calcSurface(mmm)
	fmt.Println("P1:", r)
}

func part2() {
	mmm := loadInput()
	r := calcOuterSurface(mmm)
	fmt.Println("P2:", r)
}

func calcOuterSurface(mmm []Point) int {
	xyz := makeXYZ(mmm)
	surfaces := makeSurfaces(mmm, xyz)
	topSurface := findTopZSurface(surfaces)

	// s := surfaces[topSurface]
	// fmt.Print(s.X, s.Y, s.Z, s.Side, " ->  ")
	// for _, ni := range s.Next {
	// 	n := surfaces[ni]
	// 	fmt.Print(n.X, n.Y, n.Z, n.Side, " ")
	// }
	// fmt.Println()

	surface := waveBySurface(surfaces, topSurface)

	return surface
}

func waveBySurface(surfaces []Surface, first int) int {
	current := []int{first}
	visited := map[int]int{first: 1}

	for {
		nextWave := []int{}
		for _, index := range current {
			s := surfaces[index]
			for _, nextIndex := range s.Next {
				if visited[nextIndex] != 1 {
					visited[nextIndex] = 1
					nextWave = append(nextWave, nextIndex)
				}
			}
		}

		if len(nextWave) == 0 {
			break
		}

		current = nextWave
	}

	return len(visited)
}

func findTopZSurface(surfaces []Surface) int {
	index := -1

	for i, s := range surfaces {
		if index == -1 || s.Z > surfaces[index].Z && s.Side.X == 0 && s.Side.Y == 0 && s.Side.Z == 1 {
			index = i
		}
	}

	return index
}

var NormalVectors = []Point{
	{X: 0, Y: 0, Z: 1},
	{X: 0, Y: 0, Z: -1},
	{X: 0, Y: 1, Z: 0},
	{X: 0, Y: -1, Z: 0},
	{X: 1, Y: 0, Z: 0},
	{X: -1, Y: 0, Z: 0},
}

func makeSurfaces(mmm []Point, xyz XYZ) []Surface {
	res := []Surface{}

	for _, p := range mmm {
		for _, nv := range NormalVectors {
			if hasSpace(p, xyz, nv) {
				res = append(res, Surface{X: p.X, Y: p.Y, Z: p.Z, Side: nv})
			}
		}
	}

	L := len(res)
	for i := 0; i < L; i++ {
		s := res[i]

		for j := 0; j < L; j++ {
			n := res[j]
			if isSideBySide(s, n, xyz) {
				res[i].Next = append(res[i].Next, j)
			}
		}
	}

	return res
}

func isSideBySide(s, n Surface, xyz XYZ) bool {
	// same cube
	if s.X == n.X && s.Y == n.Y && s.Z == n.Z {
		if s.Side.X != n.Side.X || s.Side.Y != n.Side.Y || s.Side.Z != n.Side.Z {
			notOpposite := s.Side.X+n.Side.X != 0 || s.Side.Y+n.Side.Y != 0 || s.Side.Z+n.Side.Z != 0
			return notOpposite && isEmptyXYZ(xyz, s.X+s.Side.X+n.Side.X, s.Y+s.Side.Y+n.Side.Y, s.Z+s.Side.Z+n.Side.Z)
		}
		return false
	}

	// have same side
	if abs(s.X-n.X) == 1 && s.Y-n.Y == 0 && s.Z-n.Z == 0 ||
		s.X-n.X == 0 && abs(s.Y-n.Y) == 1 && s.Z-n.Z == 0 ||
		s.X-n.X == 0 && s.Y-n.Y == 0 && abs(s.Z-n.Z) == 1 {
		return s.Side.X == n.Side.X && s.Side.Y == n.Side.Y && s.Side.Z == n.Side.Z
	}

	// have same edge
	if abs(s.X-n.X) == 1 && abs(s.Y-n.Y) == 1 && s.Z-n.Z == 0 ||
		s.X-n.X == 0 && abs(s.Y-n.Y) == 1 && abs(s.Z-n.Z) == 1 ||
		abs(s.X-n.X) == 1 && s.Y-n.Y == 0 && abs(s.Z-n.Z) == 1 {

		return s.X+s.Side.X == n.X+n.Side.X && s.Y+s.Side.Y == n.Y+n.Side.Y && s.Z+s.Side.Z == n.Z+n.Side.Z
	}

	return false
}

func hasSpace(p Point, xyz XYZ, side Point) bool {
	return isEmptyXYZ(xyz, p.X+side.X, p.Y+side.Y, p.Z+side.Z)
}

func isEmptyXYZ(xyz XYZ, x, y, z int) bool {
	xx, xEx := xyz[x]
	if !xEx {
		return true
	}

	yy, yEx := xx[y]
	if !yEx {
		return true
	}

	_, zEx := yy[z]

	return !zEx
}

func makeXYZ(mmm []Point) XYZ {
	res := XYZ{}
	for _, p := range mmm {

		_, xEx := res[p.X]
		if !xEx {
			res[p.X] = map[int]map[int]bool{}
		}

		_, yEx := res[p.X][p.Y]
		if !yEx {
			res[p.X][p.Y] = map[int]bool{}
		}

		res[p.X][p.Y][p.Z] = true
	}

	return res
}

func calcSurface(mmm []Point) int {
	surface := 0

	L := len(mmm)
	for i := 0; i < L; i++ {
		p1 := mmm[i]
		near := 0
		for j := 0; j < L; j++ {
			if i != j {
				p2 := mmm[j]
				if pointsConnected(p1, p2) {
					near++
				}
			}
		}
		surface += 6 - near
	}

	return surface
}

func pointsConnected(p1 Point, p2 Point) bool {
	byZ := p1.X == p2.X && p1.Y == p2.Y && abs(p1.Z-p2.Z) <= 1
	byY := p1.X == p2.X && p1.Z == p2.Z && abs(p1.Y-p2.Y) <= 1
	byX := p1.Z == p2.Z && p1.Y == p2.Y && abs(p1.X-p2.X) <= 1
	return byZ || byY || byX
}

func abs(v int) int {
	if v >= 0 {
		return v
	}
	return v * -1
}

type Point struct {
	X int
	Y int
	Z int
}

type Surface struct {
	X    int
	Y    int
	Z    int
	Side Point
	Next []int
}

type XYZ map[int]map[int]map[int]bool

func loadInput() []Point {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	res := []Point{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])

		res = append(res, Point{X: x, Y: y, Z: z})
	}

	return res
}
