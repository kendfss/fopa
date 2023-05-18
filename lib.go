package fopa

import (
// "golang.org/x/exp/slices"

// it "github.com/kendfss/iters"
)

type (
	Rule struct {
		Symb string `json:"symb"`
		Desc string `json:"name"`
	}
	derive struct{}
)

var (
	Filler = "_"

	forbiddenChars []string
	forbiddenRules []Rule
	// forbiddenChars = []string{"#", "%", "&", "{", "}", "\\", "<", ">", "*", "?", "/", " ", "\u00a0", "$", "!", "'", "\"", ":", "@", "+", "`", "|", "="}
	// forbiddenRules = []Rule{{Symb: "#", Desc: "pound"}, {Symb: "%", Desc: "percent"}, {Symb: "&", Desc: "ampersand"}, {Symb: "{", Desc: "left curly bracket"}, {Symb: "}", Desc: "right curly bracket"}, {Symb: "\\", Desc: "back slash"}, {Symb: "<", Desc: "left angle bracket"}, {Symb: ">", Desc: "right angle bracket"}, {Symb: "*", Desc: "asterisk"}, {Symb: "?", Desc: "question mark"}, {Symb: "/", Desc: "forward slash"}, {Symb: "\u00a0", Desc: "blank spaces"}, {Symb: " ", Desc: "blank spaces"}, {Symb: "$", Desc: "dollar sign"}, {Symb: "!", Desc: "exclamation point"}, {Symb: "'", Desc: "single quotes"}, {Symb: "\"", Desc: "double quotes"}, {Symb: ":", Desc: "colon"}, {Symb: "@", Desc: "at sign"}, {Symb: "+", Desc: "plus sign"}, {Symb: "`", Desc: "backtick"}, {Symb: "|", Desc: "pipe"}, {Symb: "=", Desc: "equal sign"}}
)

// remove illegal characters from a file path
func Sanitize(path string) string {
	return Sanitizef(path, Filler)
}

// remove illegal characters from a file path
func Sanitizef(path, fill string) string {
	return sanitizef(path, fill)
	// rack := forbiddenChars
	// if rack == nil {
	// 	rack = ForbiddenChars()
	// }

	// for _, char := range rack {
	// 	if string(os.PathSeparator) != char {
	// 		path = strings.ReplaceAll(path, char, filler)
	// 	}
	// }
	// return path
}

// remove runs of the fill character
func Redux(path string) string {
	return Reduxf(path, Filler)
}

// remove runs of the fill character, with a format string
func Reduxf(path, fill string) string {
	return reduxf(path, fill)
}

// Sanitize and Redux a filepath
func Clean(path string) string {
	return Cleanf(path, Filler)
}

// Sanitize and Redux a filepath, with a format string
func Cleanf(path, fill string) string {
	return cleanf(path, fill)
}
