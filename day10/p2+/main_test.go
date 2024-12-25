package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		in io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		debug   bool
	}{
// 		{
// 			"1",
// 			args{strings.NewReader(`
// ...0...
// ...1...
// ...2...
// 6543456
// 7.....7
// 8.....8
// 9.....9`)},
// 			`2`,
// 			true,
// 		},
// 		{
// 			"2",
// 			args{strings.NewReader(`
// ..90..9
// ...1.98
// ...2..7
// 6543456
// 765.987
// 876....
// 987....`)},
// 			`4`,
// 			true,
// 		},
// 		{
// 			"3",
// 			args{strings.NewReader(`
// 10..9..
// 2...8..
// 3...7..
// 4567654
// ...8..3
// ...9..2
// .....01`)},
// 			`3`,
// 			true,
// 		},
		{
			"4",
			args{strings.NewReader(`
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`)},
			`81`,
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(v bool) { debugEnable = v }(debugEnable)
			debugEnable = tt.debug
			out := &bytes.Buffer{}
			run(tt.args.in, out)
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
