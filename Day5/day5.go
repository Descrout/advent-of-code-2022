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
	inputLines := readLineByLine("input.txt")

	crates := newCrates(9)
	procs := []*Procedure{}

	for i, line := range inputLines {
		if i < 8 {
			crates.fillCrates(line)
			continue
		}

		if i > 9 {
			procs = append(procs, procFromStr(line))
		}
	}

	crates9001 := crates.copy()

	for _, proc := range procs {
		crates.applyProcedure(proc, false)
		crates9001.applyProcedure(proc, true)
	}

	fmt.Println("[Part1] Top of each stack CrateMover9000:", crates.topOfStacks())
	fmt.Println("[Part2] Top of each stack CrateMover9001:", crates9001.topOfStacks())
}

type Crates struct {
	crates [][]string
}

func newCrates(size int) *Crates {
	return &Crates{
		make([][]string, size),
	}
}

func (c *Crates) copy() *Crates {
	crates := newCrates(len(c.crates))
	copy(crates.crates, c.crates)
	return crates
}

func (c *Crates) fillCrates(line string) {
	j := 0
	for i := 1; i < len(line); i += 4 {
		char := string(line[i])
		if char != " " {
			c.crates[j] = append([]string{char}, c.crates[j]...)
		}
		j++
	}
}

func (c *Crates) applyProcedure(proc *Procedure, is9001 bool) {
	end := len(c.crates[proc.from])
	start := end - proc.move

	sublist := c.crates[proc.from][start:end]
	if !is9001 {
		sort.SliceStable(sublist, func(i, j int) bool {
			return i > j
		})
	}
	c.crates[proc.to] = append(c.crates[proc.to], sublist...)
	c.crates[proc.from] = c.crates[proc.from][:start]
}

func (c *Crates) topOfStacks() string {
	txt := ""
	for _, crate := range c.crates {
		txt += crate[len(crate)-1]
	}
	return txt
}

type Procedure struct {
	move int
	from int
	to   int
}

func procFromStr(txt string) *Procedure {
	splitted := strings.Split(txt, " ")

	move, _ := strconv.Atoi(splitted[1])
	from, _ := strconv.Atoi(splitted[3])
	to, _ := strconv.Atoi(splitted[5])

	return &Procedure{move, from - 1, to - 1}
}

func readLineByLine(filePath string) []string {
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileLines := []string{}
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines
}
