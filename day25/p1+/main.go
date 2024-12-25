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
	// sc.Split(bufio.ScanWords)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	var locks, keys []uint

	for sc.Scan() { // first line '######' for lock or '.....' for key
		isLock := sc.Bytes()[0] == '#'
		value := uint(0)
		for i := 0; i < 5; i++ {
			if !sc.Scan() {
				panic(sc.Err())
			}
			for _, bit := range sc.Bytes() {
				value <<= 1
				if bit == '#' {
					value++
				}
			}
		}
		sc.Scan() // last line '.....' for lock or '######' for key
		sc.Scan() // skip empty line
		if isLock {
			locks = append(locks, value)
		} else {
			keys = append(keys, value)
		}
	}

	ans := 0
	for _, lock := range locks {
		for _, key := range keys {
			if lock&key == 0 {
				ans++
			}
		}
	}

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
