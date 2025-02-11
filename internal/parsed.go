// generated by go generate
// do not modify by hand

package internal

const Pattern = "#|%|&|\\{|\\}|\\\\|<|>|\\*|\\?|/|\\$|!|'|\"|\\:|@|\\+|`|\\||="

var Specials = map[string]struct{}{
	"$": {},
	"(": {},
	")": {},
	"*": {},
	"+": {},
	".": {},
	":": {},
	"?": {},
	"[": {},
	"\\": {},
	"]": {},
	"^": {},
	"{": {},
	"|": {},
	"}": {},
}

var ForbiddenRules = []Rule{
	{Symb: "#", Desc: "pound"},
	{Symb: "%", Desc: "percent"},
	{Symb: "&", Desc: "ampersand"},
	{Symb: "{", Desc: "left curly bracket"},
	{Symb: "}", Desc: "right curly bracket"},
	{Symb: "\\", Desc: "back slash"},
	{Symb: "<", Desc: "left angle bracket"},
	{Symb: ">", Desc: "right angle bracket"},
	{Symb: "*", Desc: "asterisk"},
	{Symb: "?", Desc: "question mark"},
	{Symb: "/", Desc: "forward slash"},
	{Symb: "blank", Desc: "spaces"},
	{Symb: "$", Desc: "dollar sign"},
	{Symb: "!", Desc: "exclamation point"},
	{Symb: "'", Desc: "single quotes"},
	{Symb: "\"", Desc: "double quotes"},
	{Symb: ":", Desc: "colon"},
	{Symb: "@", Desc: "at sign"},
	{Symb: "+", Desc: "plus sign"},
	{Symb: "`", Desc: "backtick"},
	{Symb: "|", Desc: "pipe"},
	{Symb: "=", Desc: "equal sign"},
	{Symb: "emojis", Desc: "emojis"},
	{Symb: "alt", Desc: "codes"},
}

var ForbiddenChars = []string{
	"#",
	"%",
	"&",
	"{",
	"}",
	"\\",
	"<",
	">",
	"*",
	"?",
	"/",
	"$",
	"!",
	"'",
	"\"",
	":",
	"@",
	"+",
	"`",
	"|",
	"=",
}
