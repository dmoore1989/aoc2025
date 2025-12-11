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
		filePath = "lib/day10.txt"

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

type testCase struct {
	state []rune
	count int
	next  []int
}

func part1(fileTxt string) string {
	var count int
	for data := range strings.SplitSeq(fileTxt, "\n") {
		items := strings.Split(data, " ")
		goal := items[0][1 : len(items[0])-1]

		var buttonSwitches [][]int
		for _, buttons := range items[1 : len(items)-1] {
			buttonSwitches = append(buttonSwitches, convertToButtonSwitch(buttons))
		}
		count += findLowestCount(goal, buttonSwitches)
		fmt.Println(count)
	}

	return strconv.Itoa(count)
}

func convertToButtonSwitch(buttons string) []int {
	var buttonSwitch []int
	for _, button := range buttons[1 : len(buttons)-1] {
		buttonSwitch = append(buttonSwitch, utils.Atoi(string(button)))
	}
	return buttonSwitch
}

func findLowestCount(goal string, buttonSwitches [][]int) int {
	queue := []testCase{}
	for _, buttonSwitch := range buttonSwitches {
		queue = append(queue, testCase{
			state: []rune(strings.Repeat(".", len(goal))),
			count: 1,
			next:  buttonSwitch,
		})
	}
	var i int
	for i <= len(queue) {
		currentCase := queue[i]
		currentState := make([]rune, len(currentCase.state))
		copy(currentState, currentCase.state)
		for _, button := range currentCase.next {
			if currentState[button] == '.' {
				currentState[button] = '#'
			} else {
				currentState[button] = '.'
			}
		}

		if string(currentState) == goal {
			return currentCase.count
		}

		for _, buttonSwitch := range buttonSwitches {
			queue = append(queue, testCase{
				state: currentState,
				count: currentCase.count + 1,
				next:  buttonSwitch,
			})
		}
		i++
	}

	return 0
}

func part2(fileTxt string) string {
	return "Part 2: " + fileTxt
}
