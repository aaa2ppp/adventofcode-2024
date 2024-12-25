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
			args{strings.NewReader(`Register C: 9

Program: 2,6,5,6`)},
			`1`,
			true,
		},
		{
			"2",
			args{strings.NewReader(`Register A: 10

Program: 5,0,5,1,5,4`)},
			`0,1,2`,
			true,
		},
		{
			"3",
			args{strings.NewReader(`Register A: 2024

Program: 0,1,5,4,3,0`)},
			`4,2,5,6,7,7,7,7,3,1,0`,
			true,
		},
		{
			"4",
			args{strings.NewReader(`Register B: 29

Program: 1,7,5,5`)},
			`2`, // 26 & 7 = 2
			true,
		},
		{
			"5",
			args{strings.NewReader(`Register B: 2024
Register C: 43690

Program: 4,0,5,5`)},
			`2`, // 44354 & 7 = 2
			true,
		},
		{
			"6",
			args{strings.NewReader(`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`)},
			`4,6,3,5,6,3,5,2,1,0`,
			false,
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
