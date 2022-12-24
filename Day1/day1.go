package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func sumArray(arr []int) int {
	total := 0

	for _, num := range arr {
		total += num
	}

	return total
}

func main() {
	inputLines := readLineByLine("input.txt")

	latestElfCalories := []int{}
	maxElfCalori := 0

	for _, line := range inputLines {
		if len(line) == 0 {
			totalCalori := sumArray(latestElfCalories)
			if totalCalori > maxElfCalori {
				maxElfCalori = totalCalori
			}
			latestElfCalories = latestElfCalories[:0]
			continue
		}

		calori, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		latestElfCalories = append(latestElfCalories, calori)
	}

	fmt.Println("Most carried calories:", maxElfCalori)
}
