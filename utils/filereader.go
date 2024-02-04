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
