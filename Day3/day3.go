package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputLines := readLineByLine("input.txt")

	sumErrors := 0
	for _, line := range inputLines {
		sumErrors += int(getPriority(findError(line)))
	}
	fmt.Println("[Part1] Sum of priorities:", sumErrors)

	sumRecurring := 0

	for i := 0; i < len(inputLines); i += 3 {
		recurring := findRecurringInGroup([]string{
			inputLines[i],
			inputLines[i+1],
			inputLines[i+2],
		})

		sumRecurring += int(getPriority(recurring))
	}

	fmt.Println("[Part2] Sum of all groups:", sumRecurring)

}

func getPriority(letter byte) byte {
	if letter > 96 {
		return letter - 96
	}
	return letter - 38
}

func getUniqueAscii(text string) []byte {
	uniqueLetters := map[byte]struct{}{}
	for _, letter := range []byte(text) {
		uniqueLetters[letter] = struct{}{}
	}
	bytes := []byte{}
	for key := range uniqueLetters {
		bytes = append(bytes, key)
	}
	return bytes
}

func findRecurringInGroup(group []string) byte {
	occurences := map[byte]int{}
	for _, compartment := range group {
		compartmentUnique := getUniqueAscii(compartment)
		for _, letter := range compartmentUnique {
			if occurence, ok := occurences[letter]; ok {
				occurences[letter] = occurence + 1
			} else {
				occurences[letter] = 1
			}
		}
	}

	return getThreeOccurenceLetter(occurences)
}

func getThreeOccurenceLetter(occurences map[byte]int) byte {
	for key, value := range occurences {
		if value == 3 {
			return key
		}
	}
	panic("could not find 3 times recurring letter in a group")
}

func findError(compartment string) byte {
	occurence := map[byte]struct{}{}
	asciiList := []byte(compartment)
	half := len(asciiList) / 2

	for i := 0; i < len(asciiList); i++ {
		codeUnit := asciiList[i]
		if i < half {
			occurence[codeUnit] = struct{}{}
		} else if _, ok := occurence[codeUnit]; ok {
			return codeUnit
		}
	}

	panic("could not find any recurring letter")
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
