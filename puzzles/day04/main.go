package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	part1("input.txt")
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
