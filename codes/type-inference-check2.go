// filename: type-inference-check2.go
package main

import (
	"fmt"
)

func main() {
  i := 42
  f := 3.142
  g := 0.867 + 0.5i
  fmt.Printf("Type: %T Value: %v\n", i, i)
  fmt.Printf("Type: %T Value: %v\n", f, f)
  fmt.Printf("Type: %T Value: %v\n", g, g)
}
