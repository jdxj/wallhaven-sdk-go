package wallhaven_sdk_go

import (
	"fmt"
	"strings"
	"testing"
)

func TestSpace(t *testing.T) {
	str := ""
	runes := []rune(str)
	fmt.Printf("len: %d\n", len(runes))
}

func equalSlice(s1, s2 []string) bool {
	return strings.Compare(strings.Join(s1, ""), strings.Join(s2, "")) == 0
}

func TestTrimSpaceAndSplit(t *testing.T) {
	type data struct {
		name  string
		input string
		want  []string
	}

	cases := []data{
		{
			name:  "case1",
			input: "abc",
			want:  []string{"abc"},
		},
		{
			name:  "case2",
			input: " abc ",
			want:  []string{"abc"},
		},
		{
			name:  "case3",
			input: " abc 123 ",
			want:  []string{"abc", "123"},
		},
		{
			name:  "case4",
			input: "   abc   123   ",
			want:  []string{"abc", "123"},
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			result := TrimSpaceAndSplit(v.input)
			if !equalSlice(result, v.want) {
				t.Errorf("%s, input: %s, want: %s, get: %s\n", v.name, v.input, v.want, result)
			}
		})
	}
}
