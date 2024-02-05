package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/wardzxzxc/aoc-2023/utils"
)

type Card struct {
	cardNum      int
	winningNums  []int
	selectedNums []int
}

func main() {
	part1("input.txt")
	part2("test.txt")
}

func part1(fileName string) {
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

		winningNums := strings.Split(strings.Split(strings.SplitN(line, ":", 2)[1], "|")[0], " ")
		selectedNums := strings.Split(strings.Split(strings.SplitN(line, ":", 2)[1], "|")[1], " ")
		sumPerCard := 0
		isFirst := true

		for _, num := range selectedNums {
			if num != "" {
				if slices.Contains(winningNums, num) && isFirst {
					isFirst = false
					sumPerCard += 1
				} else if slices.Contains(winningNums, num) && !isFirst {
					sumPerCard *= 2
				} else {
					continue
				}
			}
		}
		sum += sumPerCard
	}

	fmt.Println(sum)
}

func part2(fileName string) {
	allLines := utils.ReadAllLines(fileName)
	cards := parseInput(allLines)
	fmt.Println(cards)
}

func parseInput(allLines []string) []Card {
	cards := make([]Card, 0)

	for _, line := range allLines {
		cardNum, _ := strconv.Atoi(strings.TrimSpace(line[5:strings.Index(line, ":")]))
		winningNumsString := strings.Split(
			strings.Split(strings.SplitN(line, ":", 2)[1], "|")[0],
			" ",
		)
		selectedNumsString := strings.Split(
			strings.Split(strings.SplitN(line, ":", 2)[1], "|")[0],
			" ",
		)

		winningNums := make([]int, 0)
		selectedNums := make([]int, 0)

		for _, num := range winningNumsString {
			if string(num) != "" {
				winningNum, _ := strconv.Atoi(string(num))
				winningNums = append(winningNums, winningNum)
			}
		}

		for _, num := range selectedNumsString {
			if string(num) != "" {
				selectedNum, _ := strconv.Atoi(string(num))
				selectedNums = append(selectedNums, selectedNum)
			}
		}

		cards = append(
			cards,
			Card{cardNum: cardNum, winningNums: winningNums, selectedNums: selectedNums},
		)
	}

	return cards
}
