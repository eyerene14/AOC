package main

import (
	"adventofcode/user/hello/puzzle"
	"bufio"
	"fmt"
	//"regexp"
	//"io"
	"os"
)

var Counter = 0

var Lines map[int]string

func FindDigits(scanner *bufio.Scanner) int {	

	//var digitRegexp = regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	counter := 0
	// Read each line and search for the first integer
	for scanner.Scan() {
		line := scanner.Text()
		counter++
		fmt.Println(line)
	}
	return counter
}

func main() {
	line:=FindDigits(puzzle.ReadFile("./day3/day3test.txt"))
	total:=&line
	fmt.Println(*total)
}