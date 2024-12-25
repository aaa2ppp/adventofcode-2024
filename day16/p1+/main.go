package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"io"
	"log"
	"math"
	"os"

	. "adventofcode-2024/utils"
)

func showMtrix(matrix [][]byte) {
	for _, row := range matrix {
		log.Printf("%s\n", row)
	}
}

func showGraph(graph [][]edge) {
	for node, edges := range graph {
		if len(edges) > 0 {
			log.Printf("%d: %v", node, edges)
		}
	}
}

func solve(matrix [][]byte) int {
	n := len(matrix)
	m := len(matrix[0])
	node := func(i, j, b int) int { return (i*m+j)*2 + b } // b=0 вниз, b=1 влево

	matrix[n-2][0] = '.' // создаем вход

	start := node(n-2, 0, 1)
	finish1 := node(1, m-3, 1)
	finish2 := node(1, m-2, 0)

	graph := buildGraph(matrix)

	if debugEnable {
		showMtrix(matrix)
		showGraph(graph)
		log.Println("s,f1,f2", start, finish1, finish2)
	}

	items := make([]PriorityQueueItem, len(graph))
	frontier := PriorityQueue{}
	inQueue := make(map[*PriorityQueueItem]struct{})

	for i := range items {
		items[i] = PriorityQueueItem{node: i, weight: math.MaxInt}
	}
	items[start].weight = 0

	it := &items[start]
	heap.Push(&frontier, it)
	inQueue[it] = struct{}{}

	for len(frontier) > 0 {
		it := heap.Pop(&frontier).(*PriorityQueueItem)
		delete(inQueue, it)

		if debugEnable {
			log.Println("it:", *it, graph[it.node])
		}

		if it.node == finish1 || it.node == finish2 {
			break
		}

		for _, edge := range graph[it.node] {
			neigIt := &items[edge.neig]
			if w := it.weight + edge.weight; w < neigIt.weight {
				neigIt.weight = w
				if _, ok := inQueue[neigIt]; ok {
					heap.Fix(&frontier, neigIt.index)
				} else {
					heap.Push(&frontier, neigIt)
				}
			}
		}
	}

	return min(items[finish1].weight, items[finish2].weight)
}

// An PriorityQueueItem is something we manage in a priority queue.
type PriorityQueueItem struct {
	node   int // The value of the item; arbitrary.
	weight int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*PriorityQueueItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *PriorityQueueItem, priority int) {
	item.weight = priority
	heap.Fix(pq, item.index)
}

type edge struct {
	neig   int
	weight int
}

func buildGraph(matrix [][]byte) [][]edge {
	n := len(matrix)
	m := len(matrix[0])
	node := func(i, j, b int) int { return (i*m+j)*2 + b } // b=0 вниз, b=1 влево

	graph := make([][]edge, n*m*2)

	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {

			if matrix[i][j] == '#' {
				continue
			}

			if matrix[i-1][j] != '#' {
				a := node(i-1, j, 0)

				if matrix[i+1][j] != '#' {
					b := node(i, j, 0)
					graph[a] = append(graph[a], edge{b, 1})
					graph[b] = append(graph[b], edge{a, 1})
				}
				if matrix[i][j+1] != '#' {
					b := node(i, j, 1)
					graph[a] = append(graph[a], edge{b, 1001})
					graph[b] = append(graph[b], edge{a, 1001})
				}
				if matrix[i][j-1] != '#' {
					b := node(i, j-1, 1)
					graph[a] = append(graph[a], edge{b, 1001})
					graph[b] = append(graph[b], edge{a, 1001})
				}
			}

			if matrix[i][j-1] != '#' {
				a := node(i, j-1, 1)

				if matrix[i][j+1] != '#' {
					b := node(i, j, 1)
					graph[a] = append(graph[a], edge{b, 1})
					graph[b] = append(graph[b], edge{a, 1})
				}
				if matrix[i+1][j] != '#' {
					b := node(i, j, 0)
					graph[a] = append(graph[a], edge{b, 1001})
					graph[b] = append(graph[b], edge{a, 1001})
				}
			}

			if matrix[i+1][j] != '#' && matrix[i][j+1] != '#' {
				a := node(i, j, 0)
				b := node(i, j, 1)
				graph[a] = append(graph[a], edge{b, 1001})
				graph[b] = append(graph[b], edge{a, 1001})
			}
		}
	}

	return graph
}

func run(in io.Reader, out io.Writer) {
	buf, err := io.ReadAll(in)
	if err != nil {
		panic(err)
	}
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	matrix := bytes.Split(bytes.TrimSpace(buf), []byte("\n"))

	ans := solve(matrix)

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
