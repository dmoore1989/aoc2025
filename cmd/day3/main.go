package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]

	filePath := "sample.txt"
	if args[0] == "real" {
		filePath = "lib/day3.txt"

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
	var total, largest int
	row := strings.SplitSeq(fileTxt, "\n")
	for batteries := range row {
		for i, battery1 := range batteries {
			for _, battery2 := range batteries[i+1:] {
				num, _ := strconv.Atoi(string([]rune{battery1, battery2}))
				if num > largest {
					largest = num
				}
			}
		}
		total += largest
		largest = 0
	}
	return strconv.Itoa(total)
}

func part2(fileTxt string) string {

	var total int
	row := strings.SplitSeq(fileTxt, "\n")
	for batteries := range row {
		nextAmount, _ := strconv.Atoi(findLargest(batteries, 12))
		fmt.Println(nextAmount)
		total += nextAmount
	}
	return strconv.Itoa(total)
}

func findLargest(amount string, num int) string {
	if num == 1 {
		return amount
	}

	var largest int
	finalIdx := len(amount) - num
	for i, test := range amount[:finalIdx] {
		nextSequence := findLargest(amount[i:], num-1)
		testAmount, _ := strconv.Atoi(string(test) + nextSequence)
		if testAmount > largest {
			largest = testAmount
		}
	}
	return strconv.Itoa(largest)
}
