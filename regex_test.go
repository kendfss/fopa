package fopa

import (
	"regexp"
	"strings"
	"testing"

	"github.com/kendfss/fopa/internal"
)

func TestPat(t *testing.T) {
	parts := strings.Split(DefaultPattern, "|")
	have, want := len(parts), len(ForbiddenChars())+2 // +2 because of the additional space matcher and we get an extra empty string due to splitting on an escaped pipe (before we use another pipe to separate cases)
	if have != want {
		t.Errorf("have %d sub-patterns, want %d\n\t%q\n\t%q\n\t%q", have, want, DefaultPattern, parts, ForbiddenChars())
	}
}

func TestEach(t *testing.T) {
	re := regexp.MustCompile(DefaultPattern)
	want := Filler
	for _, char := range ForbiddenChars() {
		char := string(char)
		t.Run(char, func(t *testing.T) {
			pat := char
			if _, special := internal.Specials[char]; special {
				pat = "\\" + pat
			}
			have := re.ReplaceAllString(char, Filler)
			if have != want {
				t.Errorf("have %q, want %q", have, want)
			}
		})
	}
}
