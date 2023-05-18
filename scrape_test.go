package fopa_test

import (
	"bytes"
	"crypto/sha256"
	"io/ioutil"
	"testing"

	"github.com/kendfss/but"
	"github.com/kendfss/fopa"
	"github.com/kendfss/hcat"
)

const source = "https://www.mtu.edu/umc/services/websites/writing/characters-avoid"

func TestForbiddenChars(t *testing.T) {
	if fopa.ForbiddenChars() == nil {
		t.Log("fopa.ForbiddenChars is nil, but this should not be possible. Please verify logic.")
		t.Fail()
	}
}

func TestForbiddenRules(t *testing.T) {
	if fopa.ForbiddenRules() == nil {
		t.Log("fopa.ForbiddenRules is nil, but this should not be possible. Please verify logic.")
		t.Fail()
	}
}

func TestSource(t *testing.T) {
	fetched := sha256.Sum256(hcat.Read(hcat.Scrape(source)))
	tmp, err := ioutil.ReadFile(".testdata/source.htm")
	but.Must(err)
	local := sha256.Sum256(tmp)

	if !bytes.Equal(fetched[:], local[:]) {
		t.Logf("This package sources forbidden characters from %q, but the source's code has changed. Please update the following files:\n\t%q\n\t%q", source, "./.testdata/source.htm", "./scrape.go")
		t.Fail()
	}
}
