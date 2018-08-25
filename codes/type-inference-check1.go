// filename: type-inference-check1.go
package main

import (
	"fmt"
)

func main() {
  var i int = 100
  j := i
  fmt.Printf("Type: %T Value: %v\n", j, j)
}
