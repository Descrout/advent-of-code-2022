package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputLines := readLineByLine("input.txt")

	pairs := []*Pair{}
	for _, line := range inputLines {
		pairs = append(pairs, pairFromStr(line))
	}

	fullyContains := 0
	overlaps := 0
	for _, pair := range pairs {
		if pair.fullyContains() {
			fullyContains++
			overlaps++
		} else if pair.overlaps() {
			overlaps++
		}
	}
	fmt.Println("[Part1] Fully contained pairs:", fullyContains)
	fmt.Println("[Part2] Overlapped pairs:", overlaps)
}

type Pair struct {
	min1 int
	max1 int
	min2 int
	max2 int
}

func pairFromStr(txt string) *Pair {
	commaSplit := strings.Split(txt, ",")
	first := strings.Split(commaSplit[0], "-")
	second := strings.Split(commaSplit[1], "-")

	min1, _ := strconv.Atoi(first[0])
	max1, _ := strconv.Atoi(first[1])
	min2, _ := strconv.Atoi(second[0])
	max2, _ := strconv.Atoi(second[1])

	return &Pair{min1, max1, min2, max2}
}

func (p *Pair) fullyContains() bool {
	return (p.min1 >= p.min2 && p.max1 <= p.max2) || (p.min2 >= p.min1 && p.max2 <= p.max1)
}

func (p *Pair) overlaps() bool {
	return (p.min1 >= p.min2 && p.min1 <= p.max2) || (p.min2 >= p.min1 && p.min2 <= p.max1)
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
