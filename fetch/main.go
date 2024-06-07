package main

import (
	"fmt"
	"regexp"
	"strings"

	"golang.design/x/clipboard"

	"github.com/kendfss/fopa"
)

func printChars() {
	fmt.Print("var forbiddenChars = []string{")
	for i, char := range fopa.ForbiddenChars() {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%q", char)
	}
	fmt.Print("}\n")
}

func copy_() {
	// but.Must()
	clipboard.Write(
		clipboard.FmtText,
		[]byte(
			strings.Join(
				fopa.Map(
					regexp.QuoteMeta,
					fopa.ForbiddenChars()...,
				),
				"|",
			),
		),
	)
}

func main() {
	// fmt.Printf("var forbiddenChars = %#v\n", fopa.ForbiddenChars())
	copy_()
	printChars()
	fmt.Printf("var forbiddenRules = %#v\n", fopa.ForbiddenRules())

	// resp, err := http.Get("https://www.mtu.edu/umc/services/websites/writing/characters-avoid")
	// but.Must(err)
	// println(resp.Status)
	// println(resp.StatusCode)
	// println(string(jsol.Prettify(resp.Header)))
	// data, err := ioutil.ReadAll(resp.Body)
	// but.Must(err)
	// fmt.Println(string(data))
}
