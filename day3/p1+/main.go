package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"

	. "adventofcode-2024/utils"
)

func scanMul(buf []byte) (a, b int, _ []byte) {
	var err error
	for {
		p := bytes.Index(buf, []byte("mul("))
		if p == -1 {
			return 0, 0, nil
		}
		buf = buf[p+4:]

		p = bytes.Index(buf, []byte(","))
		if p == -1 {
			continue
		}
		a, err = strconv.Atoi(UnsafeString(buf[:p]))
		if err != nil {
			continue
		}
		buf = buf[p+1:]

		p = bytes.Index(buf, []byte(")"))
		if p == -1 {
			continue
		}
		b, err = strconv.Atoi(UnsafeString(buf[:p]))
		if err != nil {
			continue
		}
		buf = buf[p+1:]

		return a, b, buf
	}
}

func solve(buf []byte) int {
	ans := 0
	var a, b int

	for {
		a, b, buf = scanMul(buf)
		if buf == nil {
			break
		}
		ans += a * b
	}

	return ans
}

func run(in io.Reader, out io.Writer) {
	buf, err := io.ReadAll(in)
	if err != nil {
		panic(err)
	}

	bw := bufio.NewWriter(out)
	defer bw.Flush()

	ans := solve(buf)
	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
