package main

import (
	"bufio"
	"io"
	"os"

	. "adventofcode-2024/utils"
)

func calc(num int) int {
	const mask = (1 << 24) - 1       // 16777216 - 1
	num = ((num << 6) ^ num) & mask  // *64
	num = ((num >> 5) ^ num) & mask  // /32
	num = ((num << 11) ^ num) & mask // *2048
	return num
}

func calc2000(num int) int {
	for i := 0; i < 2000; i++ {
		num = calc(num)
	}
	return num
}

func calcPriceN(num int, n int, buf []int) []int {
	if buf == nil {
		buf = make([]int, 0, n+1)
	}
	buf = append(buf, num%10)
	for i := 0; i < n; i++ {
		num = calc(num)
		buf = append(buf, num%10)
	}
	return buf
}

func searchSeq(prices []int) map[uint]int {
	p := prices
	ans := make(map[uint]int, len(prices))

	f := func(i int) uint {
		return uint(byte(int8(p[i] - p[i-1])))
	}

	const mask = (1 << 32) - 1
	var k uint
	k = f(1)
	k = (k << 8) + f(2)
	k = (k << 8) + f(3)

	for i := 4; i < len(prices); i++ {
		k = ((k << 8) + f(i)) & mask
		if _, ok := ans[k]; !ok {
			ans[k] = prices[i]
		}
	}

	return ans
}

func solve( /*TODO*/ ) int {

	return 0 // TODO
}

func run(in io.Reader, out io.Writer) {
	// br := bufio.NewReader(in)
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanWords)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	var (
		prices   []int
		unionSeq = map[uint]int{}
	)

	for {
		num, err := ScanInt(sc)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		prices = calcPriceN(num, 2000, prices[:0])
		seqs := searchSeq(prices)
		for k, v := range seqs {
			unionSeq[k] += v
		}
	}

	ans := 0
	for _, v := range unionSeq {
		ans = max(ans, v)
	}

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
