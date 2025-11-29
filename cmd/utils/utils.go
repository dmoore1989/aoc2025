package utils

import (
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

func Mod[T constraints.Integer](num T, mod T) T {
	if num >= 0 {
		return num % mod
	}

	return (mod + (num % mod)) % 4
}

func Abs[T constraints.Integer](num T) T {
	if num < 0 {
		return -num
	}
	return num
}
