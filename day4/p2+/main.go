package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"slices"

	. "adventofcode-2024/utils"
)

const word = "MAS"

var count = 0

func showWord(matrix [][]byte, i, j int) {
	n, m := len(matrix), len(matrix[0])

	count++
	log.Printf("%d -------", count)

	row := make([]byte, m)
	for ii := 0; ii < n; ii++ {
		for jj := 0; jj < m; jj++ {
			of := [2]int{i - ii, j - jj}
			if slices.Contains([][2]int{{-1, -1}, {-1, 1}, {0, 0}, {1, -1}, {1, 1}}, of) {
				row[jj] = matrix[ii][jj]
			} else {
				row[jj] = '.'
			}

		}
		log.Printf("%s", row)
	}
}

func check(matrix [][]byte, i, j int) bool {
	n, m := len(matrix), len(matrix[0])

	if !(1 <= i && i < n-1 && 1 <= j && j < m-1) {
		return false
	}

	if matrix[i][j] != 'A' {
		return false
	}

	if matrix[i-1][j-1] != 'M' && matrix[i-1][j-1] != 'S' {
		return false
	}

	if matrix[i-1][j-1] == 'M' && matrix[i+1][j+1] != 'S' {
		return false
	}

	if matrix[i-1][j-1] == 'S' && matrix[i+1][j+1] != 'M' {
		return false
	}

	if matrix[i+1][j-1] != 'M' && matrix[i+1][j-1] != 'S' {
		return false
	}

	if matrix[i+1][j-1] == 'M' && matrix[i-1][j+1] != 'S' {
		return false
	}

	if matrix[i+1][j-1] == 'S' && matrix[i-1][j+1] != 'M' {
		return false
	}

	return true
}

func solve(matrix [][]byte) int {
	n, m := len(matrix), len(matrix[0])
	ans := 0

	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if check(matrix, i, j) {
				if debugEnable {
					showWord(matrix, i, j)
				}
				ans++
			}
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
