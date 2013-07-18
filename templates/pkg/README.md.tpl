# {{.FirstName}}

[Documentation online](http://godoc.org/{{.GoGetName}})

**{{.FirstName}}** is a package written in Go generated automatically by `gobi`. Happy hacking!

## Install (with GOPATH set on your machine)
----------

* Step 1: Get the `{{.SecondName}}` package

```
go get {{.GoGetName}}
```

* Step 2 (Optional): Run tests

```
$ go test -v ./...
```

##Usage
----------
```
package main

import (
  "fmt"
  "os"
  "{{.GoGetName}}"
)

func main() {
  {{.SecondName}}Example, err := {{.SecondName}}.New(1, "gobi")
  if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

  {{.SecondName}}Example.SetId({{.SecondName}}Example.Id() + 1)
  {{.SecondName}}Example.SetName({{.SecondName}}Example.Name() + " is great")

  fmt.Println({{.SecondName}}Example.Id(), {{.SecondName}}Example.Name())
  // Output: 2 gobi is great
}
```

##License
----------
{{.FirstName}} is {{.License}} licensed.