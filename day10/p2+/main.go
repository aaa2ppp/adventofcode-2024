package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"

	. "adventofcode-2024/utils"
)

func showVisited(matrix [][]int) {
	for _, row := range matrix {
		log.Printf("%2d", row)
	}
}

func solve(matrix [][]byte) int {
	n := len(matrix)
	m := len(matrix[0])

	var frontier Queue[[2]int]
	visited := MakeMatrix[int](n, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '0' {
				frontier.Push([2]int{i, j})
				visited[i][j] = 1
			}
		}
	}

	count := 0

	for !frontier.Empty() {
		p := frontier.Pop()
		i, j := p[0], p[1]
		v := matrix[i][j]

		if v == '9' {
			count += visited[i][j]
			continue
		}

		for _, of := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {

			neigI := i + of[0]
			neigJ := j + of[1]

			if !(0 <= neigI && neigI < n && 0 <= neigJ && neigJ < m) {
				continue
			}

			if matrix[neigI][neigJ] != v+1 {
				continue
			}

			if visited[neigI][neigJ] == 0 {
				frontier.Push([2]int{neigI, neigJ})
			}

			visited[neigI][neigJ] += visited[i][j]
		}
	}

	return count
}

func run(in io.Reader, out io.Writer) {
	buf, err := io.ReadAll(in)
	if err != nil {
		panic(err)
	}
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	matrix := bytes.Split(bytes.TrimSpace(buf), []byte("\n"))

	ans := solve(matrix)
	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
