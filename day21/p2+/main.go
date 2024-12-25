package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	. "adventofcode-2024/utils"
)

type point struct {
	i, j int
}

// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+
// | 4 | 5 | 6 |
// +---+---+---+
// | 1 | 2 | 3 |
// +---+---+---+
//	   | 0 | A |
//	   +---+---+

var (
	keyb1 = map[byte]point{
		'7': {0, 0},
		'8': {0, 1},
		'9': {0, 2},
		'4': {1, 0},
		'5': {1, 1},
		'6': {1, 2},
		'1': {2, 0},
		'2': {2, 1},
		'3': {2, 2},
		'0': {3, 1},
		'A': {3, 2},
	}
)

// 029A -> <A^A>^^AvvvA, <A^A^>^AvvvA or <A^A^^>AvvvA
func encode1(seq string) (ans []string) {
	var buf []byte
	var dfs func(p point, s string, i int)

	dfs = func(p point, s string, i int) {
		for len(s) > 0 {
			p2 := keyb1[s[0]]

			if p.i != p2.i && p.j != p2.j && !(p.i == 3 && p2.j == 0) && !(p2.i == 3 && p.j == 0) {
				i := i
				if n := p2.i - p.i; n > 0 {
					buf = append(buf[:i], strings.Repeat("v", n)...)
					i += n
				}
				if n := p.j - p2.j; n > 0 {
					buf = append(buf[:i], strings.Repeat("<", n)...)
					i += n
				}
				if n := p.i - p2.i; n > 0 {
					buf = append(buf[:i], strings.Repeat("^", n)...)
					i += n
				}
				if n := p2.j - p.j; n > 0 {
					buf = append(buf[:i], strings.Repeat(">", n)...)
					i += n
				}
				buf = append(buf[:i], 'A')
				dfs(p2, s[1:], i+1)
			}

			if n := p2.j - p.j; n > 0 {
				buf = append(buf[:i], strings.Repeat(">", n)...)
				i += n
			}
			if n := p.i - p2.i; n > 0 {
				buf = append(buf[:i], strings.Repeat("^", n)...)
				i += n
			}
			if n := p.j - p2.j; n > 0 {
				buf = append(buf[:i], strings.Repeat("<", n)...)
				i += n
			}
			if n := p2.i - p.i; n > 0 {
				buf = append(buf[:i], strings.Repeat("v", n)...)
				i += n
			}

			buf = append(buf[:i], 'A')
			i++
			p = p2
			s = s[1:]
		}

		ans = append(ans, string(buf[:i]))
	}

	buf = append(buf, 'A')
	dfs(keyb1['A'], seq, 1)

	return ans
}

//     +---+---+
//     | ^ | A |
// +---+---+---+
// | < | v | > |
// +---+---+---+

// XXX почему можно выбирать только одну?
var keyb2 = map[string][]string{
	"^^": {"AA"},
	"AA": {"AA"},
	"<<": {"AA"},
	"vv": {"AA"},
	">>": {"AA"},
	"^A": {"A>A"},
	"^<": {"Av<A"},
	"^v": {"AvA"},
	// "^>": {"Av>A", "A>vA"},
	"^>": {"Av>A"},
	// "^>": {"A>vA"},
	"A^": {"A<A"},
	"A<": {"Av<<A"},
	// "Av": {"Av<A", "A<vA"},
	// "Av": {"Av<A"},
	"Av": {"A<vA"},
	"A>": {"AvA"},
	"<^": {"A>^A"},
	"<A": {"A>>^A"},
	"<v": {"A>A"},
	"<>": {"A>>A"},
	"v^": {"A^A"},
	// "vA": {"A^>A", "A>^A"},
	"vA": {"A^>A"},
	// "vA": {"A>^A"},
	"v<": {"A<A"},
	"v>": {"A>A"},
	// ">^": {"A^<"A, "A<^A"},
	// ">^": {"A^<A"},
	">^": {"A<^A"},
	">A": {"A^A"},
	"><": {"A<<A"},
	">v": {"A<A"},
}

func countKeyb2(depth int) map[string]int {
	ans := make(map[string]int)
	buf2 := make(map[string]int, len(keyb2))

	for key := range keyb2 {
		buf1 := make(map[string]int, len(keyb2))
		buf1[key] = 1

		for i := 1; i < depth; i++ {

			clear(buf2)

			for key2 := range buf1 {
				n := buf1[key2]
				seq2 := keyb2[key2][0] // XXX [0]

				for j := 1; j < len(seq2); j++ {
					key3 := seq2[j-1 : j+1]
					buf2[key3] += n
				}
			}

			if debugEnable {
				log.Printf("countKeyb2: %s %d %v", key, i, buf2)
			}

			buf1, buf2 = buf2, buf1
		}

		for key2, n := range buf1 {
			ans[key] += n * (len(keyb2[key2][0]) - 1) // XXX [0]
		}
	}

	return ans
}

func countString(s string, keyb map[string]int) int {
	ans := 0
	for i := 1; i < len(s); i++ {
		key := s[i-1 : i+1]
		ans += keyb[key]
	}
	return ans
}

func run(in io.Reader, out io.Writer, depth int) {
	// br := bufio.NewReader(in)
	sc := bufio.NewScanner(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	keyb2 := countKeyb2(depth)
	if debugEnable {
		log.Println("keyb2:", keyb2)
	}

	ans := 0
	for sc.Scan() {
		seq := sc.Text()
		v, err := strconv.Atoi(seq[:len(seq)-1])
		if err != nil {
			panic(err)
		}

		seqs1 := encode1(seq)
		minimum := math.MaxInt

		for _, s := range seqs1 {
			n := countString(s, keyb2)
			if debugEnable {
				log.Printf("%s: %d/%d %s", seq, len(s)-1, n, s)
			}
			minimum = min(minimum, n)
		}

		ans += v * minimum
	}

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout, 25)
}
