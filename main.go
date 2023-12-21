package main

import (
	"adventofcode/user/hello/puzzle"
	"fmt"
)

func main() {
	c := make(chan string)
	go puzzle.ReadFile("/Users/ijaramil/notes/advent_of_code_23/day3/day3test.txt",c)
	//puzzle.GetLines()
	fmt.Print(<-c)
}

// OUTPUT: Hello ,Go!
