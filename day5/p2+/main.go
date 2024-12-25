package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	. "adventofcode-2024/utils"
)

func solve(graph map[int][]int, aa []int) int {

	if check(graph, aa) {
		return 0
	}

	ma := make(map[int]struct{}, len(aa))
	for _, a := range aa {
		ma[a] = struct{}{}
	}

	graph2 := make(map[int][]int, len(aa))
	for _, a := range aa {
		// for _, b := range graph[a] {
		// 	if _, ok := ma[b]; ok {
		// 		graph2[a] = append(graph2[a], b)
		// 	}
		// }
		graph2[a] = graph[a] // и так сойдет...
	}

	if debugEnable {
		log.Println("graph2:", graph2)
	}

	order := sort(graph2)

	if debugEnable {
		log.Println("order:", order)
	}

	if !check(graph2, order) {
		panic("solve: bad order!")
	}

	m := len(aa) / 2
	for _, v := range order {

		if _, ok := ma[v]; !ok {
			continue
		}

		if m == 0 {

			if debugEnable {
				log.Println("+", v)
			}

			return v
		}

		m--
	}

	panic("solve: not found middle item!")
}

func check(graph map[int][]int, aa []int) bool {

	visited := make(map[int]struct{}, len(graph))

	for _, node := range aa {

		if _, ok := visited[node]; ok {

			if debugEnable {
				log.Printf("[%d] oops!..", node)
			}

			return false
		}

		if debugEnable {
			log.Printf("[%d]", node)
		}

		for _, neig := range graph[node] {
			if debugEnable {
				log.Println(neig)
			}

			visited[neig] = struct{}{}
		}
	}

	if debugEnable {
		log.Println("== ok")
	}

	return true
}

const (
	white = 0
	grey  = 1
	black = 2
)

func sort(graph map[int][]int) []int {

	var (
		dfs   func(node int) bool
		order = make([]int, 0, len(graph))
		color = make(map[int]int, len(graph))
		loop  = make([]int, 0, len(graph))
	)

	dfs = func(node int) bool {
		if color[node] == grey {
			loop = append(loop, node)
			return false
		}

		if color[node] == black {
			return true
		}

		color[node] = grey

		for _, neig := range graph[node] {
			if !dfs(neig) {
				loop = append(loop, node)
				return false
			}
		}

		color[node] = black
		order = append(order, node)

		return true
	}

	for node := range graph {
		if !dfs(node) {
			panic(fmt.Sprintf("sort: found loop: %v", loop))
		}
	}

	return order
}

func run(in io.Reader, out io.Writer) {

	sc := bufio.NewScanner(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	graph := make(map[int][]int, 100)

	for sc.Scan() {

		line := bytes.TrimSpace(sc.Bytes())
		if len(line) == 0 {
			break
		}

		ab := bytes.Split(line, []byte("|"))
		a, err := strconv.Atoi(UnsafeString(ab[0]))
		if err != nil {
			panic(err)
		}

		b, err := strconv.Atoi(UnsafeString(ab[1]))
		if err != nil {
			panic(err)
		}

		graph[b] = append(graph[b], a)
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}

	if debugEnable {
		log.Println("graph:", graph)
	}

	ans := 0

	for sc.Scan() {

		line := bytes.TrimSpace(sc.Bytes())
		ww := bytes.Split(line, []byte(","))
		aa := make([]int, len(ww))

		for i, w := range ww {
			a, err := strconv.Atoi(UnsafeString(w))

			if err != nil {
				panic(err)
			}

			aa[i] = a
		}

		if debugEnable {
			log.Println("aa:", aa)
		}

		ans += solve(graph, aa)
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
