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
		filePath = "lib/day12.txt"

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

func part1(fileTxt string) string {
	var count int
	blockMap := map[int]int{
		0: 7,
		1: 6,
		2: 5,
		3: 7,
		4: 7,
		5: 7,
	}

	rows := strings.Split(fileTxt, "\n")[30:]
	for _, row := range rows {
		area := utils.Atoi(row[0:2]) * utils.Atoi(row[3:5])
		items := strings.Split(row, " ")[1:]
		var shapeArea int
		for i, item := range items {
			shapeArea += utils.Atoi(item) * blockMap[i]
		}
		if shapeArea <= area {
			count += 1
		}
	}

	return strconv.Itoa(count)
}

func part2(fileTxt string) string {
	return "Part 2: " + fileTxt
}
