package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var latestElfCalories = []int{}
var topThreeElfCalories = []int{0, 0, 0}

const topElfCaloriesLength = 3

func main() {
	inputLines := readLineByLine("input.txt")

	for _, line := range inputLines {
		if len(line) == 0 {
			totalCalori := sumArray(latestElfCalories)
			updateTopThree(totalCalori)
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

	fmt.Println("[Part1] Most carried calories:", topThreeElfCalories[0])
	topThreeSum := sumArray(topThreeElfCalories)
	fmt.Println("[Part2] Top three calories sum:", topThreeSum)
}

func updateTopThree(newCalories int) {
	for i := 0; i < topElfCaloriesLength; i++ {
		if newCalories > topThreeElfCalories[i] {
			oldCalori := topThreeElfCalories[i]
			topThreeElfCalories[i] = newCalories
			updateTopThree(oldCalori)
			break
		}
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

func sumArray(arr []int) int {
	total := 0

	for _, num := range arr {
		total += num
	}

	return total
}
