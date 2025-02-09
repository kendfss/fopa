package fopa_test

import (
	"testing"

	"github.com/kendfss/fopa"
)

func TestClean(t *testing.T) {
	have := fopa.Clean("file#with@bad*chars.txt")
	want := "file_with_bad_chars.txt"
	if have != want {
		t.Logf("have %q, want %q", have, want)
	}

	have = fopa.Cleanf("file#with@bad*chars.txt", "-")
	want = "file-with-bad-chars.txt"
	if have != want {
		t.Logf("have %q, want %q", have, want)
	}
}

func TestSanitize(t *testing.T) {
	have := fopa.Sanitize("file#with@bad*chars.txt") // Replace forbidden chars
	want := "file_with_bad_chars.txt"
	if have != want {
		t.Logf("have %q, want %q", have, want)
	}
}

func TestRedux(t *testing.T) {
	have := fopa.Redux("file___with___chars.txt") // Remove consecutive replacements
	want := "file_with_chars.txt"
	if have != want {
		t.Logf("have %q, want %q", have, want)
	}
}
