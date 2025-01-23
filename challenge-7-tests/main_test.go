// challenges/testing/begin/main_test.go
package main

import (
	"fmt"
	"testing"
)

type counter interface {
	count(string) int
}

type testcase struct {
	input string
	want  int
}

func baseTest(t *testing.T, c counter, tests []testcase) {
	for i, tc := range tests {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			t.Parallel()
			got := c.count(tc.input)
			if tc.want != got {
				t.Fatalf("wanted %d but got %d", tc.want, got)
			}
		})
	}
}

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	baseTest(t, letterCounter{}, []testcase{
		{input: "HELLO", want: 5},
		{input: "CAFÉ", want: 4},
		{input: "1A2", want: 1},
		{input: "纳尼", want: 2},
		{input: ":A?2", want: 1},
	})
}

// write a test for numberCounter.count
func TestNumberCount(t *testing.T) {
	baseTest(t, numberCounter{}, []testcase{
		{input: "HELLO", want: 0},
		{input: "CAFÉ", want: 0},
		{input: "1A2", want: 2},
		{input: "纳尼", want: 0},
		{input: ":A?2", want: 1},
	})
}

// write a test for symbolCounter.count
func TestSymbolCount(t *testing.T) {
	baseTest(t, symbolCounter{}, []testcase{
		{input: "HELLO", want: 0},
		{input: "CAFÉ", want: 0},
		{input: "1A2", want: 0},
		{input: "纳尼", want: 0},
		{input: ":A?2", want: 2},
	})
}
