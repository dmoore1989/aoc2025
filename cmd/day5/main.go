package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/dmoore1989/aoc2025/cmd/utils"
)

func main() {
	args := os.Args[1:]

	filePath := "sample.txt"
	if args[0] == "real" {
		filePath = "lib/day5.txt"

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
	var freshCount int
	segments := strings.Split(fileTxt, "\n\n")
	freshRange := [][2]int{}
	for fresh := range strings.SplitSeq(segments[0], "\n") {
		minMax := strings.Split(fresh, "-")
		min := utils.Atoi(minMax[0])
		max := utils.Atoi(minMax[1])
		freshRange = append(freshRange, [2]int{min, max})
	}

	for itemStr := range strings.SplitSeq(segments[1], "\n") {
		item := utils.Atoi(itemStr)
		for _, test := range freshRange {
			if item >= test[0] && item <= test[1] {
				freshCount += 1
				break
			}
		}
	}
	return strconv.Itoa(freshCount)
}

func part2(fileTxt string) string {
	segments := strings.Split(fileTxt, "\n\n")
	freshRange := [][2]int{}
	for fresh := range strings.SplitSeq(segments[0], "\n") {
		minMax := strings.Split(fresh, "-")
		min := utils.Atoi(minMax[0])
		max := utils.Atoi(minMax[1])
		freshRange = append(freshRange, [2]int{min, max})
	}
	// for _, x := range freshRange {
	// 	fmt.Println(x)
	// }

	slices.SortFunc(freshRange, func(a [2]int, b [2]int) int {
		if a[0] < b[0] {
			return -1
		} else if a[0] == b[0] {
			return 0
		} else {
			return 1
		}
	})

	i := 0

	for i < len(freshRange)-1 {
		if freshRange[i][1] >= freshRange[i+1][0] {
			if freshRange[i+1][1] > freshRange[i][1] {
				freshRange[i][1] = freshRange[i+1][1]
			}
			freshRange = slices.Delete(freshRange, i+1, i+2)
		} else {
			i += 1
		}
	}
	fmt.Println("========")
	var total int

	for _, n := range freshRange {
		total += n[1] - n[0] + 1
	}

	return strconv.Itoa(total)
}
