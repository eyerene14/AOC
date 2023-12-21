package main

import (
	"adventofcode/user/hello/readdata"
	"fmt"
	//"io"
	//"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Totals struct {
	line int
	game Game
}

type Game struct {
	line int
	num  int
}

var rules = &map[string]int{"red": 12,
	"green": 13,
	"blue":  14,
}

func FindColor(file string) int {
	//fmt.Print(file)
	scanner := bufio.NewScanner(file)

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

func FindDigits() {

}

func main() {
	c := make(chan int)
	go readdata.ReadFileLines("./day1_12123/day1_data.txt", c)
	lines := <- c
	fmt.Println(lines)
}
