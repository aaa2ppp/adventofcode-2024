package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	. "adventofcode-2024/utils"
)

func solve(matrix [][]byte, program []byte) int {

	i, j := searchStart(matrix)
	if i == -1 {
		panic("can' found start")
	}

	step := 0
loop:
	for _, c := range program {

		var di, dj int
		switch c {
		case '^':
			di, dj = -1, 0
		case 'v':
			di, dj = 1, 0
		case '<':
			di, dj = 0, -1
		case '>':
			di, dj = 0, 1
		default:
			// skip other
			continue loop
		}

		if debugEnable {
			log.Println("--- Step:", step, "(i,j):", i, j)
		}

		i, j = doStep(matrix, i, j, di, dj)

		if debugEnable {
			showMatrix(matrix)
		}

		step++
	}

	return calcSumGPS(matrix)
}

func calcSumGPS(matrix [][]byte) int {
	n := len(matrix)
	m := len(matrix[0])

	sum := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 'O' {
				sum += i*100 + j
			}
		}
	}

	return sum
}

func doStep(matrix [][]byte, i, j, di, dj int) (int, int) {

	ii, jj := i+di, j+dj
loop:
	for {
		switch c := matrix[ii][jj]; c {
		case 'O':
			ii, jj = ii+di, jj+dj
			continue loop
		case '.':
			break loop
		case '#':
			return i, j
		default:
			panic(fmt.Errorf("unknown char: %c (%d,%d)", c, ii, jj))
		}
	}

	matrix[i][j] = '.'
	i, j = i+di, j+dj
	matrix[i][j] = '@'

	if i != ii || j != jj {
		matrix[ii][jj] = 'O'
	}

	return i, j
}

func searchStart(matrix [][]byte) (int, int) {
	n := len(matrix)
	m := len(matrix[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '@' {
				return i, j
			}
		}
	}

	return -1, -1
}

func showMatrix(matrix [][]byte) {
	for _, row := range matrix {
		log.Printf("%s\n", row)
	}
}

func run(in io.Reader, out io.Writer) {
	buf, err := io.ReadAll(in)
	if err != nil {
		panic(err)
	}
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	p := bytes.Index(buf, []byte("\n\n"))
	matrix := bytes.Split(buf[:p], []byte("\n"))

	if debugEnable {
		showMatrix(matrix)
	}

	program := buf[p+2:]

	ans := solve(matrix, program)

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
