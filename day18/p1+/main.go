package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	. "adventofcode-2024/utils"
)

func solve(desk [][]byte) int {
	n := len(desk)

	count := 0

	start := [2]int{0, 0}
	finish := [2]int{n - 1, n - 1}

	var frontier Queue[[2]int]

	visited := MakeMatrix[int](n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			visited[i][j] = -1
		}
	}

	frontier.Push(start)
	visited[start[0]][start[1]] = 0

	for !frontier.Empty() {
		count++
		p := frontier.Pop()
		if p == finish {
			break
		}
		i, j := p[0], p[1]

		for _, of := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			neigI := i + of[0]
			neigJ := j + of[1]

			if !(0 <= neigI && neigI < n && 0 <= neigJ && neigJ < n) {
				continue
			}

			if desk[neigI][neigJ] != '.' {
				continue
			}

			if visited[neigI][neigJ] != -1 {
				continue
			}

			visited[neigI][neigJ] = visited[i][j] + 1
			frontier.Push([2]int{neigI, neigJ})
		}
	}

	return visited[finish[0]][finish[1]]
}

func newDesk(n int, points [][2]int) [][]byte {

	matrix := MakeMatrix[byte](n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = '.'
		}
	}

	for _, p := range points {
		x, y := p[0], p[1]
		matrix[y][x] = '#'
	}

	return matrix
}

func showDesk(desk [][]byte) {
	for _, row := range desk {
		log.Printf("%s\n", row)
	}
}

func run(in io.Reader, out io.Writer) {
	// br := bufio.NewReader(in)
	sc := bufio.NewScanner(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	n, err := ScanInt(sc)
	if err != nil {
		panic(err)
	}

	m, err := ScanInt(sc)
	if err != nil {
		panic(err)
	}

	var points [][2]int

	for i := 0; i < m; i++ {
		if !sc.Scan() {
			if err := sc.Err(); err != nil {
				panic(err)
			}
			break
		}

		xy := strings.Split(UnsafeString(sc.Bytes()), ",")
		if len(xy) != 2 {
			panic("want two num")
		}

		x, err := strconv.Atoi(xy[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(xy[1])
		if err != nil {
			panic(err)
		}

		points = append(points, [2]int{x, y})
	}

	desk := newDesk(n, points)

	if debugEnable {
		showDesk(desk)
	}

	ans := solve(desk)

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
