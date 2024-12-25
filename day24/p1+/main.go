package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	. "adventofcode-2024/utils"
)

type graph map[string]*node

type node struct {
	graph      graph
	op         string
	left       string
	right      string
	calculated bool
	value      bool
}

func (node *node) String() string {
	return fmt.Sprintf("{%v %v %v %v %v}", node.op, node.left, node.right, node.calculated, node.value)
}

func (node *node) Value() bool {
	if !node.calculated {
		left := node.graph[node.left]
		right := node.graph[node.right]
		switch node.op {
		case "OR":
			node.value = left.Value() || right.Value()
		case "AND":
			node.value = left.Value() && right.Value()
		case "XOR":
			node.value = left.Value() != right.Value()
		}
	}
	return node.value
}

type item struct {
	key  string
	node *node
}

func (it item) String() string {
	return fmt.Sprintf("{%v %v}", it.key, it.node)
}

func solve(graph graph) int {

	order := make([]item, 0)

	for key, node := range graph {
		if key[0] == 'z' {
			node.Value() // calculate value to debug
			order = append(order, item{key, node})
		}
	}

	sort.Slice(order, func(i, j int) bool {
		return order[i].key > order[j].key
	})

	if debugEnable {
		log.Println("order:", order)
	}

	ans := 0
	for _, it := range order {
		ans <<= 1
		if it.node.Value() {
			ans++
		}
	}

	return ans
}

func run(in io.Reader, out io.Writer) {

	sc := bufio.NewScanner(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	graph := make(graph)

	for sc.Scan() {

		// x01: 0

		b := bytes.TrimSpace(sc.Bytes())
		if len(b) == 0 {
			break
		}

		key := string(b[:3])
		switch c := b[5]; c {
		case '0':
			graph[key] = &node{graph: graph, calculated: true, value: false}
		case '1':
			graph[key] = &node{graph: graph, calculated: true, value: true}
		default:
			panic(fmt.Errorf("%c: unknown value", c))
		}
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

		if left > right {
			left, right = right, left
		}

		node := &node{
			graph: graph,
			op:    op,
			left:  left,
			right: right,
		}

		graph[key] = node
	}

	if debugEnable {
		log.Println("graph:", graph)
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
