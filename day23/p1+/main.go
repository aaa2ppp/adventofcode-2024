package main

import (
	"bufio"
	"io"
	"os"

	. "adventofcode-2024/utils"
)

func solve(graph map[uint16][]uint16) int {

	ans  := map[uint64]struct{}{}

	for k := range graph {
		if k>>8 == 't' {
			searchTriplex(graph, k, ans)
		}
	}

	return len(ans)
}

func searchTriplex(graph map[uint16][]uint16, k uint16, dst map[uint64]struct{}) {

	neigs := make(map[uint16]struct{}, len(graph[k]))
	visited := make(map[uint16]struct{}, len(graph[k]))

	for _, n := range graph[k] {
		neigs[n] = struct{}{}
	}

	for _, n1 := range graph[k] {
		visited[n1] = struct{}{}

		for _, n2 := range graph[n1] {
			if _, ok := visited[n2]; ok {
				continue
			}

			if _, ok := neigs[n2]; ok {
				a, b, c := k, n1, n2

				// sort
				if a > b {
					a, b = b, a
				}
				if b > c {
					b, c = c, b
				}
				if a > b {
					a, b = b, a
				}

				kk := uint64(a)<<32 + uint64(b)<<16 + uint64(c)
				dst[kk] = struct{}{}
			}
		}
	}
}

func run(in io.Reader, out io.Writer) {
	// br := bufio.NewReader(in)
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanWords)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	graph := map[uint16][]uint16{}
	toUint16 := func(b []byte) uint16 {
		return uint16(b[0])<<8 + uint16(b[1])
	}

	for sc.Scan() {
		a := toUint16(sc.Bytes()[0:])
		b := toUint16(sc.Bytes()[3:])
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}

	ans := solve(graph)

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
