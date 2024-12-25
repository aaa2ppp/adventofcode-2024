package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"

	. "adventofcode-2024/utils"
)

type button struct {
	x, y int
	price  int
}

type prize struct {
	x, y int
}

func solve(a, b button, p prize) int {
	minimum := math.MaxInt

	for x, y := 0, 0; x <= p.x && y <= p.y; x, y = x+a.x, y+a.y {
		if (p.x-x)%b.x == 0 && (p.y-y)%b.y == 0 && (p.x-x)/b.x == (p.y-y)/b.y {
			cost := (x/a.x)*a.price + (p.x-x)/b.x*b.price
			minimum = min(minimum, cost)
		}
	}

	if minimum == math.MaxInt {
		minimum = 0
	}

	return minimum
}

func run(in io.Reader, out io.Writer) {
	br := bufio.NewReader(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	ans := 0
	for {
		var x, y int

		if _, err := fmt.Fscanf(br, "Button A: X+%d, Y+%d\n", &x, &y); err != nil {
			if err == io.ErrUnexpectedEOF {
				break
			}
			panic(err)
		}
		a := button{x, y, 3}

		if _, err := fmt.Fscanf(br, "Button B: X+%d, Y+%d\n", &x, &y); err != nil {
			panic(err)
		}
		b := button{x, y, 1}

		if _, err := fmt.Fscanf(br, "Prize: X=%d, Y=%d\n", &x, &y); err != nil {
			panic(err)
		}
		p := prize{x, y}

		if debugEnable {
			log.Print(a, b, p)
		}

		ans += solve(a, b, p)

		fmt.Fscanln(br)
	}

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
