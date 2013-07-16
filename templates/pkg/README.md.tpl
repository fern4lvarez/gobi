# {{.FirstName}}
[Documentation online](http://godoc.org/{{.Host}}/{{.UserId}}/{{.FirstName}})

**{{.FirstName}}** is a package written in Go generated automatically by ´gobi´. Happy hacking! 

## Install (with GOPATH set on your machine)

* Step 1: Get the `{{.SecondName}}` package

```
go get {{.Host}}/{{.UserId}}/{{.Name}}
```

* Step 2 (Optional): Run tests

```
$ go test -v ./...
```

##Usage

```
package main

import (
  "fmt"
  "{{.Host}}/{{.UserId}}/{{.Name}}"
)

func main() {
  // TODO
  }
```

##License
{{.FirstName}} is {{.License}} licensed.