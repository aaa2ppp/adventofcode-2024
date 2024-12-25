package main

import (
	"bufio"
	"io"
	"os"

	. "adventofcode-2024/utils"
)

func solve(aa, bb []int) int {
	freq := make(map[int]int, len(bb))

	for _, v := range bb {
		freq[v]++
	}

	ans := 0
	for _, v := range aa {
		ans += v * freq[v]
	}

	return ans
}

func run(in io.Reader, out io.Writer) {
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanWords)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	var aa, bb []int

	for {
		a, b, err := ScanTwoInt(sc)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		aa = append(aa, a)
		bb = append(bb, b)
	}

	ans := solve(aa, bb)

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
