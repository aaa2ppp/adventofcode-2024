package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	. "adventofcode-2024/utils"
)


// x00 XOR y00 -> z00 
// x00 AND y00 -> d00

// x01 XOR y01 -> a01
// x01 AND y01 -> b01
// a01 XOR d00 -> z01
// a01 AND d00 -> c01
// b01 OR  c01 -> d01

// ...

// x(i) XOR y(i)   -> a(i)
// x(i) AND y(i)   -> b(i)
// a(i) XOR d(i-1) -> z(i)
// a(i) AND d(i-1) -> c(i)
// b(i) OR  c(i)   -> d(i)

// ...

// x(n) XOR y(n)   -> a(n)
// x(n) AND y(n)   -> b(n)
// a(n) XOR d(n-1) -> z(n)
// a(n) AND d(n-1) -> c(n)
// b(n) OR  c(n)   -> z(n+1)


type graph map[string]*node

type node struct {
	op    string
	left  string
	right string
}

func (node *node) String() string {
	return opKey(node.left, node.op, node.right)
}

func opKey(a, op, b string) string {
	if a > b {
		a, b = b, a
	}
	return fmt.Sprintf("{%s %s %s}", a, op, b)
}

func solve(graph graph, n int) []string {

	var ans []string

	swap := func(key1, key2 string) {
		if len(ans) == 8 {
			panic("too many swaps")
		}

		if debugEnable {
			log.Println("Swap:", key1, key2)
		}

		graph[key1], graph[key2] = graph[key2], graph[key1]
		ans = append(ans, key1, key2)
	}

	// XXX ручной подбор
	// swap("nqk", "z07")
	// swap("pcp", "fgt")
	// swap("fpq", "z24")
	// swap("srn", "z32")

mainLoop:
	for {
		ops := map[string]string{}
		for key, node := range graph {
			ops[opKey(node.left, node.op, node.right)] = key
		}

		var a, b, c, d, z string

		z = ops[opKey("x00", "XOR", "y00")]
		d = ops[opKey("x00", "AND", "y00")]

		if debugEnable {
			log.Printf("%d: d:%s z:%s", 0, d, z)
		}

		if z != "z00" {
			swap(z, "z00")
			continue mainLoop
		}

		for i := 1; i <= n; i++ {
			x := fmt.Sprintf("x%02d", i)
			y := fmt.Sprintf("y%02d", i)
			wantZ := fmt.Sprintf("z%02d", i)

			a = ops[opKey(x, "XOR", y)]
			z = ops[opKey(a, "XOR", d)]

			if z == "" {

				node := graph[wantZ]
				if node.left == d {
					swap(a, node.right)
					continue mainLoop
				}
				if node.right == d {
					swap(a, node.left)
					continue mainLoop
				}

				log.Printf("%d: Oops!.. a:%s d:%s, %v not found, want %s:%v", i, a, d, opKey(a, "XOR", d), wantZ, graph[wantZ])
				os.Exit(1)
			}

			if z != wantZ {
				swap(z, wantZ)
				continue mainLoop
			}

			b = ops[opKey(x, "AND", y)]
			c = ops[opKey(a, "AND", d)]
			d = ops[opKey(b, "OR", c)]

			if debugEnable {
				log.Printf("%d: a:%s b:%s c:%s d:%s z:%s", i, a, b, c, d, z)
			}
		}

		break
	}

	return ans
}

func run(in io.Reader, out io.Writer) {

	sc := bufio.NewScanner(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	graph := graph{}

	n := 0
	for sc.Scan() {

		// x01: 0

		b := bytes.TrimSpace(sc.Bytes())
		if len(b) == 0 {
			break
		}

		idx, err := strconv.Atoi(UnsafeString(b[1:3]))
		if err != nil {
			panic(err)
		}

		n = max(n, idx)
	}

	for sc.Scan() {

		// ntg XOR fgs -> mjb
		//  0   1   2  3   4

		token := bytes.Split(sc.Bytes(), []byte(" "))

		var (
			left  = string(token[0])
			op    = string(token[1])
			right = string(token[2])
			key   = string(token[4])
		)

		node := &node{
			op:    op,
			left:  left,
			right: right,
		}

		graph[key] = node
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}

	ans := solve(graph, n)

	sort.Strings(ans)
	bw.WriteString(strings.Join(ans, ","))
	bw.WriteByte('\n')
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
