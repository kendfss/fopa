package fopa

import (
	"fmt"
	"regexp"
)

const pat = "#|%|&|\\{|\\}|\\|<|>|\\*|\\?|/| |\\$|!|'|\"|:|@|\\+|`|\\||=" + "\\s"

func sanitizef(arg, fill string) string {
	re := regexp.MustCompile(pat)
	return re.ReplaceAllString(arg, fill)
}

func reduxf(arg, fill string) string {
	re := regexp.MustCompile(fmt.Sprintf("%s+", fill))
	return re.ReplaceAllString(arg, fill)
}

func cleanf(arg, fill string) string {
	return reduxf(sanitizef(arg, fill), fill)
}
