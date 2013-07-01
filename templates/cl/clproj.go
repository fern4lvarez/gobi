package main

import (
	"flag"
	"fmt"
)

func main() {
	var f string

	flag.StringVar(&f, "gobi", "is awesome", "Just prints.")
	flag.Parse()

	fmt.Println(f)
}
