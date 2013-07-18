# {{.FirstName}}
=====
[Documentation online](http://godoc.org/{{.GoGetName}})

**{{.SecondName}}** is a web application written in Go generated automatically by `gobi`. It contains a Procfile and a .godir, so it's ready to be deployed on platforms like Heroku or cloudControl. Happy hacking!

## Install (with GOPATH set on your machine)
----------

* Step 1: Get the `{{.SecondName}}` package

```
go get {{.GoGetName}}
```

* Step 2 (Optional): Run tests

```
$ go test -v .
```

##Usage
----------

* Run it locally

```
$ {{.SecondName}}
Listening on 5555 ...
```

##License
----------
{{.FirstName}} is {{.License}} licensed.