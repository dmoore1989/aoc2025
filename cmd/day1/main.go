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
		filePath = "lib/day1.txt"

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
	position := 50
	count := 0
	turns := strings.SplitSeq(fileTxt, "\n")

	for turn := range turns {
		amount, _ := strconv.Atoi(turn[1:])

		if turn[0] == 'L' {
			position -= amount
		} else {
			position += amount
		}
		position = utils.Mod(position, 100)

		if position == 0 {
			count += 1
		}
	}

	return strconv.Itoa(count)
}

func part2(fileTxt string) string {
	position := 50
	count := 0
	turns := strings.SplitSeq(fileTxt, "\n")

	for turn := range turns {
		amount, _ := strconv.Atoi(turn[1:])

		for i := 0; i < amount; i++ {
			if turn[0] == 'L' {
				position -= 1
			} else {
				position += 1
			}

			if position == 0 {
				count += 1
			}
			position = utils.Mod(position, 100)
		}
	}

	return strconv.Itoa(count)
}
