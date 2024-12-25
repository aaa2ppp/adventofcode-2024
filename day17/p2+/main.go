package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	// . "adventofcode-2024/utils"
)

var (
	tasks = []struct {
		name string
		ra   int
		rb   int
		rc   int
		code []byte
		exec func(ra int, x, y int) byte
		x, y int
	}{
		{
			"1",
			0,
			0,
			0,
			[]byte{2, 4, 1, 1, 7, 5, 0, 3, 1, 4, 4, 0, 5, 5, 3, 0},
			exec,
			1,
			4,
		},
		{
			"2",
			0,
			0,
			0,
			[]byte{2, 4, 1, 6, 7, 5, 4, 6, 1, 4, 5, 5, 0, 3, 3, 0},
			exec,
			6,
			4,
		},
	}
)

// exec
// ---
// see ../p1/1.2.debug
// vm: RA:1 RB:0 RC:0 Code:[2 4 1 1 7 5 0 3 1 4 4 0 5 5 3 0]
// 00: BST RA 	// RA=1 RB=0 RC=0
// 02: BXL 1 	// RA=1 RB=1 RC=0
// 04: CDV RB 	// RA=1 RB=0 RC=0
// 06: ADV 3 	// RA=1 RB=0 RC=1
// 08: BXL 4 	// RA=0 RB=0 RC=1
// 10: BXC 0 	// RA=0 RB=4 RC=1
// 12: OUT RB 	// RA=0 RB=5 RC=1
// 14: JNZ 0 	// RA=0 RB=5 RC=1
// 5
func exec(ra, x, y int) byte {
	rb := ra & 7
	rb = rb ^ x
	rc := ra >> rb
	// ra >>= 3
	rb = rb ^ y
	rb = rb ^ rc
	return byte(rb & 7)
}

// exec2
// ----
// see ../p1/2.2.debug
// vm: &{RA:1 RB:0 RC:0 Code:[2 4 1 6 7 5 4 6 1 4 5 5 0 3 3 0] pos:0 instruction:[0xc48ec0 0xc48f80 0xc48fc0 0xc49060 0xc490a0 0xc490e0 0xc49200 0xc492c0] out:[]}
// 00: BST RA 	// RA=1 RB=0 RC=0
// 02: BXL 6 	// RA=1 RB=1 RC=0
// 04: CDV RB 	// RA=1 RB=7 RC=0
// 06: BXC RC 	// RA=1 RB=7 RC=0
// 08: BXL 4 	// RA=1 RB=7 RC=0
// 0a: OUT RB 	// RA=1 RB=3 RC=0
// 0c: ADV 3 	// RA=1 RB=3 RC=0
// 0e: JNZ 0 	// RA=0 RB=3 RC=0
// 3
func exec2(ra int) byte {
	rb := ra & 7
	rb = rb ^ 6
	rc := ra >> rb
	rb = rb ^ rc
	rb = rb ^ 4
	// ra >>= 3
	return byte(rb & 7)
}

func solve(code []byte, exec func(ra, x, y int) byte, x, y int) int {

	var dfs func(i int, ra int) (int, bool)

	dfs = func(i int, ra int) (int, bool) {
		if i == -1 {
			return ra, true
		}

		ra <<= 3
		for v := 0; v < 8; v++ {
			ra := ra | v
			if code[i] != exec(ra, x, y) {
				continue
			}
			if debugEnable {
				log.Printf("%02d: bingo! %b", i, ra)
			}
			if ra, ok := dfs(i-1, ra); ok {
				return ra, true
			}
		}

		if debugEnable {
			log.Printf("%02d: oops!..", i)
		}
		return -1, false
	}

	ans, _ := dfs(len(code)-1, 0)
	return ans
}

func run(_ io.Reader, out io.Writer) {

	bw := bufio.NewWriter(out)
	defer bw.Flush()

	for _, t := range tasks {
		ans := solve(t.code, t.exec, t.x, t.y)
		fmt.Printf("%s: %d\n", t.name, ans)
	}
}

// ----------------------------------------------------------------------------

var _, debugEnable = os.LookupEnv("DEBUG")

func main() {
	_ = debugEnable
	run(os.Stdin, os.Stdout)
}
