package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"

	. "adventofcode-2024/utils"
)

type point struct {
	i, j int
}

func searchPath(maze [][]byte) (path []point) {

	cur := searchStart(maze)
	prev := point{-1, -1}

	for maze[cur.i][cur.j] != 'E' {
		path = append(path, cur)

		for _, of := range []point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			neig := point{cur.i + of.i, cur.j + of.j}
			if neig != prev && maze[neig.i][neig.j] != '#' {
				prev = cur
				cur = neig
				break
			}
		}
	}

	path = append(path, cur)

	return path
}

func searchStart(maze [][]byte) point {
	n := len(maze)
	m := len(maze[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if maze[i][j] == 'S' {
				return point{i, j}
			}
		}
	}
	return point{-1, -1}
}

func countCheats(path []point) map[int]int {

	count := make(map[int]int, len(path))

	for i, p1 := range path {
		for j := i + 2; j < len(path); j++ {
			p2 := path[j]

			s := Abs(p1.i-p2.i) + Abs(p1.j-p2.j)
			if s > 20 {
				continue
			}

			if t := j - i - s; t > 0 {
				count[t]++
			}
		}
	}

	return count
}

// bad result
func countCheats2(path []point) map[int]int {

	count := make(map[int]int, len(path))
	index := make(map[point]int, len(path))
	for i, p := range path {
		index[p] = i
	}

	for i, p1 := range path {
		for r := 1; r <= 20; r++ {
			for di, dj := 1, r-1; di <= r; di, dj = di+1, dj-1 {
				di, dj := di, dj
				for k := 0; k < 4; k++ {
					p2 := point{p1.i + di, p1.j + dj}

					j, ok := index[p2]
					if !ok || j < i {
						continue
					}

					s := Abs(p1.i-p2.i) + Abs(p1.j-p2.j)
					// if s > 20 {
					// 	continue
					// }

					if t := j - i - s; t > 0 {
						count[t]++
					}

					di, dj = -dj, di
				}
			}
		}
	}

	return count
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

	path := searchPath(maze)
	if debugEnable {
		log.Println("path:", path)
	}

	count := countCheats(path)
	if debugEnable {
		log.Println("count:", count)
	}

	sum := make([]int, len(path))
	for k, v := range count {
		sum[k] = v
	}

	for i := len(sum) - 2; i >= 0; i-- {
		sum[i] += sum[i+1]
	}

	var ans []int
	t, err := ScanInt(sc)
	for err == nil {
		ans = append(ans, sum[t])
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
