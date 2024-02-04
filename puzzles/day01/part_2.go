package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func getAllKeys(m map[string]string) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	return keys

}

func extractNumbers(input string) []string {
	var wordToDigit = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	re := regexp.MustCompile(`\d|` + strings.Join(getAllKeys(wordToDigit), "|"))
	matches := re.FindAllString(input, -1)

	var digits []string

	for _, match := range matches {
		if _, err := strconv.Atoi(match); err == nil {
			digits = append(digits, match)
		} else if spelledOutNum, found := wordToDigit[match]; found {
			digits = append(digits, spelledOutNum)
		}
	}
	return digits
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("error")
		return
	}

	defer file.Close()

	var wordToDigit = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		var combined string

		re := regexp.MustCompile(`(?:` + strings.Join(getAllKeys(wordToDigit), "|") + `)`)

		for i, char := range line {
			if unicode.IsDigit(char) {
				combined = string(char)
				// found match and break
				break
			} else {
				word := line[:i+1]
				match := re.FindString(word)

				if match != "" {
					combined = wordToDigit[match]
					// found spelled word match and break
					break
				}
			}
		}

		// find reverse first occurrence
		for i := len(line) - 1; i >= 0; i-- {
			runeValue := rune(line[i])
			if unicode.IsDigit(runeValue) {
				combined += string(runeValue)
				// found match and break
				break
			} else {
				word := line[i:]
				match := re.FindString(word)

				if match != "" {
					combined += wordToDigit[match]
					// found spelled word match and break
					break
				}
			}
		}

		combinedInt, _ := strconv.Atoi(combined)
		sum += combinedInt
		fmt.Println(line)
		fmt.Println(combined)
	}

	fmt.Println(sum)
}
