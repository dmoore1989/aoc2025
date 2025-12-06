package utils

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

var Directions = map[rune][2]int{
	'<': {0, -1},
	'^': {-1, 0},
	'v': {1, 0},
	'>': {0, 1},
}

var Ordinals = map[rune][2]int{
	'W': {0, -1},
	'N': {-1, 0},
	'S': {1, 0},
	'E': {0, 1},
}

func ToSliceNum(s []string) []int {
	var a []int
	for _, n := range s {
		a = append(a, Atoi(n))
	}
	return a
}

func Mod[T constraints.Integer](num T, mod T) T {
	if num >= 0 {
		return num % mod
	}

	return (mod + (num % mod)) % mod
}

func Abs[T constraints.Integer](num T) T {
	if num < 0 {
		return -num
	}
	return num
}

func InsideSlice(index, length int) bool {
	return index >= 0 && index < length
}

// It's advent of code, I couldn't care less if my integer conversion has an error...
func Atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
