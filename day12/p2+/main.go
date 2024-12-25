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
	flags := MakeMatrix[byte](n, m) // 128 -visited, 1, 2, 4, 8 - top, rigth, bottom. left

	ans := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited := flags[i][j] & 128; visited == 0 {
				area, corners := countArea(matrix, flags, i, j)
				ans += area * corners
			}
		}
	}

	return ans
}

func countArea(matrix, flags [][]byte, i, j int) (area, corners int) {
	n, m := len(matrix), len(matrix[0])

	var frontier Queue[[2]int]
	frontier.Push([2]int{i, j})
	flags[i][j] |= 128 // mark visited

	var cells [][2]int

	// обходим соседние по часовой стрелке: top, rigth, bottom, left
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for !frontier.Empty() {
		p := frontier.Pop()
		i, j := p[0], p[1]
		v := matrix[i][j]

		area++
		cells = append(cells, [2]int{i, j})

		for k, of := range dirs {

			ii, jj := i+of[0], j+of[1]

			if !(0 <= ii && ii < n && 0 <= jj && jj < m) {
				flags[i][j] |= 1 << k // mark side
				continue
			}

			if neigV := matrix[ii][jj]; neigV != v {
				flags[i][j] |= 1 << k // mark side
				continue
			}

			if visited := flags[ii][jj] & 128; visited == 0 {
				flags[ii][jj] |= 128 // mark visited
				frontier.Push([2]int{ii, jj})
			}
		}
	}

	for _, c := range cells {
		i, j := c[0], c[1]
		f := flags[i][j]

		k1 := 3
		side1 := f & (1 << k1)
		for k2 := 0; k2 < 4; k2++ {

			side2 := f & (1 << k2)

			if side1 != 0 && side2 != 0 {
				// внешний угол
				corners++
			}

			if side1 == 0 && side2 == 0 {

				of1 := dirs[k1]
				i1, j1 := i+of1[0], j+of1[1]
				f1 := flags[i1][j1]

				of2 := dirs[k2]
				i2, j2 := i+of2[0], j+of2[1]
				f2 := flags[i2][j2]

				if f1&(1<<k2) != 0 && f2&(1<<k1) != 0 {
					// внутренний угол
					corners++
				}
			}

			k1, side1 = k2, side2
		}
	}

	return area, corners
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
