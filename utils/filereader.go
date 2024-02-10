package utils

import (
	"bufio"
	"os"
	"path"
	"runtime"
)

func ReadAllLines(name string) []string {
	_, callerName, _, ok := runtime.Caller(1)

	if !ok {
		panic("caller cannot be found, cannot build path")
	}

	inputFile, err := os.Open(path.Join(path.Dir(callerName), name))
	if err != nil {
		panic(err)
	}

	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadParagraphs(name string) [][]string {
	var paragraphs [][]string
	currPara := make([]string, 0)

	file := readFile(name)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			paragraphs = append(paragraphs, currPara)
			currPara = make([]string, 0)
		} else {
			currPara = append(currPara, line)
		}
	}

	if len(currPara) > 1 {
		paragraphs = append(paragraphs, currPara)
	}

	return paragraphs
}

func readFile(name string) *os.File {
	_, callerName, _, ok := runtime.Caller(2)

	if !ok {
		panic("caller cannot be found, cannot build path")
	}

	inputFile, err := os.Open(path.Join(path.Dir(callerName), name))
	if err != nil {
		panic(err)
	}
	return inputFile
}
