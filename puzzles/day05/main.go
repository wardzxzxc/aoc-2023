package main

import (
	"fmt"
	"strings"

	"github.com/wardzxzxc/aoc-2023/utils"
)

type ResourceMap struct {
	dest        int
	source      int
	rangeLength int
}

type Puzzle struct {
	seeds       []int
	seedToSoil  []ResourceMap
	soilToFert  []ResourceMap
	fertToWater []ResourceMap
	waterToList []ResourceMap
	lightToTemp []ResourceMap
	tempToHumid []ResourceMap
	humidToLoc  []ResourceMap
}

func main() {
	part1("test.txt")
}

func part1(fileName string) {
	fmt.Println(parseInput(fileName))
}

func parseInput(fileName string) Puzzle {
	paragraphs := utils.ReadParagraphs(fileName)

	seedsString := strings.Split(strings.Split(paragraphs[0][0], ":")[1], " ")

	seeds := make([]int, 0)
	for _, seed := range seedsString {
		if seed != "" {
			seeds = append(seeds, utils.ConvertStrToInt(seed))
		}
	}

	var allMappings [][]ResourceMap

	for i := 1; i < len(paragraphs); i++ {
		allMappings = append(allMappings, parseMapping(paragraphs[i]))
	}

	return Puzzle{
		seeds:       seeds,
		seedToSoil:  allMappings[0],
		soilToFert:  allMappings[1],
		fertToWater: allMappings[2],
		waterToList: allMappings[3],
		lightToTemp: allMappings[4],
		tempToHumid: allMappings[5],
		humidToLoc:  allMappings[6],
	}
}

func parseMapping(paragraph []string) []ResourceMap {
	mappings := make([]ResourceMap, 0)

	for i := 1; i < len(paragraph); i++ {
		mappingsSlice := strings.Split(paragraph[i], " ")
		mapping := ResourceMap{
			dest: utils.ConvertStrToInt(
				mappingsSlice[0],
			), source: utils.ConvertStrToInt(mappingsSlice[1]),
			rangeLength: utils.ConvertStrToInt(mappingsSlice[2]),
		}
		mappings = append(mappings, mapping)
	}
	return mappings
}
