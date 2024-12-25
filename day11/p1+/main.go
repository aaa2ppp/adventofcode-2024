package main

import (
	"bufio"
	"io"
	"os"
	"strconv"

	. "adventofcode-2024/utils"
)

func solve(aa []int, count int) int {
	bb := make([]int, 0, len(aa)*2)

	for ; count > 0; count-- {
		bb = bb[:0]

		for _, v := range aa {
			if v == 0 {
				bb = append(bb, 1)
				continue
			}

			s := strconv.Itoa(v)

			if n := len(s); n%2 == 0 {
				n /= 2
				v1, err := strconv.Atoi(s[:n])
				if err != nil {
					panic(err)
				}
				v2, err := strconv.Atoi(s[n:])
				if err != nil {
					panic(err)
				}
				bb = append(bb, v1, v2)
				continue
			}

			bb = append(bb, v*2024)
		}

		aa, bb = bb, aa
	}

	return len(aa)
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

	ans := solve(aa, 25)

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
