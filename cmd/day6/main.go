package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/dmoore1989/aoc2025/cmd/utils"
)

func main() {
	args := os.Args[1:]

	filePath := "sample.txt"
	if args[0] == "real" {
		filePath = "lib/day6.txt"

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
	var total int
	r := regexp.MustCompile(`\s+`)
	rows := strings.Split(fileTxt, "\n")

	var operations []string
	var operands [][]int
	for i, line := range rows {
		fmt.Println(line)
		line = strings.TrimSpace(line)
		line = r.ReplaceAllString(line, ` `)
		lineArr := strings.Split(line, " ")
		if i == len(rows)-1 {
			operations = lineArr
		} else {
			lineInt := utils.ToSliceNum(lineArr)
			operands = append(operands, lineInt)
		}
	}

	for i, operation := range operations {
		amount := 0
		if operation == "*" {
			amount = 1
		}
		for _, operand := range operands {
			if operation == "+" {
				amount += operand[i]
			} else {
				amount *= operand[i]
			}
		}
		total += amount
	}

	return strconv.Itoa(total)
}

func part2(fileTxt string) string {
	var total int
	var currentOp rune
	numbers := [][]int{{}}
	var ops []rune
	rows := strings.Split(fileTxt, "\n")
	for i := 0; i < len(rows[0]); i++ {
		numberStr := ""
		if rows[len(rows)-1][i] != ' ' {
			currentOp = rune(rows[len(rows)-1][i])
		}
		r := regexp.MustCompile(`\s+`)
		for _, char := range rows[:len(rows)-1] {
			numberStr += string(char[i])
		}
		numberStr = r.ReplaceAllString(numberStr, "")
		if numberStr == "" {
			numbers = append(numbers, []int{})
			ops = append(ops, currentOp)
		} else {
			numbers[len(numbers)-1] = append(numbers[len(numbers)-1], utils.Atoi(numberStr))
		}
	}
	ops = append(ops, currentOp)

	for i, nums := range numbers {
		operation := ops[i]
		amount := 0
		if operation == '*' {
			amount = 1
		}
		for _, num := range nums {
			if operation == '+' {
				amount += num
			} else {
				amount *= num
			}
		}
		total += amount
	}

	return strconv.Itoa(total)
}
