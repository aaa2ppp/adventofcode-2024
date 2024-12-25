package main

import (
	"bufio"
	"io"
	"os"
	"sort"

	. "adventofcode-2024/utils"
)

func solve(aa, bb []int) int {
	sort.Ints(aa)
	sort.Ints(bb)

	ans := 0
	for i := 0; i < len(aa); i++ {
		ans += Abs(aa[i] - bb[i])
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
