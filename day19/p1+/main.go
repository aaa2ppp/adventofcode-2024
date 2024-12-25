package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	. "adventofcode-2024/utils"
)

func solve(towels, desings [][]byte) int {
	hashes := make(map[string]struct{}, len(towels))
	for _, t := range towels {
		hashes[UnsafeString(t)] = struct{}{}
	}

	count := 0
	for _, d := range desings {
		if check(UnsafeString(d), hashes) {
			count++
		}
	}

	return count
}

func check(d string, ts map[string]struct{}) bool {
	if debugEnable {
		log.Println("check:", d)
	}
	var dfs func(i int) bool

	visited := make([]bool, len(d))
	dfs = func(i int) bool {

		if i == len(d) {
			return true
		}

		if visited[i] {
			return false
		}
		visited[i] = true

		for l := 1; i+l <= len(d); l++ {
			s := d[i : i+l]
			if _, ok := ts[s]; ok {
				if debugEnable {
					log.Println("found:", i, s)
				}
				if dfs(i + l) {
					return true
				}
			}
		}

		return false
	}

	return dfs(0)
}

func run(in io.Reader, out io.Writer) {

	buf, err := io.ReadAll(in)
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(bytes.TrimSpace(buf), []byte("\n"))
	towels := bytes.Split(lines[0], []byte(", "))
	designs := lines[2:]

	if debugEnable {
		log.Printf("towels: %s", towels)
		log.Printf("designs: %s", designs)
	}

	ans := solve(towels, designs)

	fmt.Fprintln(out, ans)
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
