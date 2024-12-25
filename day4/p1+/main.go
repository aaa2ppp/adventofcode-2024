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

func showWord(matrix [][]byte, i, j int, of [2]int, k int) {
	n, m := len(matrix), len(matrix[0])

	count++
	log.Printf("%d --------", count)

	if of[0] == -1 {
		i -= k - 1
		of[0] = 1
	}
	if of[1] == -1 {
		j -= k - 1
		of[1] = 1
	}

	row := make([]byte, m)
	for ii := 0; ii < n; ii++ {
		for jj := 0; jj < m; jj++ {
			if k > 0 && ii == i && jj == j {
				row[jj] = matrix[ii][jj]
				i += of[0]
				j += of[1]
				k--
			} else {
				row[jj] = '.'
			}
		}
		log.Printf("%s", row)
	}
}

func check(matrix [][]byte, i, j int, of [2]int, word string) bool {
	n, m := len(matrix), len(matrix[0])

	for len(word) > 0 {
		if !(0 <= i && i < n && 0 <= j && j < m) {
			return false
		}
		if matrix[i][j] != word[0] {
			return false
		}
		i += of[0]
		j += of[1]
		word = word[1:]
	}

	return true
}

func solve(matrix [][]byte) int {
	n, m := len(matrix), len(matrix[0])
	ans := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for _, of := range [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
				if !check(matrix, i, j, of, word) {
					continue
				}

				if debugEnable {
					showWord(matrix, i, j, of, len(word))
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
