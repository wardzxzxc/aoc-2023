package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/wardzxzxc/aoc-2023/utils"
)

func main() {
	part1("input.txt")
	part2("input.txt")
}

// Given index in string, find the entire number
func getNumber(idx int, line string) int {
	numString := string(line[idx])

	if !unicode.IsDigit(rune(line[idx])) {
		return 0
	}

	// Go backwards and find
	for i := idx - 1; i >= 0; i-- {
		char := string(line[i])
		if !unicode.IsDigit(rune(line[i])) {
			break
		} else {
			numString = char + numString
		}
	}

	// Go forward and find
	for j := idx + 1; j < len(line); j++ {
		char := string(line[j])
		if !unicode.IsDigit(rune(line[j])) {
			break
		} else {
			numString = numString + char
		}
	}

	i, err := strconv.Atoi(numString)
	if err != nil {
		panic(err)
	}

	return i
}

func getSurroundingNumbers(row int, col int, lines []string) []int {
	rowDirs := [3]int{1, 0, -1}
	colDirs := [3]int{1, 0, -1}
	nums := make(map[int]bool)
	finalNums := make([]int, 0)

	for _, rowDir := range rowDirs {
		for _, colDir := range colDirs {
			rowOther := rowDir + row
			colOther := colDir + col

			// Continue if we have exceeded the boundaries
			if rowOther < 0 || rowOther > len(lines) || colOther < 0 || colOther > len(lines[0]) {
				continue
			}

			// To not take the element in question
			if rowDir == 0 && colDir == 0 {
				continue
			}

			num := getNumber(colOther, lines[rowOther])

			// Check if number was found before as each digit in the number could be within the 9-boxes range, causing duplicates
			if num != 0 {
				if _, ok := nums[num]; !ok {
					nums[num] = true
				}
			}
		}
	}

	for key := range nums {
		finalNums = append(finalNums, key)
	}

	return finalNums
}

func part1(inputFile string) {
	lines := utils.ReadAllLines(inputFile)
	sum := 0

	for i, line := range lines {
		for j, rune := range line {
			if !unicode.IsDigit(rune) && string(rune) != "." {
				nums := getSurroundingNumbers(i, j, lines)
				for _, num := range nums {
					sum += num
				}
			}
		}
	}

	fmt.Println(sum)
}

func part2(inputFile string) {
	lines := utils.ReadAllLines(inputFile)
	sum := 0

	for i, line := range lines {
		for j, rune := range line {
			if string(rune) == "*" {
				nums := getSurroundingNumbers(i, j, lines)

				if len(nums) == 2 {
					sum += nums[0] * nums[1]
				}
			}
		}
	}

	fmt.Println(sum)
}
