package fopa

import (
	"fmt"
	"regexp"

	"github.com/kendfss/fopa/internal"
)

// const DefaultPattern = "#|%|&|\\{|\\}|\\|<|>|\\*|\\?|/| |\\$|!|'|\"|:|@|\\+|`|\\||=" + "\\s"
const DefaultPattern = internal.Pattern + `|\s`

func sanitizef(arg, fill string) string {
	re := regexp.MustCompile(DefaultPattern)
	return re.ReplaceAllString(arg, fill)
}

func reduxf(arg, fill string) string {
	re := regexp.MustCompile(fmt.Sprintf("%s+", fill))
	return re.ReplaceAllString(arg, fill)
}

func cleanf(arg, fill string) string {
	return reduxf(sanitizef(arg, fill), fill)
}
