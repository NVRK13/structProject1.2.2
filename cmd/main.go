package main

import (
	"fmt"

	Utils "structproject1.2.2/pkg/utils"
)

func main() {
	s := "Hello"
	n := Utils.CountChars(s)
	fmt.Printf("Bytes in %q: %d\n", s, n)

}
