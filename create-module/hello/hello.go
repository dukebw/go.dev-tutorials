package main

import (
  "fmt"
  "github.com/dukebw/go.dev-tutorials/create-module/greetings"
)

func main() {
  message := greetings.Hello("Gladys")
  fmt.Println(message)
}
