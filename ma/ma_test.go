package main

import (
	"testing"
)

func TestIsDir(t *testing.T) {
	tcs := []struct {
		input  string
		wanted bool
	}{
		{
			input:  "a/b/c",
			wanted: true,
		},
		{
			input:  "a/b/c/",
			wanted: true,
		},
		{
			input:  "a.b/c/c.d",
			wanted: false,
		},
		{
			input:  "a/c.b/c.d",
			wanted: false,
		},
	}

	for _, tc := range tcs {
		if got := isDir(tc.input); got != tc.wanted {
			t.Errorf("input = %q; wanted = %v; got = %v", tc.input, tc.wanted, got)
		}
	}
}

func TestSplit(t *testing.T) {
	tcs := []struct {
		input  string
		wanted []string
	}{
		{
			input:  "a/b/c",
			wanted: []string{"a/b/c", ""},
		},
		{
			input:  "a/b/c/",
			wanted: []string{"a/b/c/", ""},
		},
		{
			input:  "a.b/c/c.d",
			wanted: []string{"a.b/c/", "c.d"},
		},
		{
			input:  "a/c.b/c.d",
			wanted: []string{"a/c.b/", "c.d"},
		},
	}

	for _, tc := range tcs {
		if d, f := split(tc.input); d != tc.wanted[0] || f != tc.wanted[1] {
			t.Errorf("input = %q; wanted = %v; got = %v", tc.input, tc.wanted, []string{d, f})
		}
	}
}
