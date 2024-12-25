package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"testing"

	. "adventofcode-2024/utils"
)

var testData = func() []byte {
	d, err := os.ReadFile("../data/1")
	if err != nil {
		panic(err)
	}
	return d
}()

var testAns = func() []byte {
	d, err := os.ReadFile("./1.a")
	if err != nil {
		panic(err)
	}
	return d
}()

func Test_run_solve(t *testing.T) {
	test_run(t, solve)
}

func Test_run_unionSolve(t *testing.T) {
	test_run(t, ufindSolve)
}

func Test_run_dequeSolve(t *testing.T) {
	test_run(t, dequeSolve)
}

func test_run(t *testing.T, solve func(n int, points [][2]int) int) {
	type args struct {
		in io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		debug   bool
	}{
		{
			"1",
			args{strings.NewReader(`7
12
5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`)},
			`6,1`,
			true,
		},
		{
			"2",
			args{bytes.NewReader(testData)},
			UnsafeString(testAns),
			false,
		},
		// {
		// 	"3",
		// 	args{strings.NewReader(``)},
		// 	``,
		// 	true,
		// },
		// {
		// 	"4",
		// 	args{strings.NewReader(``)},
		// 	``,
		// 	true,
		// },
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(v bool) { debugEnable = v }(debugEnable)
			debugEnable = tt.debug
			out := &bytes.Buffer{}
			run(tt.args.in, out, solve)
			if gotOut := out.String(); trimLines(gotOut) != trimLines(tt.wantOut) {
				t.Errorf("run() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func trimLines(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " \t\r\n")
	}
	for n := len(lines); n > 0 && lines[n-1] == ""; n-- {
		lines = lines[:n-1]
	}
	return strings.Join(lines, "\n")
}

func Benchmark_run(b *testing.B) {
	b.Run("solve", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			run(bytes.NewReader(testData), io.Discard, solve)
		}
	})
	b.Run("dequeSolve", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			run(bytes.NewReader(testData), io.Discard, dequeSolve)
		}
	})
	b.Run("unionSolve", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			run(bytes.NewReader(testData), io.Discard, ufindSolve)
		}
	})
}

func Benchmark_solve(b *testing.B) {

	bench := func(prefix, name string, n int, points [][2]int, solve func(n int, points [][2]int) int) {
		var ans int
		b.Run(prefix+name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ans = solve(n, points)
			}
		})
		b.Log("ans:", ans, points[ans])
	}

	batchBench := func(prefix string, n int, points [][2]int) {
		// bench(prefix, "binsearch+bfs", n, points, solve)
		bench(prefix, "binsearch+bfs2", n, points, dequeSolve)
		bench(prefix, "ufind", n, points, ufindSolve)
	}

	{
		sc := bufio.NewScanner(bytes.NewReader(testData))
		sc.Split(bufio.ScanWords)
		n, points := scanData(sc)
		batchBench("AoC:", n, points)
	}

	benchN := func(n int) {
		points := makeWallPoints(n)
		batchBench(fmt.Sprintf("wall%d:", n), n, points)
	}

	benchN(100)
	benchN(1000)
}

func makeRandomPoints(n int) [][2]int {
	points := make([][2]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			points[i*n+j] = [2]int{i, j}
		}
	}
	points = points[1 : len(points)-1]
	rand.Shuffle(len(points), func(i, j int) {
		points[i], points[j] = points[j], points[i]
	})
	return points
}

func makeWallPoints(n int) [][2]int {
	points := make([][2]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			points[i*n+j] = [2]int{i, j}
		}
	}
	points = points[1 : len(points)-1]
	for i, j := 0, n/2*n-1; i < n; i, j = i+1, j+1 {
		points[i], points[j] = points[j], points[i]
	}
	rand.Shuffle(len(points)-n, func(i, j int) {
		points[n+i], points[n+j] = points[n+j], points[n+i]
	})
	return points
}
