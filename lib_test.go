package fopa_test

import (
	"testing"

	"github.com/kendfss/fopa"
)

func TestClean(t *testing.T) {
	arg := "file#with@bad*chars.txt"
	have := fopa.Clean(arg)
	want := "file_with_bad_chars.txt"
	if have != want {
		t.Errorf("have %q (%d), want %q %d", have, len(have), want, len(want))
	}

	arg = "file#with@bad*chars.txt"
	have = fopa.Cleanf(arg, "-")
	want = "file-with-bad-chars.txt"
	if have != want {
		t.Errorf("have %q (%d), want %q %d", have, len(have), want, len(want))
	}
}

func TestSanitize(t *testing.T) {
	have := fopa.Sanitize("file#with@bad*chars.txt") // Replace forbidden chars
	want := "file_with_bad_chars.txt"
	if have != want {
		t.Errorf("have %q, want %q", have, want)
	}
}

func TestRedux(t *testing.T) {
	have := fopa.Redux("file___with___chars.txt") // Remove consecutive replacements
	want := "file_with_chars.txt"
	if have != want {
		t.Errorf("have %q, want %q", have, want)
	}
}
