package main

import (
	"fmt"
	"log"
	"maps"
	"math"
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
		filePath = "lib/day8.txt"

	}

	fileArr, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fileStr := strings.TrimSuffix(string(fileArr), "\n")

	var answer string
	if args[1] == "1" {
		answer = part1(fileStr, args[0])
	} else {
		answer = part2(fileStr)
	}
	fmt.Println(answer)

}

type distance struct {
	pointA string
	pointB string
	length float64
}

func calcDistance(point1, point2 string) float64 {
	p := strings.Split(point1, ",")
	a1 := float64(utils.Atoi(p[0]))
	a2 := float64(utils.Atoi(p[1]))
	a3 := float64(utils.Atoi(p[2]))

	p = strings.Split(point2, ",")
	b1 := float64(utils.Atoi(p[0]))
	b2 := float64(utils.Atoi(p[1]))
	b3 := float64(utils.Atoi(p[2]))

	return math.Sqrt(math.Pow((a1-b1), 2.0) + math.Pow((a2-b2), 2.0) + math.Pow((a3-b3), 2.0))
}

func part1(fileTxt, dataSet string) string {
	distances := []distance{}
	circuits := []map[string]bool{}
	points := strings.Split(fileTxt, "\n")
	for i, p1 := range points {
		pointMap := make(map[string]bool)
		pointMap[p1] = true
		circuits = append(circuits, pointMap)
		for _, p2 := range points[i+1:] {
			distances = append(distances, distance{
				pointA: p1,
				pointB: p2,
				length: calcDistance(p1, p2),
			})
		}
	}
	slices.SortFunc(distances, func(a distance, b distance) int {
		if a.length < b.length {
			return -1
		} else if a.length > b.length {
			return 1
		} else {
			return 0
		}
	})

	n := 10
	if dataSet == "real" {
		n = 1000
	}

	for i := 0; i < n; i++ {
		currentDist := distances[i]
		a := slices.IndexFunc(circuits, func(x map[string]bool) bool {
			_, ok := x[currentDist.pointA]
			return ok
		})
		b := slices.IndexFunc(circuits, func(x map[string]bool) bool {
			_, ok := x[currentDist.pointB]
			return ok
		})
		if a != b {
			maps.Copy(circuits[a], circuits[b])
			circuits = slices.Delete(circuits, b, b+1)
		}

	}

	slices.SortFunc(circuits, func(a map[string]bool, b map[string]bool) int {
		if len(a) > len(b) {
			return -1
		} else if len(a) < len(b) {
			return 1
		} else {
			return 0
		}
	})

	return strconv.Itoa(len(circuits[0]) * len(circuits[1]) * len(circuits[2]))
}

func part2(fileTxt string) string {
	distances := []distance{}
	circuits := []map[string]bool{}
	points := strings.Split(fileTxt, "\n")
	for i, p1 := range points {
		pointMap := make(map[string]bool)
		pointMap[p1] = true
		circuits = append(circuits, pointMap)
		for _, p2 := range points[i+1:] {
			distances = append(distances, distance{
				pointA: p1,
				pointB: p2,
				length: calcDistance(p1, p2),
			})
		}
	}
	slices.SortFunc(distances, func(a distance, b distance) int {
		if a.length < b.length {
			return -1
		} else if a.length > b.length {
			return 1
		} else {
			return 0
		}
	})
	i := 0
	for len(circuits) > 1 {
		currentDist := distances[i]
		a := slices.IndexFunc(circuits, func(x map[string]bool) bool {
			_, ok := x[currentDist.pointA]
			return ok
		})
		b := slices.IndexFunc(circuits, func(x map[string]bool) bool {
			_, ok := x[currentDist.pointB]
			return ok
		})

		if a != b {
			maps.Copy(circuits[a], circuits[b])
			circuits = slices.Delete(circuits, b, b+1)
		}

		i++
	}

	finConnection := distances[i-1]
	pointAX := utils.Atoi(strings.Split(finConnection.pointA, ",")[0])
	pointBX := utils.Atoi(strings.Split(finConnection.pointB, ",")[0])

	return strconv.Itoa(pointAX * pointBX)
}
