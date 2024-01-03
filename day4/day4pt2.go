package main

import (
	"adventofcode/user/hello/puzzle"
	"bufio"
	"fmt"
	"regexp"

	//"strconv"
	//"strings"
	"golang.org/x/exp/slices"
)

type Winners struct{
	line int
	count int
}

type Pc struct{
	parent int
	children []int
}
var pc Pc
var tracker=make(map[int]int)
var counter = 0
/*
func AddCopies(tracker map[int]int, max int) map[int]int{
	counter:=puzzle.Counter
	for ; counter<max+1; counter++{
		if _,ok:=tracker[counter]; !ok{
			tracker[counter+1]++
		}
		tracker[counter]++
	}
	return tracker
}*/

func FindWinners(wns []string, ons []string) Winners {
	countWins,countIters:=0,0
	for countWins == 0 || countIters == 0{
		for _,v:= range wns{
			countIters++
			if slices.Contains(ons,v){	
				countWins++
				tracker[puzzle.Counter+countWins]++
				fmt.Print(countIters)	
			}
		}
		break
	}
	
	fmt.Println(tracker)
	return Winners{puzzle.Counter,countWins}
}

func ReadLines(scanner *bufio.Scanner) {

	var digitRegexp = regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	// Read each line and search for the first integer
	for scanner.Scan() {
		puzzle.Counter++
		line := scanner.Text()
		d:=digitRegexp.FindAllString(line,-1)
		l:=6
		//l=11
		leftn:=d[1:l]
		rightn:=d[l:]
		winners:=&Winners{}
		*winners=FindWinners(leftn,rightn)
	}
	
	//return total
}

func main() {
	//line:=FindDigits(puzzle.ReadFile("./day4/day4data.txt"))
	ReadLines(puzzle.ReadFile("./day4/day4test.txt"))
	fmt.Println(pc)
}