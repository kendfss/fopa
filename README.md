fopa
---

[Fo][src]rbidden [Pa][src]ths
Sanitize a file's path so that it has consistent across common operating systems and the web.

```go
package main

import "github.com/kendfss/fopa"

func main() {
    path := "%# "
    println(fopa.Sanitize(path)) // ___
    println(fopa.Sanitizef(path, "*")) // ***, but this is also forbidden
}
```


[src]: https://www.mtu.edu/umc/services/websites/writing/characters-avoid

### todo
- [ ] move scrape.go to internal
    - [ ] move Forbidden{Char,Rule} funcs to root
- [ ] [generate](https://github.com/dave/jennifer) {rule,char} table
- [ ] 
