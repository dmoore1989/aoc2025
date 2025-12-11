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
		filePath = "lib/day11.txt"

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
	path := make(map[string][]string)
	for line := range strings.SplitSeq(fileTxt, "\n") {
		items := strings.Split(line, ": ")
		path[items[0]] = strings.Split(items[1], " ")
	}

	return strconv.Itoa(explorePaths("you", path))
}

func explorePaths(point string, path map[string][]string) int {
	if point == "out" {
		return 1
	}
	var count int
	for _, next := range path[point] {
		count += explorePaths(next, path)
	}
	return count
}

func explorePaths2(point string, path map[string][]string, visitDac, visitFft bool, memo map[string]map[bool]map[bool]int) int {
	if _, ok := memo[point][visitDac][visitFft]; ok {
		return memo[point][visitDac][visitFft]
	}

	if point == "out" {
		if visitDac && visitFft {
			return 1
		} else {
			return 0
		}
	}
	if point == "fft" {
		visitFft = true
	}
	if point == "dac" {
		visitDac = true
	}

	var count int
	for _, next := range path[point] {
		count += explorePaths2(next, path, visitDac, visitFft, memo)
	}

	if _, ok := memo[point]; !ok {
		memo[point] = make(map[bool]map[bool]int)
	}
	if _, ok := memo[point][visitDac]; !ok {
		memo[point][visitDac] = make(map[bool]int)
	}
	memo[point][visitDac][visitFft] = count
	return count
}

func part2(fileTxt string) string {
	path := make(map[string][]string)
	for line := range strings.SplitSeq(fileTxt, "\n") {
		items := strings.Split(line, ": ")
		path[items[0]] = strings.Split(items[1], " ")
	}

	memo := make(map[string]map[bool]map[bool]int)
	return strconv.Itoa(explorePaths2("svr", path, false, false, memo))
}
