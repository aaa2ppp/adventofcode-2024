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
			log.Println("--- Step:", step, "(i,j):", i, j, "(di,dj):", di, dj)
		}

		if dj != 0 {
			i, j = doHorStep(matrix, i, j, dj)
		}

		if di != 0 {
			i, j = doVerStep(matrix, i, j, di)
		}

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
			if matrix[i][j] == '[' {
				sum += i*100 + j
				j++
			}
		}
	}

	return sum
}

func doHorStep(matrix [][]byte, i, j, dj int) (int, int) {

	ii, jj := i, j+dj
loop:
	for {
		switch c := matrix[ii][jj]; c {
		case '[', ']':
			jj += dj * 2
			continue loop
		case '.':
			break loop
		case '#':
			return i, j
		default:
			panic(fmt.Errorf("unknown char: %c (%d,%d)", c, ii, jj))
		}
	}

	for jj != j {
		matrix[ii][jj] = matrix[ii][jj-dj]
		jj -= dj
	}
	matrix[i][j] = '.'

	return i, j + dj
}

type doVerStepFunc func(matrix [][]byte, i, j, di int) (int, int)

var doVerStep = func() doVerStepFunc {

	// создаем замыкание для повторного использования frontier и visited

	type point struct {
		i, j int
	}

	var (
		frontier Queue[point]
		visited  Stack[point]
	)

	return func(matrix [][]byte, i, j, di int) (int, int) {

		frontier.Clear()
		visited.Clear()

		frontier.Push(point{i, j})
		visited.Push(point{i, j})

	loop:
		for !frontier.Empty() {
			p := frontier.Pop()

			neig := point{p.i + di, p.j}
			switch c := matrix[neig.i][neig.j]; c {
			case '#':
				return i, j
			case '.':
				continue loop
			case '[':
				//
			case ']':
				neig.j--
			default:
				panic(fmt.Errorf("unknown char: %c (%d,%d)", c, neig.i, neig.j))
			}

			if neig == visited.Top() {
				continue
			}

			frontier.Push(neig)
			visited.Push(neig)

			neig.j++
			frontier.Push(neig)
			visited.Push(neig)
		}

		for !visited.Empty() {
			p := visited.Pop()
			if matrix[p.i][p.j] != '.' {
				matrix[p.i+di][p.j] = matrix[p.i][p.j]
				matrix[p.i][p.j] = '.'
			}
		}

		return i + di, j
	}
}()

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

func transferMatrix(oldMatrix [][]byte) [][]byte {
	n := len(oldMatrix)
	m := len(oldMatrix[0])

	newMatrix := MakeMatrix[byte](n, m*2)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch c := oldMatrix[i][j]; c {
			case '#':
				newMatrix[i][j*2] = '#'
				newMatrix[i][j*2+1] = '#'
			case '.':
				newMatrix[i][j*2] = '.'
				newMatrix[i][j*2+1] = '.'
			case '@':
				newMatrix[i][j*2] = '@'
				newMatrix[i][j*2+1] = '.'
			case 'O':
				newMatrix[i][j*2] = '['
				newMatrix[i][j*2+1] = ']'
			default:
				panic(fmt.Errorf("transferMatrix: unknown char: %c (%d,%d)", c, i, j))
			}
		}
	}

	return newMatrix
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

	matrix = transferMatrix(matrix)

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
