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
		// {
		// 	"1",
		// 	args{strings.NewReader(``)},
		// 	``,
		// 	true,
		// },
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

func Test_solve(t *testing.T) {
	type args struct {
		aa    []int
		count int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		debug bool
	}{
		{
			"1",
			args{[]int{0, 1, 10, 99, 999}, 1},
			7,
			true,
		},
		{
			"2",
			args{[]int{125, 17}, 6},
			22,
			true,
		},
		{
			"3",
			args{[]int{125, 17}, 25},
			55312,
			true,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(v bool) { debugEnable = v }(debugEnable)
			debugEnable = tt.debug
			if got := solve(tt.args.aa, tt.args.count); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
