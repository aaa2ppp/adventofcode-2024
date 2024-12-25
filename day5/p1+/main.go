package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strconv"

	. "adventofcode-2024/utils"
)

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

		n := len(ww)
		aa := make([]int, n)

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

		if check(graph, aa) {
			m := aa[n/2]

			if debugEnable {
				log.Println("+", m)
			}

			ans += m
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
