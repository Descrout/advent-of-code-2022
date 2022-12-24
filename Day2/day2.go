package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scores = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

// Some languages do negative number modulus different
// So instead of % wrap around
// I put "C" as first and "A" as last
var symbols = []string{"C", "A", "B", "C", "A"}

func main() {
	inputLines := readLineByLine("input.txt")

	scorePart1 := 0
	scorePart2 := 0
	for _, line := range inputLines {
		splitted := strings.Split(line, " ")
		other := splitted[0]
		mine := splitted[1]
		scorePart1 += getMyScore(other, mine)
		scorePart2 += getMyScorePart2(other, mine)
	}

	fmt.Println("[Part1] Your score is:", scorePart1)
	fmt.Println("[Part2] Your score is:", scorePart2)
}

func getMyScore(other, mine string) int {
	myScore := scores[mine]
	otherScore := scores[other]

	diff := myScore - otherScore
	if diff == 0 {
		myScore += 3
	} else if diff == 1 || diff == -2 {
		myScore += 6
	}

	return myScore
}

func getMyScorePart2(other, mine string) int {
	symbolIdx := scores[other]

	switch mine {
	case "X":
		return getMyScore(other, symbols[symbolIdx-1])
	case "Z":
		return getMyScore(other, symbols[symbolIdx+1])
	default:
		return getMyScore(other, other)
	}
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
