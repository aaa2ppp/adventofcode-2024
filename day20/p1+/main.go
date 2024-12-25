package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"

	. "adventofcode-2024/utils"
)

func showCount(count [][]int) {
	for _, row := range count {
		log.Printf("%4d", row)
	}
}

type point struct {
	i, j int
}

func countMase(maze [][]byte) (count [][]int, path []point) {
	n := len(maze)
	m := len(maze[0])

	// search start & finish
	var start, finish point
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch maze[i][j] {
			case 'S':
				start = point{i, j}
			case 'E':
				finish = point{i, j}
			}
		}
	}

	count = MakeMatrix[int](n, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			count[i][j] = -1
		}
	}

	var frontier Queue[point]
	frontier.Push(finish)
	count[finish.i][finish.j] = 0

	for !frontier.Empty() {
		p := frontier.Pop()
		if p == start {
			break
		}

		c := count[p.i][p.j] + 1
		for _, of := range []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			neig := point{p.i + of.i, p.j + of.j}

			if maze[neig.i][neig.j] == '#' {
				continue
			}

			if count[neig.i][neig.j] == -1 {
				count[neig.i][neig.j] = c
				frontier.Push(neig)
			}
		}
	}

	// restore path
	path = make([]point, 0, count[start.i][start.j]+1)
	p := finish

	for p != start {
		path = append(path, p)

		c := count[p.i][p.j] + 1
		for _, of := range []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			neig := point{p.i + of.i, p.j + of.j}
			if count[neig.i][neig.j] == c {
				p = neig
				break
			}
		}
	}
	path = append(path, p)

	return count, path
}

func countCheat(count [][]int, path []point, t int) int {
	n := len(count)
	m := len(count[0])

	ans := 0

	for _, p := range path {
		c := count[p.i][p.j] - 2

		for _, of := range []point{{-2, 0}, {-1, 1}, {0, 2}, {1, 1}, {2, 0}, {1, -1}, {0, -2}, {-1, -1}} {
		// for _, of := range []point{{-2, 0}, {0, 2}, {2, 0}, {0, -2}} {
			neig := point{p.i + of.i, p.j + of.j}

			if !(0 <= neig.i && neig.i < n && 0 <= neig.j && neig.j < m) {
				continue
			}

			if count[neig.i][neig.j] == -1 {
				continue
			}

			if n := c - count[neig.i][neig.j]; n >= t {
				ans++
			}
		}
	}

	return ans
}

func run(in io.Reader, out io.Writer) {
	// br := bufio.NewReader(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	buf, err := io.ReadAll(in)
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(bytes.TrimSpace(buf), []byte("\n"))

	sc := bufio.NewScanner(bytes.NewReader(lines[0]))
	sc.Split(bufio.ScanWords)
	maze := lines[1:]

	if err != nil {
		panic(err)
	}

	count, path := countMase(maze)
	if debugEnable {
		showCount(count)
		log.Println("path:", path)
	}

	var ans []int

	t, err := ScanInt(sc)
	for err == nil {
		ans = append(ans, countCheat(count, path, t))
		t, err = ScanInt(sc)
	}

	if err != io.EOF {
		panic(err)
	}

	WriteInts(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
