package main

import (
	"bufio"
	"io"
	"os"

	. "adventofcode-2024/utils"
)

func solve( /*TODO*/ ) int {

	return 0 // TODO
}

func run(in io.Reader, out io.Writer) {
	// br := bufio.NewReader(in)
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanWords)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	// TODO

	ans := solve( /*TODO*/ )

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
