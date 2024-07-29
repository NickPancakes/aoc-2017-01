package internal_test

import (
	"testing"

	"github.com/nickpancakes/aoc-2017-01/internal"
)

func TestSolveCaptcha(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			in:  "1122",
			out: "3",
		},
		{
			in:  "1111",
			out: "4",
		},
		{
			in:  "1234",
			out: "0",
		},
		{
			in:  "91212129",
			out: "9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := internal.SolveCaptcha(tt.in); got != tt.out {
				t.Errorf("SolveCaptcha(\"%s\") = \"%s\" , want \"%s\"", tt.in, got, tt.out)
			}
		})
	}
}
