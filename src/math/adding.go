package main

import "fmt"

func add(x, y int) int {
  fmt.Println("hue")
  return x + y
}

func hue(x string, y int) (int, string) {
  return y, x
}

func main(){

  fmt.Println("adding func", add(42, 13))

  a, b := hue("hello, you're number", 1)
  fmt.Println(a, b)

  fmt.Println(split(17))
}

func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}
