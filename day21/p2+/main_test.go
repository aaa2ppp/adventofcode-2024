package main

import (
	"bytes"
	"io"
	"slices"
	"strings"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		in    io.Reader
		depth int
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		debug   bool
	}{
		{
			"1",
			args{strings.NewReader(`029A
980A
179A
456A
379A`),
				2,
			},
			`126384`,
			true,
		},
		// {
		// 	"2",
		// 	args{strings.NewReader(``)},
		// 	``,
		// 	true,
		// },
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
			run(tt.args.in, out, tt.args.depth)
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

func Test_encode1(t *testing.T) {
	type args struct {
		seq string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		debug bool
	}{
		{
			"029A",
			args{"029A"},
			[]string{"<A^A>^^AvvvA", "<A^A^>^AvvvA", "<A^A^^>AvvvA"},
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(v bool) { debugEnable = v }(debugEnable)
			debugEnable = tt.debug
			if got := encode1(tt.args.seq); !slices.Contains(tt.want, got[0]) {
				t.Errorf("encode1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_encode2(t *testing.T) {
// 	type args struct {
// 		seq string
// 	}
// 	tests := []struct {
// 		name  string
// 		args  args
// 		want  string
// 		debug bool
// 	}{
// 		{
// 			"029A",
// 			args{"029A"},
// 			"v<<A>>^A<A>AvA<^AA>A<vAAA>^A",
// 			true,
// 		},
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			defer func(v bool) { debugEnable = v }(debugEnable)
// 			debugEnable = tt.debug
// 			if got := encode2(encode1(tt.args.seq)); got[0] != tt.want {
// 				t.Errorf("encode2() = %d %v, want %d %v", len(got), got, len(tt.want), tt.want)
// 			}
// 		})
// 	}
// }
