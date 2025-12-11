package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dmoore1989/aoc2025/cmd/utils"
	"zappem.net/pub/math/polygon"
)

func main() {
	args := os.Args[1:]

	filePath := "sample.txt"
	if args[0] == "real" {
		filePath = "lib/day9.txt"

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
	var maxArea int
	points := strings.Split(fileTxt, "\n")
	for i, point1 := range points {
		p := strings.Split(point1, ",")
		x1 := utils.Atoi(p[0])
		y1 := utils.Atoi(p[1])
		for _, point2 := range points[i+1:] {
			p = strings.Split(point2, ",")
			x2 := utils.Atoi(p[0])
			y2 := utils.Atoi(p[1])

			area := (utils.Abs(x2-x1) + 1) * (utils.Abs(y2-y1) + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return strconv.Itoa(maxArea)
}

func part2(fileTxt string) string {
	var maxArea int
	points := strings.Split(fileTxt, "\n")
	lightShape := polygon.Shape{
		PS: []polygon.Point{},
	}

	for _, point1 := range points {
		p := strings.Split(point1, ",")
		x1 := float64(utils.Atoi(p[0]))
		y1 := float64(utils.Atoi(p[1]))
		lightShape.PS = append(lightShape.PS, polygon.Point{X: x1, Y: y1})
	}

	for i, point1 := range points {
		pOne := strings.Split(point1, ",")
		x1 := utils.Atoi(pOne[0])
		y1 := utils.Atoi(pOne[1])

		for _, point2 := range points[i+1:] {
			pTwo := strings.Split(point2, ",")
			x2 := utils.Atoi(pTwo[0])
			y2 := utils.Atoi(pTwo[1])

			area := (utils.Abs(x2-x1) + 1) * (utils.Abs(y2-y1) + 1)
			if area > maxArea && rectInShape(x1, y1, x2, y2, &lightShape) {
				maxArea = area
			}
		}
	}

	return strconv.Itoa(maxArea)
}

func rectInShape(x1, y1, x2, y2 int, shape *polygon.Shape) bool {
	var minX, maxX int
	if x1 > x2 {
		maxX = x1
		minX = x2
	} else {
		maxX = x2
		minX = x1
	}
	for i := minX; i <= maxX; i++ {
		point := polygon.Point{X: float64(i), Y: float64(y1)}
		fmt.Println("first", point, shape)
		if !point.Inside(shape) {
			return false
		}

		fmt.Println("next", point)
		point = polygon.Point{X: float64(i), Y: float64(y2)}
		if !point.Inside(shape) {
			return false
		}
	}

	var minY, maxY int
	if y1 > y2 {
		maxY = y1
		minY = y2
	} else {
		maxY = y2
		minY = y1
	}
	for i := minY; i <= maxY; i++ {
		point := polygon.Point{X: float64(x1), Y: float64(i)}
		if !point.Inside(shape) {
			return false
		}

		point = polygon.Point{X: float64(x2), Y: float64(i)}
		if !point.Inside(shape) {
			return false
		}
	}

	return true
}
