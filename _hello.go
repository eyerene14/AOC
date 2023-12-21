// SOURCE: https://go.dev/doc/code
// USAGE: 1. go install adventofcode/user/hello 2. hello
package main

import (
	"adventofcode/user/hello/morestrings"
	"fmt"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG, olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
// OUTPUT: Hello ,Go!