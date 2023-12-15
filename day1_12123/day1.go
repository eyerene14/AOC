// source: https://gobyexample.com/reading-files
// source: https://go.dev/blog/slices-intro
package main

import (
	"bufio"
	"fmt"
	"strings"
	//"io"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// source: https://go.dev/blog/slices-intro
var digitRegexp = regexp.MustCompile("[0-9]+")
var b []int
func FindDigits(filepath string) int {

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		check(err)
	}
	defer file.Close()

	//fmt.Print(file)
	scanner := bufio.NewScanner(file)

	var total int
	//var m [] int
	//counter := 0
	var digitRegexp = regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{1}]*`)
	// Read each line and search for the first integer
	for scanner.Scan() {
		line := scanner.Text()
		//left := digitRegexp.Find(line)
		all:=digitRegexp.FindAllString(line,-1)
		right_i := strings.Split(all[len(all)-1], "")
		left_i := all[0]
		m,_:= strconv.Atoi(right_i[len(right_i)-1:][0])
		b,_:= strconv.Atoi(left_i[:1])
		total=b+m
		fmt.Printf("%d%d\n",m,b)
	}
		/*m = string(line[int(left_i[0])])
		//var right_i []int
		for i := len(line[left_i[1]:]); i > 0; i-- {
			if right_i := digitRegexp.FindIndex(line[left_i[1]:]); right_i != nil {
				fmt.Println(right_i)
				fmt.Println(string(line[int(right_i[0])]), m)
				break
			}
		}
		//m+=string(line[int(right_i[1])])*/
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return total
}

func main() {

	//fmt.Println(FindDigits("./day1_12123/120123_puzzle.txt"))
	fmt.Println(FindDigits("./day1_12123/1line.txt"))
}
