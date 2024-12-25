package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"

	. "adventofcode-2024/utils"
)

const word = "XMAS"

var count = 0

func showWord(matrix [][]byte, visited [][]bool) {
	count++
	log.Printf("%d --------", count)

	n, m := len(matrix), len(matrix[0])
	row := make([]byte, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[i][j] {
				row[j] = matrix[i][j]
			} else {
				row[j] = '.'
			}
		}
		log.Printf("%s", row)
	}
}

func search(matrix [][]byte, visited [][]bool, i, j int, word string) int {
	visited[i][j] = true
	defer func() { visited[i][j] = false }()

	if matrix[i][j] != word[0] {
		return 0
	}

	word = word[1:]
	if len(word) == 0 {
		if debugEnable {
			showWord(matrix, visited)
		}
		return 1
	}

	n, m := len(matrix), len(matrix[0])
	count := 0

	for _, of := range [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
		neigI := i + of[0]
		neigJ := j + of[1]

		if !(0 <= neigI && neigI < n && 0 <= neigJ && neigJ < m) {
			continue
		}

		if visited[neigI][neigJ] {
			continue
		}

		count += search(matrix, visited, neigI, neigJ, word)
	}

	return count
}

func solve(matrix [][]byte) int {
	n, m := len(matrix), len(matrix[0])
	visited := MakeMatrix[bool](n, m)
	ans := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans += search(matrix, visited, i, j, word)
		}
	}

	return ans
}

func run(in io.Reader, out io.Writer) {
	buf, err := io.ReadAll(in)
	if err != nil {
		panic(err)
	}
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	if n := len(buf); buf[n-1] == '\n' {
		buf = buf[:n-1]
	}
	matrix := bytes.Split(buf, []byte("\n"))

	ans := solve(matrix)
	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
