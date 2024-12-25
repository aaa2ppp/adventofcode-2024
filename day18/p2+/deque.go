package main

import (
	"sort"

	. "adventofcode-2024/utils"
)

func dequeSolve(n int, points [][2]int) int {
	desk := newDesk2(n, points)

	frontier := NewDequeSize[[2]int](n * 2)
	visited := MakeMatrix[byte](n, n)

	var count byte
	ans := sort.Search(len(points), func(i int) bool {
		count++
		return !searchPath(i, desk, frontier, visited, count)
	})

	return ans
}
