package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"

	. "adventofcode-2024/utils"
)

func solve(matrix [][]byte) int {
	n, m := len(matrix), len(matrix[0])

	ans := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited := matrix[i][j] & 128; visited == 0 {
				area, perimeter := countArea(matrix, i, j)
				ans += area * perimeter
			}
		}
	}

	return ans
}

func countArea(matrix [][]byte, i, j int) (area, perimeter int) {
	n, m := len(matrix), len(matrix[0])

	var frontier Queue[[2]int]
	frontier.Push([2]int{i, j})
	matrix[i][j] |= 128 // mark visited

	for !frontier.Empty() {
		p := frontier.Pop()
		i, j := p[0], p[1]
		v := matrix[i][j] & 127

		area++

		for _, of := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			neigI, neigJ := i+of[0], j+of[1]
			if !(0 <= neigI && neigI < n && 0 <= neigJ && neigJ < m) {
				perimeter++
				continue
			}

			if neigV := matrix[neigI][neigJ] & 127; neigV != v {
				perimeter++
				continue
			}

			if visited := matrix[neigI][neigJ] & 128; visited == 0 {
				matrix[neigI][neigJ] |= 128 // mark visited
				frontier.Push([2]int{neigI, neigJ})
			}
		}
	}

	return area, perimeter
}

func showMatrix(matrix [][]byte) {
	for _, row := range matrix {
		log.Printf("%s", row)
	}
}

func run(in io.Reader, out io.Writer) {
	buf, err := io.ReadAll(in)
	if err != nil {
		panic(err)
	}
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	matrix := bytes.Split(bytes.TrimSpace(buf), []byte("\n"))
	if debugEnable {
		showMatrix(matrix)
	}

	ans := solve(matrix)
	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
