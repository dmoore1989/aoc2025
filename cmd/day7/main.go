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
		filePath = "lib/day7.txt"

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
	rows := strings.Split(fileTxt, "\n")
	var point int
	// Find S
	for i, char := range rows[0] {
		if char == 'S' {
			point = i
			break
		}
	}

	splitCount := findSplits(rows, point)
	return strconv.Itoa(splitCount)
}

func findSplits(chart []string, point int) int {
	var count int
	for i, row := range chart {
		if row[point] == '|' {
			break
		}
		if row[point] == '^' {
			count += 1
			count += findSplits(chart[i:], point-1)
			count += findSplits(chart[i:], point+1)
			break
		}
		r := []rune(chart[i])
		r[point] = '|'
		chart[i] = string(r)
	}

	return count
}

func part2(fileTxt string) string {
	rows := strings.Split(fileTxt, "\n")
	var point int
	memo := make(map[[2]int]int)
	// Find S
	for i, char := range rows[0] {
		if char == 'S' {
			point = i
			break
		}
	}

	splitCount := findPossibilities(rows, 0, point, memo) + 1
	return strconv.Itoa(splitCount)
}

func findPossibilities(chart []string, rowStart, point int, memo map[[2]int]int) int {
	if value, ok := memo[[2]int{rowStart, point}]; ok {
		return value
	}
	var count int
	for i, row := range chart[rowStart:] {
		if row[point] == '^' {
			count += 1
			count += findPossibilities(chart, rowStart+i, point-1, memo)
			count += findPossibilities(chart, rowStart+i, point+1, memo)
			break
		}
	}

	memo[[2]int{rowStart, point}] = count
	return count
}
