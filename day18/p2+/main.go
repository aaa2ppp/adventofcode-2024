package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	. "adventofcode-2024/utils"
)

func solve(n int, points [][2]int) int {
	desk := newDesk2(n, points)

	frontier := NewQueueSize[[2]int](n * 2)
	visited := MakeMatrix[byte](n, n)

	var count byte
	ans := sort.Search(len(points), func(i int) bool {
		count++
		return !searchPath(i, desk, frontier, visited, count)
	})

	return ans
}

type queue[T any] interface {
	Empty() bool
	Pop() T
	Push(v T)
	Clear()
}

func searchPath(ii int, desk [][]int, frontier queue[[2]int], visited [][]byte, count byte) bool {
	n := len(desk)

	start := [2]int{0, 0}
	finish := [2]int{n - 1, n - 1}

	frontier.Clear()

	frontier.Push(start)
	visited[start[0]][start[1]] = count

	for !frontier.Empty() {
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

			if desk[neigI][neigJ] <= ii {
				continue
			}

			if visited[neigI][neigJ] == count {
				continue
			}

			visited[neigI][neigJ] = count
			frontier.Push([2]int{neigI, neigJ})
		}
	}

	return visited[finish[0]][finish[1]] == count
}

func newDesk2(n int, points [][2]int) [][]int {

	matrix := MakeMatrix[int](n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = 100500 // +infinity
		}
	}

	for i, p := range points {
		x, y := p[0], p[1]
		matrix[y][x] = i
	}

	return matrix
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

func scanData(sc *bufio.Scanner) (int, [][2]int) {
	n, err := ScanInt(sc)
	if err != nil {
		panic(err)
	}

	m, err := ScanInt(sc)
	if err != nil {
		panic(err)
	}
	_ = m

	var points [][2]int

	for {
		if !sc.Scan() {
			if err := sc.Err(); err != nil {
				panic(err)
			}
			break
		}

		xy := UnsafeString(sc.Bytes())
		p := strings.Index(xy, ",")
		if p == -1 {
			panic("comma not found, want two num sperated by comma")
		}

		x, err := strconv.Atoi(xy[:p])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(xy[p+1:])
		if err != nil {
			panic(err)
		}

		points = append(points, [2]int{x, y})
	}

	return n, points
}

func run(in io.Reader, out io.Writer, solve func(n int, points [][2]int) int) {
	// br := bufio.NewReader(in)
	sc := bufio.NewScanner(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	n, points := scanData(sc)
	ans := solve(n, points)

	if debugEnable {
		log.Println("idx:", ans, "point:", points[ans])
		showDesk(newDesk(n, points[:ans+1]))
	}

	WriteInts(bw, points[ans][:], WriteOpts{Sep: ',', End: '\n'})
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout, solve)
}
