package fopa

import (
	"github.com/kendfss/fopa/internal"
)

// type Rule struct {
// 	Symb, Desc string
// }

// var forbiddenRules []Rule

var Filler = "_"

// remove illegal characters from a file path
func Sanitize(path string) string {
	return Sanitizef(path, Filler)
}

// remove illegal characters from a file path
func Sanitizef(path, fill string) string {
	return sanitizef(path, fill)
}

// Redux remove runs of the fill character
func Redux(path string) string {
	return Reduxf(path, Filler)
}

// Redux removes runs of the fill character, with a format string
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

// // ForbiddenRules returns a slice of Symbol x Description pairs.
// func ForbiddenRules() []Rule {
// 	if forbiddenRules == nil {
// 		forbiddenRules = make([]Rule, len(internal.ForbiddenRules))
// 		for i, e := range internal.ForbiddenRules {
// 			forbiddenRules[i] = Rule{Symb: e.Symb, Desc: e.Desc}
// 		}
// 	}
// 	return forbiddenRules
// }

// ForbiddenChars returns a slice of the characters this library forbids
func ForbiddenChars() []string {
	return internal.ForbiddenChars
}
