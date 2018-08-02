package main

import (
  "fmt"
  "xlsx"
)

func main() {
  inflation()
  wageGrowth()
  advisoryBoard()

  excelFileName := "/home/go/go/src/programs/yo.xlsx"
    xlFile, err := xlsx.OpenFile(excelFileName)
    if err != nil {
        fmt.Println("lel")
    }
  }

    func (c *A) String() (string, error) {
	     return c.FormattedValue()
    }


func inflation() {
}
func wageGrowth() {
}
func advisoryBoard() {
}
