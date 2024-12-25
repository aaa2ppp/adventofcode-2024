package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	. "adventofcode-2024/utils"
)

func solve(n, m, x, y, vx, vy, count int) int {
	x = ((x+vx*count)%n + n) % n
	y = ((y+vy*count)%m + m) % m

	if x < n/2 && y < m/2 {
		return 0
	}
	if x > n/2 && y < m/2 {
		return 1
	}
	if x < n/2 && y > m/2 {
		return 2
	}
	if x > n/2 && y > m/2 {
		return 3
	}

	return -1
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

	var q [4]int
	for {
		var x, y, vx, vy int
		_, err := fmt.Fscanf(br, "p=%d,%d v=%d,%d\n", &x, &y, &vx, &vy)
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				break
			}
		}
		i := solve(n, m, x, y, vx, vy, 100)
		if i >= 0 {
			q[i]++
		}
	}

	if debugEnable {
		log.Println("q:", q)
	}

	ans := q[0] * q[1] * q[2] * q[3]

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
