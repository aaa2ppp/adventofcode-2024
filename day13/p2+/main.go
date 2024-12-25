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
	x, y  int
	price int
}

type prize struct {
	x, y int
}

func solve(a, b button, p prize) int {

	switch {
	case a.y*p.x > a.x*p.y && b.y*p.x < b.x*p.y ||
		a.y*p.x < a.x*p.y && b.y*p.x > b.x*p.y:

		if debugEnable {
			log.Println("может быть не более одного решения")
		}

		n := (b.x*p.y - b.y*p.x)
		m := a.y*b.x - a.x*b.y
		if n%m != 0 {
			if debugEnable {
				log.Println("oops!..")
			}
			return 0
		}

		x := (n * a.x) / m
		y := (n * a.y) / m
		if (p.x-x)%b.x != 0 && (p.y-y)%b.y != 0 && (p.x-x)/b.x != (p.y-y)%b.y {
			if debugEnable {
				log.Println("oops!..")
			}
			return 0
		}

		cost := x/a.x*a.price + (p.x-x)/b.x*b.price
		if debugEnable {
			log.Printf("cost: %d %d %d", cost, x/a.x, (p.x-x)/b.x)
		}

		return cost

	case a.y*p.x == a.x*p.y && b.y*p.x == b.x*p.y:

		if debugEnable {
			log.Println("может быть много решений")
		}

		d := Gcd(a.x, b.x)
		if p.x%b.x%d != 0 {
			if debugEnable {
				log.Println("oops!..")
			}
			return 0
		}

		minimum := math.MaxInt
		for i := 0; i < 2; i++ {
			x := p.x % b.x
			for x%a.x != 0 {
				x += b.x
			}
			cost := x/a.x*a.price + (p.x-x)/b.x*b.price
			if debugEnable {
				if i == 0 {
					log.Printf("cost%d: %d %d %d", i, cost, x/a.x, (p.x-x)/b.x)
				} else {
					log.Printf("cost%d: %d %d %d", i, cost, (p.x-x)/b.x, x/a.x)
				}
			}
			minimum = min(minimum, cost)
			a, b = b, a
		}

		return minimum

	default:
		if debugEnable {
			log.Println("не может быть решений")
		}
		return 0
	}
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
		p := prize{x + 10000000000000, y + 10000000000000}

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
