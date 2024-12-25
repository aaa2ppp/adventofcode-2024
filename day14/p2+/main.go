package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	. "adventofcode-2024/utils"
)

type point struct {
	x, y   int
	vx, vy int
}

func solve(bw *bufio.Writer, n, m int, pp []point, count int) int {
	matrix := MakeMatrix[byte](m, n+1)
	for _, row := range matrix {
		row[n] = '\n'
	}

	for t := 0; t < n*m; t++ {
		fmt.Fprintf(bw, "--- t: %d\n", t)
		
		for _, row := range matrix {
			for j := 0; j < n; j++ {
				row[j] = '.'
			}
		}

		for _, p := range pp {
			matrix[p.y][p.x] = 'x'
		}

		for _, row := range matrix {
			bw.Write(row)
		}

		bw.WriteByte('\n')
		for i := range pp {
			pp[i].x = ((pp[i].x+pp[i].vx)%n + n) % n
			pp[i].y = ((pp[i].y+pp[i].vy)%m + m) % m
		}
	}

	return 0
}

func run(in io.Reader, out io.Writer) {
	br := bufio.NewReader(in)
	// sc := bufio.NewScanner(in)
	// sc.Split(bufio.ScanWords)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	var n, m int
	_, err := fmt.Fscanf(br, "%dx%d\n", &n, &m)
	if err != nil {
		panic(err)
	}

	if debugEnable {
		log.Println("n:", n, "m:", m)
	}

	pp := make([]point, 0, 1024)

	for {
		var x, y, vx, vy int
		_, err := fmt.Fscanf(br, "p=%d,%d v=%d,%d\n", &x, &y, &vx, &vy)
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			}
		}
		pp = append(pp, point{x, y, vx, vy})
	}

	solve(bw, n, m, pp, 10)
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
