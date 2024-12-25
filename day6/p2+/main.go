package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"

	. "adventofcode-2024/utils"
)

const (
	UP = iota
	RIGTH
	DOWN
	LEFT
)

var step = [][2]int{
	UP:    {-1, 0},
	RIGTH: {0, 1},
	DOWN:  {1, 0},
	LEFT:  {0, -1},
}

func showMatrix(matrix [][]byte) {
	n := len(matrix)
	m := len(matrix[0])
	row := make([]byte, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] < 1<<4 {
				row[j] = 'X'
			} else {
				row[j] = matrix[i][j]
			}
		}
		log.Printf("%s", row)
	}
	log.Println()
}

func solve(matrix [][]byte) int {
	i, j, dir := searchStart(matrix)
	if i == -1 {
		panic("start not found")
	}

	count := 0

	var dfs func(i, j, dir int)

	dfs = func(i, j, dir int) {
		defer func(v byte) { matrix[i][j] = v }(matrix[i][j])

		if matrix[i][j] >= 1<<4 {
			matrix[i][j] = 0
		}

		i2, j2, turnCount, finish := doStep(matrix, i, j, dir)

		if finish {
			return
		}

		if turnCount == 4 {
			count++
			if debugEnable {
				showMatrix(matrix)
			}
			return
		}

		dir = (dir + turnCount) % 4
		flag := byte(1) << dir

		if matrix[i][j]&flag != 0 {
			count++
			if debugEnable {
				showMatrix(matrix)
			}
			return
		}

		matrix[i][j] |= flag

		dfs(i2, j2, dir)
	}

	for {
		if matrix[i][j] >= 1<<4 {
			matrix[i][j] = 0
		}

		i2, j2, turnCount, finish := doStep(matrix, i, j, dir)

		if finish {
			break
		}

		if turnCount == 4 {
			panic("walled")
		}

		if turnCount != 0 && debugEnable {
			showMatrix(matrix)
		}

		dir = (dir + turnCount) % 4
		flag := byte(1) << dir

		if matrix[i][j]&flag != 0 {
			panic("loop")
		}

		if matrix[i2][j2] >= 1<<4 {
			func() {
				defer func(v byte) { matrix[i2][j2] = v }(matrix[i2][j2])
				matrix[i2][j2] = 'O'
				dfs(i, j, dir)
			}()
		}

		matrix[i][j] |= flag
		i, j = i2, j2
	}

	if debugEnable {
		showMatrix(matrix)
	}

	return count
}

func searchStart(matrix [][]byte) (i, j, dir int) {
	n := len(matrix)
	m := len(matrix[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch matrix[i][j] {
			case '^':
				return i, j, UP
			case 'v':
				return i, j, DOWN
			case '>':
				return i, j, RIGTH
			case '<':
				return i, j, LEFT
			}
		}
	}

	return -1, -1, 0
}

func doStep(matrix [][]byte, i, j, dir int) (i2, j2, turnCount int, finish bool) {
	n := len(matrix)
	m := len(matrix[0])

	for ; turnCount < 4; turnCount++ {
		i2 = i + step[dir][0]
		j2 = j + step[dir][1]

		if !(0 <= i2 && i2 < n && 0 <= j2 && j2 < m) {
			finish = true
			break
		}

		if matrix[i2][j2] != '#' && matrix[i2][j2] != 'O' {
			break
		}

		dir = (dir + 1) % 4
	}

	return i2, j2, turnCount, finish
}

func checkLoop(matrix [][]byte, i, j, dir int) bool {
	i, j, turnCount, finish := doStep(matrix, i, j, dir)

	if finish {
		return false
	}

	if turnCount == 4 {
		if debugEnable {
			showMatrix(matrix)
		}
		return true
	}

	dir = (dir + turnCount) % 4
	flag := byte(1) << dir

	if matrix[i][j] < 1<<4 && matrix[i][j]&flag != 0 {
		if debugEnable {
			showMatrix(matrix)
		}
		return true
	}

	return false
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
