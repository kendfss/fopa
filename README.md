# fopa

FoPa (forbidden paths) is a Go package that helps clean and sanitize file paths by removing problematic characters that could cause issues in file systems or URLs. It follows the character restrictions [guidelines from MTU][source]'s web services recommendations.

## Features

- Removes illegal/problematic characters from file paths
- Reduces consecutive occurrences of replacement characters
- Configurable replacement character
- Supports both single paths and path lists
- Command-line interface included

## Installation

```bash
go get github.com/kendfss/fopa
```

## Usage

### As a Package

```go
import "github.com/kendfss/fopa"

// Basic usage
cleanPath := fopa.Clean("file#with@bad*chars.txt")
// Result: "file_with_bad_chars.txt"

// Custom replacement character
cleanPath := fopa.Cleanf("file#with@bad*chars.txt", "-")
// Result: "file-with-bad-chars.txt"

// Individual operations
sanitized := fopa.Sanitize("file#with@bad*chars.txt")  // Replace forbidden chars
reduced := fopa.Redux("file___with___chars.txt")       // Remove consecutive replacements
```

### Command Line Interface

```bash
# Basic usage
fopa "file#with@bad*chars.txt"

# Process multiple files
fopa "file1#.txt" "file2*.txt"

# Custom fill character
fopa -f "-" "file#with@bad*chars.txt"

# Split paths before processing
fopa -s "/path/with#bad/chars:/another/path*"
```

### Piping Support

```bash
echo "file#1.txt\nfile*2.txt" | fopa
```

### api

```go
package fopa // import "github.com/kendfss/fopa"


// VARIABLES

var Filler = "_"

// FUNCTIONS

func Clean(path string) string
    Sanitize and Redux a filepath

func Cleanf(path, fill string) string
    Sanitize and Redux a filepath, with a format string

func ForbiddenChars() []string
    ForbiddenChars returns a slice of the characters this library forbids

func Join(parts ...string) string
    Join cleans each segment before joining the lot

func Redux(path string) string
    Redux remove runs of the fill character

func Reduxf(path, fill string) string
    Redux removes runs of the fill character, with a format string

func Sanitize(path string) string
    remove illegal characters from a file path

func Sanitizef(path, fill string) string
    remove illegal characters from a file path

func SplitClean(path string) string
    SplitClean splits the path before cleaning each segment
```

## Forbidden Characters

The following characters are automatically replaced:

- `#` (pound)
- `%` (percent)
- `&` (ampersand)
- `{` `}` (curly brackets)
- `\` (back slash)
- `<` `>` (angle brackets)
- `*` (asterisk)
- `?` (question mark)
- `/` (forward slash)
- `$` (dollar sign)
- `!` (exclamation point)
- `'` `"` (quotes)
- `:` (colon)
- `@` (at sign)
- `+` (plus sign)
- ``` ` ``` (backtick)
- `|` (pipe)
- `=` (equal sign)
- spaces and other whitespace
<!-- - emojis -->
<!-- - alt codes -->

## License

[BSD](./LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## todo
- [ ] emojis
- [ ] alt codes

[source]: https://www.mtu.edu/umc/services/websites/writing/characters-avoid
