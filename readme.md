fopa
---

[Fo]()rbidden [Pa]()ths
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
