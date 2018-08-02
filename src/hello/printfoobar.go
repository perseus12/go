package main

import (
  "fmt"
)

var (
   bar string = "bar"
)

func main() {
  return fmt.Sprintf("foo: %s", bar)
}
