package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	// . "adventofcode-2024/utils"
)

type item struct {
	key   uint16
	count int
}

func solve(graph map[uint16]map[uint16]struct{}) []uint16 {

	order := make([]item, 0, len(graph))
	for k, v := range graph {
		order = append(order, item{k, len(v)})
	}

	sort.Slice(order, func(i, j int) bool {
		return order[i].count > order[j].count
	})

	if debugEnable {
		log.Println("order:", order)
	}

	maximum := 0
	var ans []uint16
	var group []uint16

	var dfs func(i int)

	dfs = func(i int) {
		if i == len(order) {
			if len(group) == maximum {
				ans = append(ans[:0], group...)
			}
			return
		}
		if len(group)+(len(order)-i) <= maximum {
			if len(group) == maximum {
				ans = append(ans[:0], group...)
			}
			return
		}
		k := order[i].key
		if len(graph[k]) <= maximum {
			if len(group) == maximum {
				ans = append(ans[:0], group...)
			}
			return
		}

		bingo := true
		for _, p := range group {
			if _, ok := graph[p][k]; !ok {
				bingo = false
				break
			}
		}

		if bingo {
			group = append(group, k)
			maximum = max(maximum, len(group))
			dfs(i + 1)
			group = group[:len(group)-1]
		}

		dfs(i + 1)
	}

	dfs(0)
	return ans
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

	graph := map[uint16]map[uint16]struct{}{}
	toUint16 := func(b []byte) uint16 {
		return uint16(b[0])<<8 + uint16(b[1])
	}

	for sc.Scan() {
		a := toUint16(sc.Bytes()[0:])
		b := toUint16(sc.Bytes()[3:])

		if _, ok := graph[a]; !ok {
			graph[a] = map[uint16]struct{}{}
		}
		if _, ok := graph[b]; !ok {
			graph[b] = map[uint16]struct{}{}
		}

		graph[a][b] = struct{}{}
		graph[b][a] = struct{}{}
	}

	if debugEnable {
		log.Println("graph:", graph)
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}

	ans := solve(graph)
	if debugEnable {
		log.Println("len(ans):", len(ans))
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})

	bw.WriteByte(byte(ans[0] >> 8))
	bw.WriteByte(byte(ans[0] & 0xff))

	for _, v := range ans[1:] {
		bw.WriteByte(',')
		bw.WriteByte(byte(v >> 8))
		bw.WriteByte(byte(v & 0xff))
	}
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
