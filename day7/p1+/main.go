package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	. "adventofcode-2024/utils"
)

func check(v int, aa []int) bool {
	n := 1 << len(aa)

	for f := 0; f < n; f++ {
		if debugEnable {
			log.Printf("%08b", f)
		}

		v2 := aa[0]

		for i, a := range aa[1:] {

			if f&(1<<i) == 0 {
				v3 := v2 * a

				if debugEnable {
					log.Printf("%d * %d = %d", v2, a, v3)
				}

				v2 = v3

			} else {
				v3 := v2 + a

				if debugEnable {
					log.Printf("%d + %d = %d", v2, a, v3)
				}

				v2 = v3
			}

			if v2 > v {
				break
			}
		}

		if v2 == v {
			return true
		}
	}
	return false
}

func run(in io.Reader, out io.Writer) {
	sc := bufio.NewScanner(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	ans := 0

	for sc.Scan() {

		line := UnsafeString(sc.Bytes())

		if debugEnable {
			log.Println("***", line)
		}

		p := strings.Index(line, ":")
		v, err := strconv.Atoi(line[:p])

		if err != nil {
			panic(err)
		}

		ww := strings.Split(strings.TrimSpace(line[p+1:]), " ")
		aa := make([]int, len(ww))

		for i, w := range ww {
			v, err := strconv.Atoi(w)
			if err != nil {
				panic(err)
			}
			aa[i] = v
		}

		if check(v, aa) {
			ans += v
		}
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
