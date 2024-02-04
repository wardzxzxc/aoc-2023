package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1("input.txt")
	part2("input.txt")
}

func part2(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error reading input")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		allGamesCombined := strings.SplitN(line, ":", 2)[1]
		allGamesSeparated := strings.Split(allGamesCombined, ";")
		minCubes := make(map[string]int)
		for _, round := range allGamesSeparated {
			selections := strings.Split(round, ",")
			for _, selection := range selections {
				amountAndColor := strings.Split(strings.TrimSpace(selection), " ")
				amount, _ := strconv.Atoi(amountAndColor[0])
				color := amountAndColor[1]

				value, ok := minCubes[color]

				// if not found, put in map
				if !ok {
					minCubes[color] = amount
					continue
				}

				// if found but value smaller than amount, means new min
				if value < amount {
					minCubes[color] = amount
					continue
				}

			}
		}

		power := 1

		for _, v := range minCubes {
			power *= v
		}

		sum += power
	}

	fmt.Println(sum)
}

func part1(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error reading input")
		return
	}

	defer file.Close()

	availCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	scanner := bufio.NewScanner(file)

	currentGame := 1
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		allGamesCombined := strings.SplitN(line, ":", 2)[1]
		allGamesSeparated := strings.Split(allGamesCombined, ";")
		possible := true
		for _, round := range allGamesSeparated {
			selections := strings.Split(round, ",")
			for _, selection := range selections {
				amountAndColor := strings.Split(strings.TrimSpace(selection), " ")
				amount, _ := strconv.Atoi(amountAndColor[0])
				color := amountAndColor[1]
				if amount > availCubes[color] {
					// exit early if more than declared
					possible = false
					break
				}
			}
		}
		// if passes checks, add to sum
		if possible {
			sum += currentGame
		}
		currentGame += 1
	}

	fmt.Println(sum)
}
