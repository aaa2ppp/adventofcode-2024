package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"slices"
	"strconv"

	. "adventofcode-2024/utils"
)

func check1(report []int) bool {

	prevV := report[0]
	prevS := Sign(report[1] - prevV)

	if prevS == 0 {
		return false
	}

	for i := 1; i < len(report); i++ {
		v := report[i]
		d := v - prevV
		s := Sign(d)
		d = Abs(d)

		if s != prevS {
			return false
		}

		if d > 3 {
			return false
		}

		prevV = v
	}

	return true
}

func check2(report []int) bool {
	if check1(report) {
		return true
	}

	for i := range report {
		if check1(copyIgnore(report, i)) {
			return true
		}
	}

	return false
}

func copyIgnore(src []int, i int) []int {
	dst := make([]int, 0, len(src)-1)
	for j, v := range src {
		if i != j {
			dst = append(dst, v)
		}
	}
	return dst
}

func solve(reports [][]int) int {
	ans := 0

	for i, report := range reports {

		if check2(report) {
			ans++
			if debugEnable {
				log.Printf("%d: %v yes", i+1, report)
			}
		} else {
			if debugEnable {
				log.Printf("%d: %v no", i+1, report)
			}
		}
	}
	return ans
}

func run(in io.Reader, out io.Writer) {
	sc := bufio.NewScanner(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	var reports [][]int
	var report []int

	for sc.Scan() {
		line := bytes.TrimSpace(sc.Bytes())
		report = report[:0]

		for _, s := range bytes.Split(line, []byte{' '}) {
			v, err := strconv.Atoi(UnsafeString(s))
			if err != nil {
				panic(err)
			}
			report = append(report, v)
		}

		reports = append(reports, slices.Clone(report))
	}

	ans := solve(reports)

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
