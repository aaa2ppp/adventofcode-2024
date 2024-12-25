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
		{
			"1",
			args{strings.NewReader(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`)},
			`3749`,
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

func Test_check(t *testing.T) {
	type args struct {
		v  int
		aa []int
	}
	tests := []struct {
		name string
		args args
		want bool
		debug bool
	}{
		{
			"4: 1 1 2",
			args{4, []int{1, 1, 2}},
			true,
			true,
		},
		{
			"190: 10 19",
			args{190, []int{10, 19}},
			true,
			true,
		},
		{
			"3267: 81 40 27",
			args{3267, []int{81, 40, 27}},
			true,
			true,
		},
		{
			"83: 17 5",
			args{83, []int{17, 5}},
			false,
			true,
		},
		{
			"156: 15 6",
			args{156, []int{15, 6}},
			false,
			true,
		},

		{
			"7290: 6 8 6 15",
			args{7290, []int{6, 8, 6, 15}},
			false,
			true,
		},
		// 161011: 16 10 13
		// 192: 17 8 14
		// 21037: 9 7 18 13
		{
			"292: 11 6 16 20",
			args{292, []int{11, 6, 16, 20}},
			true,
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(v bool) { debugEnable = v }(debugEnable)
			debugEnable = tt.debug
			if got := check(tt.args.v, tt.args.aa); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
			}
		})
	}
}
