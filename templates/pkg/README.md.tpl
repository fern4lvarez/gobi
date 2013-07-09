# {{.FirstName}}
[Documentation online](http://godoc.org/{{.User}}/{{.FirstName}})

**{{.FirstName}}** is a package written in Go generated automatically by ´gobi´. Happy hacking! 

## Install (with GOPATH set on your machine)

* Step 1: Get the `{{.SecondName}}` package

```
go get {{.User}}/{{.Name}}
```

* Step 2 (Optional): Run tests

```
$ go test -v ./...
```

##Usage

### API

```
package main

import (
  "fmt"
  "{{.User}}/{{.Name}}"
)

func main() {
  // TODO
  }
```

##License
{{.FirstName}} is MIT licensed, see [here](https://{{.User}}/{{.FirstName}}/blob/master/LICENSE)