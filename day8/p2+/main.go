package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"

	. "adventofcode-2024/utils"
)

func showMatrix(matrix [][]byte) {
	for _, row := range matrix {
		log.Printf("%s", row)
	}
}

func showAntennas(antennas [][][2]int) {
	for c := '0'; c <= 'z'; c++ {
		if len(antennas[c]) > 0 {
			log.Printf("%c: %v", c, antennas[c])
		}
	}
}

func isAlphanum(c byte) bool {
	return ('0' <= c && c <= '9') || ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z')
}

func solve(matrix [][]byte) int {
	n := len(matrix)
	m := len(matrix[0])

	// search antennas
	antennas := make([][][2]int, 128)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			c := matrix[i][j]
			if isAlphanum(c) {
				antennas[c] = append(antennas[c], [2]int{i, j})
			}
		}
	}

	if debugEnable {
		showAntennas(antennas)
	}

	for i := '0'; i <= 'z'; i++ {
		placeAntinodes(matrix, antennas[i])
	}

	if debugEnable {
		showMatrix(matrix)
	}

	// count antennas
	count := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			c := matrix[i][j]
			if c == '#' {
				count++
			}
		}
	}

	return count
}

func placeAntinodes(matrix [][]byte, points [][2]int) {
	n := len(matrix)
	m := len(matrix[0])

	for k, p1 := range points {
		for _, p2 := range points[k+1:] {
			di := p2[0] - p1[0]
			dj := p2[1] - p1[1]
			for i, j := p1[0], p1[1]; 0 <= i && i < n && 0 <= j && j < m; i, j = i+di, j+dj {
				matrix[i][j] = '#'
			}
			for i, j := p1[0]-di, p1[1]-dj; 0 <= i && i < n && 0 <= j && j < m; i, j = i-di, j-dj {
				matrix[i][j] = '#'
			}
		}
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

	ans := solve(matrix)

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
