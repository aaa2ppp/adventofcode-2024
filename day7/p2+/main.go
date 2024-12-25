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

	var dfs func(v2 int, i int) bool

	dfs = func(v2 int, i int) bool {
		if v2 > v {
			return false
		}

		if i == len(aa) {
			return v2 == v
		}

		if dfs(v2+aa[i], i+1) {
			return true
		}

		if dfs(v2*aa[i], i+1) {
			return true
		}

		if aa[i] == 0 {
			v2 *= 10
		} else {
			for n := aa[i]; n > 0; n /= 10 {
				v2 *= 10
			}
		}

		if dfs(v2+aa[i], i+1) {
			return true
		}

		return false
	}

	return dfs(aa[0], 1)
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
