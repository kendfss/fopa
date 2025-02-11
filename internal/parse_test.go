package internal

import (
	"hash/crc32"
	"os"
	"regexp"
	"testing"

	"github.com/k0kubun/pp"
	"github.com/kendfss/but"
	"github.com/kendfss/hcat"
)

func TestForbiddenChars(t *testing.T) {
	chars := GetForbiddenChars()
	if chars == nil {
		t.Log("fopa.ForbiddenChars is nil, but this should not be possible. Please verify logic.")
		t.Fail()
		pp.Println(chars)
	}
}

func TestForbiddenRules(t *testing.T) {
	rules := GetForbiddenRules()
	if rules == nil {
		t.Log("fopa.ForbiddenRules is nil, but this should not be possible. Please verify logic.")
		t.Fail()
		pp.Println(rules)
	}
}

func TestSource(t *testing.T) {
	want := crc32.ChecksumIEEE(hcat.Read(hcat.Scrape(SourceURL)))
	tmp, err := os.ReadFile(SourceFilePath)
	but.Must(err)
	have := crc32.ChecksumIEEE(tmp)
	if have != want {
		t.Logf("This package sources forbidden characters from %q, but the source's code has changed. Please update the following files:\n\t%q\n\t%q", SourceURL, "./.testdata/source.htm", "./scrape.go")
		t.Fail()
	}
}

func TestRegex(t *testing.T) {
	_, err := regexp.Compile(Pattern)
	if err != nil {
		t.Error(err)
	}
}
