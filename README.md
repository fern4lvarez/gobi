```
              _     _ 
             | |   (_)
   __ _  ___ | |__  _ 
  / _' |/ _ \| '_ \| |
 | (_| | (_) | |_) | |
  \__' |\___/|_.__/|_|
   __/ |              
  |___/
```

[![Build Status](https://drone.io/github.com/fern4lvarez/gobi/status.png)](https://drone.io/github.com/fern4lvarez/gobi/latest) 
[Documentation online](http://godoc.org/github.com/fern4lvarez/gobi)

`gobi` is a command line tool that will make your Go development just faster. 
It might stand for *Go Boilerplater Injector*, but just think of it as a fun, tiny tool to create and manage quickly your applications written in Go.

`gobi` is on current development, so it's planned to get more features as long as time passes by. As of version 0.1.x these are the main current features:

* Create command line applications ready to use.
* Create Go packages with a basic test suite and example included.
* Create a web application with Bootstrap assets and ready to deploy on most popular PaaS.
* Two-level path projects.
* Create your profile with your desired configuration.
* LICENSE, README, VERSION, .gitignore and other files included out of the box.


## Install (with GOPATH set on your machine)
----------

* Step 1: Get the package. Then you will be able to use `gobi` as an executable. 

```
go get github.com/fern4lvarez/gobi
```

* Step 2 (Optional): Run tests

```
$ go test -v .
```

##Usage
-------

You'll have to introduce your configuration the first time you use `gobi`:
```
$ gobi
No configuration found! I'd like to know more about you. 
Name: // Your real name.
Username: // Your user name.
Host: // Host of your projects. Currently only github.com, bitbucket.org and code.google.com are supported.
Email: // Your email address.
License: The license applying to your projects. (Supporting AGPL, Apache, BSD, BSD3-Clause, Eclipse, GPLv2, GPLv3, LGPLv2.1, LGPLv3, MIT, Mozilla, PublicDomain, WTFPL and no-license)
```

A file called `.gobi.json` will be created on your `$HOME` directory containing all your configuration. If you want to restart your configuration, you have to remove this file and execute `gobi` again. Dynamic management of the configuration is planned to be implemented.

If you need help:
```
$ gobi help
```

If you want to know the current `gobi` version:
```
$ gobi version
```

If you want to know who you are (so what's your configuration):
```
$ gobi whoami
```

If you want to create a command line application:
```
$ gobi cl <APPNAME>
```

If you want to create a Go package:
```
$ gobi pkg <APPNAME>
```

If you want to create a web application:
```
$ gobi web <APPNAME>
```

In all cases `<APPNAME>` can have one or two levels and can't be empty. (Examples: `regexp`, `fmt`, `net/http`, `crypto/md5`)


##TODO
* Better Tests (unit and functional tests)
* Manage configuration (restart config, update fields, etc.)
* Manage projects (delete, date created, date last modified, etc.)
* `go get` projects after created
* Git management (init, add and commit to new project's repo)
* Fallback (undo everything when creation process fails)
* Introduce CI on projects
* Automatic update of gobi
* Create files asynchronously using go routines
* Lots of refactoring needed
* Your suggestion [HERE](https://github.com/fern4lvarez/gobi/issues)


##Contribute!
You all are welcome to take a seat and make a contribution to this repo: reviews, issues, feature suggestions, possible code or functionality enhancements... Everything is appreciated!

**NOTE**: If you work on a fork of this repository, you have to export a `GOBIPATH` environment variable with the `go get` name of your fork, i.e.:

    export GOBIPATH=github/YOUR_NAME/gobi


##License
`gobi` is MIT licensed, see [here](https://github.com/fern4lvarez/gobi/blob/master/LICENSE)
