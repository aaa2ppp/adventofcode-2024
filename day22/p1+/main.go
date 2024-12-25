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

func solve( /*TODO*/ ) int {

	return 0 // TODO
}

func run(in io.Reader, out io.Writer) {
	// br := bufio.NewReader(in)
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanWords)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	ans := 0
	for {
		num, err := ScanInt(sc)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		ans += calc2000(num)
	}

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
