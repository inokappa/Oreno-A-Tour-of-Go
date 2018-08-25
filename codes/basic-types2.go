// filename: basic-types2.go
package main

import (
	"fmt"
)

var (
	Foo    string     = "Foo"
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", Foo, Foo)
}
