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
	var buf2 []byte

	for {
		p := bytes.Index(buf, []byte("don't()"))

		if p == -1 {
			buf2 = buf
			buf = nil
		} else {
			buf2 = buf[:p]
			buf = buf[p+7:]
		}

		for {
			a, b, buf2 = scanMul(buf2)
			if buf2 == nil {
				break
			}
			ans += a * b
		}

		p = bytes.Index(buf, []byte("do()"))
		if p == -1 {
			break
		}
		buf = buf[p+4:]
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
