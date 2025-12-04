package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dmoore1989/aoc2025/cmd/utils"
)

func main() {
	args := os.Args[1:]

	filePath := "sample.txt"
	if args[0] == "real" {
		filePath = "lib/day4.txt"

	}

	fileArr, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fileStr := strings.TrimSuffix(string(fileArr), "\n")

	var answer string
	if args[1] == "1" {
		answer = part1(fileStr)
	} else {
		answer = part2(fileStr)
	}
	fmt.Println(answer)

}

var adjacentChecks = [][2]int{
	{0, -1},
	{-1, 0},
	{1, 0},
	{0, 1},
	{1, -1},
	{-1, 1},
	{1, 1},
	{-1, -1},
}

func part1(fileTxt string) string {
	var accCount int
	grid := strings.Split(fileTxt, "\n")
	for i, row := range grid {
		for j, space := range row {
			if space != '@' {
				continue
			}
			rollCount := 0
			for _, check := range adjacentChecks {
				newI := i + check[0]
				newJ := j + check[1]
				if utils.InsideSlice(newI, len(grid)) && utils.InsideSlice(newJ, len(row)) {
					if grid[newI][newJ] == '@' {
						rollCount += 1
					}
				}
			}
			if rollCount < 4 {
				accCount += 1
			}
		}
	}

	return strconv.Itoa(accCount)
}

func part2(fileTxt string) string {
	var accCount int
	grid := strings.Split(fileTxt, "\n")
	hasRolls := true
	for hasRolls {
		tpSet := make(map[[2]int]bool)
		hasRolls = false
		for i, row := range grid {
			for j, space := range row {
				if space != '@' {
					continue
				}
				rollCount := 0
				for _, check := range adjacentChecks {
					newI := i + check[0]
					newJ := j + check[1]
					if utils.InsideSlice(newI, len(grid)) && utils.InsideSlice(newJ, len(row)) {
						if grid[newI][newJ] == '@' {
							rollCount += 1
						}
					}
				}
				if rollCount < 4 {
					hasRolls = true
					accCount += 1
					tpSet[[2]int{i, j}] = true
				}
			}
		}

		for point := range tpSet {
			row := grid[point[0]]
			r := []rune(row)
			r[point[1]] = 'x'
			grid[point[0]] = string(r)
		}
	}

	return strconv.Itoa(accCount)
}
