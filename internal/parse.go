package internal

//go:generate go run ./parse

import (
	"bytes"
	_ "embed"
	"html"
	"strings"

	hq "github.com/antchfx/htmlquery"
	xhtml "golang.org/x/net/html"

	"github.com/kendfss/but"
)

const (
	SourceURL = "https://www.mtu.edu/umc/services/websites/writing/characters-avoid"
)

type Rule struct {
	Symb string `json:"symb"`
	Desc string `json:"name"`
}

var (
	doc *xhtml.Node

	//go:embed testdata/src.htm
	docBuf []byte
)

func init() {
	getDoc()
}

func GetForbiddenRules() (out []Rule) {
	for _, ul := range hq.Find(doc, "//ul") {
		if hq.SelectAttr(ul, "class") == "split" {
			items := hq.Find(ul, "//li")
			for _, li := range items {
				parts := strings.Fields(hq.InnerText(li))
				char := html.UnescapeString(parts[0])
				desc := strings.Join(parts[1:], " ")
				out = append(out, Rule{})
				out[len(out)-1].Symb = char
				if len(desc) > 0 {
					out[len(out)-1].Desc = desc
				} else {
					out[len(out)-1].Desc = char
				}
			}
		}
	}
	return
}

func GetForbiddenChars() (out []string) {
	for _, ul := range hq.Find(doc, "//ul") {
		if hq.SelectAttr(ul, "class") == "split" {
			items := hq.Find(ul, "//li")
			for _, li := range items {
				parts := strings.Fields(hq.InnerText(li))
				char := html.UnescapeString(parts[0])
				if len([]rune(char)) == 1 {
					out = append(out, char)
				}
			}
		}
	}
	return
}

func getDoc() {
	buf := bytes.NewBuffer(docBuf)
	var err error
	doc, err = hq.Parse(buf)
	// doc, err = hq.LoadURL(SourceFilePath)
	but.Must(err)
}
