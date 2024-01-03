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

var length=6
var total=0
var winners=make([]bool,5)
var Lines map[int][]bool
type Totals struct{
	sum int
	run_sum int
}
var found map[int]Totals

func Double(count,sum int)int{
	if count > 2{
		return sum + sum
	}
	return sum + 1
}

func CheckWinners(wns []string, ons []string) []bool {
	winners=make([]bool,len(wns))
	found=make(map[int]Totals)
	count,sum:=0,0
	for j,v:= range wns{
		//nums:=Winners{v,strings.Contains(pool,v)}
		winners[j]=slices.Contains(ons,v)
		//fmt.Print(" ", v)
		if slices.Contains(ons,v){
			count++		
			sum=Double(count,sum)
		}	
		
	}

	
	total+=sum	
	found[puzzle.Counter]=Totals{count,sum}
	fmt.Printf("found= %v total=%d",found, total)	
	return winners
}

func FindDigits(scanner *bufio.Scanner) int {

	var digitRegexp = regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	// Read each line and search for the first integer
	for scanner.Scan() {
		puzzle.Counter++
		line := scanner.Text()
		d:=digitRegexp.FindAllString(line,-1)
		//length=6
		length=11
		leftn:=d[1:length]
		rightn:=d[length:]
		//fmt.Println(puzzle.Counter,leftn,rightn)
		Lines=make(map[int][]bool)
		Lines[puzzle.Counter]=CheckWinners(leftn,rightn)
		fmt.Println(Lines)
	}
	
	return total
}

func nonmain() {
	line:=FindDigits(puzzle.ReadFile("./day4/day4data.txt"))
	//line:=FindDigits(puzzle.ReadFile("./day4/day4test2.txt"))
	total:=&line
	fmt.Println(*total)
}