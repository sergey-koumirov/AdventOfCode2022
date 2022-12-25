package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part2() {
	cmds := loadInput()

	rootInfo := NodeInfo{Name: "/", IsDir: true}
	infos := []*NodeInfo{&rootInfo}
	root := Node{Children: map[string]*Node{}, Info: &rootInfo}

	processCommands(cmds, &infos, &root)

	free := 70_000_000 - infos[0].Size
	need := 30_000_000 - free
	best := int64(70_000_000)

	for _, info := range infos {
		if info.IsDir {
			if info.Size >= need && info.Size < best {
				best = info.Size
			}
		}
	}

	fmt.Println("P2:", best)
}

func part1() {
	cmds := loadInput()

	rootInfo := NodeInfo{Name: "/", IsDir: true}
	infos := []*NodeInfo{&rootInfo}
	root := Node{Children: map[string]*Node{}, Info: &rootInfo}

	processCommands(cmds, &infos, &root)

	sum := int64(0)
	for _, info := range infos {
		if info.IsDir && info.Size <= 100000 {
			sum += info.Size
		}
	}

	fmt.Println("P1:", sum)
}

func processCommands(cmds []Raw, infos *[]*NodeInfo, root *Node) {
	current := root
	L := len(cmds)
	for i := 1; i < L; i++ {
		cmd := cmds[i]

		if cmd[0] == "$" && cmd[1] == "ls" {
			for {
				i++
				addNode(current, cmds[i], infos)
				if i == L-1 || cmds[i+1][0] == "$" {
					break
				}
			}
		} else if cmd[0] == "$" && cmd[1] == "cd" {
			if cmd[2] == ".." {
				current = current.Parent
			} else {
				temp, ex := current.Children[cmd[2]]
				current = temp
				if !ex {
					fmt.Println("PANIC EX !!!")
				}
			}
		} else {
			fmt.Println("PANIC !!!")
		}
	}
}

func addNode(current *Node, cmd Raw, infos *[]*NodeInfo) {
	tempInfo := NodeInfo{}

	if cmd[0] == "dir" {
		tempInfo.IsDir = true
		tempInfo.Name = cmd[1]
	} else {
		size, _ := strconv.ParseInt(cmd[0], 10, 64)
		tempInfo.IsDir = false
		tempInfo.Name = cmd[1]
		tempInfo.Size = size

		pointer := current

		for {
			pointer.Info.Size += size
			pointer = pointer.Parent
			if pointer == nil {
				break
			}
		}

	}

	tempNode := Node{
		Parent:   current,
		Children: map[string]*Node{},
		Info:     &tempInfo,
	}

	current.Children[tempInfo.Name] = &tempNode
	*infos = append(*infos, &tempInfo)
}

type Raw []string

type NodeInfo struct {
	IsDir bool
	Name  string
	Size  int64
}

type Node struct {
	Parent   *Node
	Children map[string]*Node
	Info     *NodeInfo
}

func loadInput() []Raw {
	commands := []Raw{}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		pp := strings.Split(line, " ")
		commands = append(commands, pp)
	}

	return commands
}
