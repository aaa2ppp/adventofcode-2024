package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"

	. "adventofcode-2024/utils"
)

func solve(aa []int, count int) int {
	ma := make(map[int]int, len(aa)*2)
	mb := make(map[int]int, len(aa)*2)

	for _, v := range aa {
		ma[v]++
	}

	var convbuf [24]byte

	for ; count > 0; count-- {

		if debugEnable {
			log.Println("len(ma):", len(ma))
		}

		clear(mb)

		for k, v := range ma {
			if k == 0 {
				mb[1] += v
				continue
			}

			b := strconv.AppendInt(convbuf[:0], int64(k), 10)
			s := UnsafeString(b)

			if n := len(s); n%2 == 0 {
				n /= 2

				k1, err := strconv.Atoi(s[:n])
				if err != nil {
					panic(err)
				}

				k2, err := strconv.Atoi(s[n:])
				if err != nil {
					panic(err)
				}

				mb[k1] += v
				mb[k2] += v

				continue
			}

			mb[k*2024] += v
		}

		ma, mb = mb, ma
	}

	n := 0
	for _, v := range ma {
		n += v
	}

	return n
}

func run(in io.Reader, out io.Writer) {
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanWords)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	var aa []int

	for {
		v, err := ScanInt(sc)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		aa = append(aa, v)
	}

	ans := solve(aa, 75)

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
