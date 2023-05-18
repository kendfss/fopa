package fopa

import (
	"html"
	"strings"

	hq "github.com/antchfx/htmlquery"
	xhtml "golang.org/x/net/html"

	"github.com/kendfss/but"
)

const (
	source = "https://www.mtu.edu/umc/services/websites/writing/characters-avoid"
	tag    = "//div[@class]"
	tagVal = "grid-x section-padding  medium-2-up xlarge-4-up "
)

var (
	doc       *xhtml.Node
	unescapes = map[string]string{
		"\u00a0": " ",
	}
)

func getForbiddenRules() (out []Rule) {
	doc, err := hq.LoadURL(source)
	but.Must(err)
	for _, div := range hq.Find(doc, tag) {
		if hq.SelectAttr(div, "class") == tagVal {
			ps := hq.Find(div, "//p")
			out = make([]Rule, len(ps))
			for i, p := range ps {
				parts := strings.Split(hq.InnerText(p), " ")
				out[i].Symb = unescapeString(parts[0])
				out[i].Desc = strings.Join(parts[1:], " ")
			}
		}
	}

	return
}

func ForbiddenRules() []Rule {
	if forbiddenRules == nil {
		if doc == nil {
			getDoc()
		}
		forbiddenRules = getForbiddenRules()
	}

	return forbiddenRules
}

func getForbiddenChars() (out []string) {
	doc, err := hq.LoadURL(source)
	but.Must(err)
	for _, div := range hq.Find(doc, tag) {
		if hq.SelectAttr(div, "class") == tagVal {
			ps := hq.Find(div, "//p")
			out = make([]string, len(ps))
			for i, p := range ps {
				out[i] = unescapeString(strings.Split(hq.InnerText(p), " ")[0])
			}
		}
	}

	return
}

func ForbiddenChars() []string {
	if forbiddenChars == nil {
		if doc == nil {
			getDoc()
		}
		forbiddenChars = Map(html.UnescapeString, getForbiddenChars()...)
	}

	return forbiddenChars
}

func getDoc() {
	var err error
	doc, err = hq.LoadURL(source)
	but.Must(err)
}

func Map[Arg, Val any](fn func(Arg) Val, args ...Arg) []Val {
	out := make([]Val, len(args))
	for i, elem := range args {
		out[i] = fn(elem)
	}
	return out
}

func unescapeString(arg string) string {
	str, ok := unescapes[arg]
	if !ok {
		str = arg
	}
	return str
}
