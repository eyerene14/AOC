package main

import (
	"bufio"
	"fmt"
	"regexp"
	//"io"
	"os"
)

var Counter = 0

var Lines map[int][][]int

type Vertex struct{
	loc [][]int
	e Edge
}

type Edge struct{
	coord [][]int
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filepath string) map[int][][]int {

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		Check(err)
	}

	scanner := bufio.NewScanner(file)
	Lines = make(map[int][][]int)
	//var digitRegexp = regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	var digitRegexp = regexp.MustCompile(`[^\d.]`)
	// [&%$*#]Read each line and search for the first integer
	for scanner.Scan() {
		line := scanner.Text()
		Lines[Counter] = digitRegexp.FindAllStringIndex(line,-1)
		fmt.Println(Lines)
		Counter++
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	defer file.Close()

	return Lines
}

func AddEdge(l map[int][][]int){

	for i:=0; i < len(Lines); i++{
		col:=l[0][0]
		for row :=0; row < 3; row++{
			e:= Edge{}
			col++
			fmt.Println(e)
		}
	}
}

func main(){
	Lines = ReadFile("/Users/ijaramil/notes/advent_of_code_23/day3/day3test.txt")
	AddEdge(Lines)
}
