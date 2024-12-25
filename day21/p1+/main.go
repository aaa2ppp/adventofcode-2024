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

var keyb2 = map[string][]string{
	"^^": {""},
	"AA": {""},
	"<<": {""},
	"vv": {""},
	">>": {""},
	"^A": {">"},
	"^<": {"v<"},
	"^v": {"v"},
	"^>": {"v>", ">v"},
	"A^": {"<"},
	"A<": {"v<<"},
	"Av": {"v<", "<v"},
	"A>": {"v"},
	"<^": {">^"},
	"<A": {">>^"},
	"<v": {">"},
	"<>": {">>"},
	"v^": {"^"},
	"vA": {"^>", ">^"},
	"v<": {"<"},
	"v>": {">"},
	">^": {"^<", "<^"},
	">A": {"^"},
	"><": {"<<"},
	">v": {"<"},
}

func encode2(seq string, ans []string) []string {
	var buf []byte
	var dfs func(seq string, i int)

	dfs = func(s string, i int) {
		for len(s) > 1 {
			key := s[0:2]
			vals := keyb2[key]

			for len(vals) > 1 {
				i := i
				buf = append(buf[:i], vals[0]...)
				i += len(vals[0])
				buf = append(buf[:i], 'A')
				i++
				dfs(s[1:], i)
				vals = vals[1:]
			}

			buf = append(buf[:i], vals[0]...)
			i += len(vals[0])
			buf = append(buf[:i], 'A')
			i++
			s = s[1:]
		}

		ans = append(ans, string(buf[:i]))
	}

	buf = append(buf, 'A')
	dfs(seq, 1)
	return ans
}

func solve(seq string, n int) int {
	// v, err := strconv.Atoi(seq[:len(seq)-1])
	// if err != nil {
	// 	panic(err)
	// }

	// seq = encode1(seq)
	// for i := 0; i < n; i++ {
	// 	seq = encode2(seq)
	// }

	// if debugEnable {
	// 	log.Println(len(seq), seq)
	// }

	// return len(seq) * v
	return 0 // TODO
}

func run(in io.Reader, out io.Writer) {
	// br := bufio.NewReader(in)
	sc := bufio.NewScanner(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	ans := 0

	var buf1, buf2 []string

	for sc.Scan() {

		seq := sc.Text()

		v, err := strconv.Atoi(seq[:len(seq)-1])
		if err != nil {
			panic(err)
		}

		buf1 = append(buf1[:0], encode1(seq)...)

		for i := 0; i < 2; i++ {
			buf2 = buf2[:0]

			for _, seq := range buf1 {
				buf2 = encode2(seq, buf2)
			}

			buf1, buf2 = buf2, buf1
		}

		minimum := math.MaxInt
		for _, seq := range buf1 {
			minimum = min(minimum, len(seq)-1)
		}

		if debugEnable {
			count := 0
			var last string
			for _, seq := range buf1 {
				if len(seq)-1 == minimum {
					// log.Println(seq, len(s), s)
					count++
					last = seq
				}
			}
			log.Printf("%s %d 1/%d/%d: %s", seq, minimum, count, len(buf1), last)
		}

		ans += v * minimum
	}

	WriteInt(bw, ans, DefaultWriteOpts())
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
