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
		filePath = "lib/day2.txt"

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
	var invalidSum int
	ranges := strings.SplitSeq(fileTxt, ",")
	for ids := range ranges {
		idsArr := strings.Split(ids, "-")
		start, _ := strconv.Atoi(idsArr[0])
		end, _ := strconv.Atoi(idsArr[1])
		for i := start; i <= end; i++ {
			if repeatingDigitsOne(i) {
				invalidSum += i
			}
		}
	}

	return strconv.Itoa(invalidSum)
}

func part2(fileTxt string) string {
	var invalidSum int
	ranges := strings.SplitSeq(fileTxt, ",")
	for ids := range ranges {
		idsArr := strings.Split(ids, "-")
		start, _ := strconv.Atoi(idsArr[0])
		end, _ := strconv.Atoi(idsArr[1])
		for i := start; i <= end; i++ {
			if repeatingDigitsTwo(i) {
				invalidSum += i
			}
		}
	}

	return strconv.Itoa(invalidSum)
}

func repeatingDigitsOne(num int) bool {
	intStr := strconv.Itoa(num)
	midPoint := len(intStr) / 2
	numOne := intStr[0:midPoint]
	numTwo := intStr[midPoint:]

	return numOne == numTwo
}

func repeatingDigitsTwo(num int) bool {
	intStr := strconv.Itoa(num)
	midPoint := len(intStr) / 2
	for intLength := 1; intLength <= midPoint; intLength++ {
		repeats := true
		i := 0
		for (i + 2*intLength) <= len(intStr) {
			numOne := intStr[i : i+intLength]
			numTwo := intStr[i+intLength : (i + 2*intLength)]
			if numOne != numTwo {
				repeats = false
				break
			}
			i += intLength
		}
		if repeats && i+intLength == len(intStr) {
			return true
		}
	}

	return false
}
