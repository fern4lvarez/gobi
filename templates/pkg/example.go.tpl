package main

import (
  "fmt"
  "os"
  "{{.Host}}/{{.UserId}}/{{.Name}}"
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