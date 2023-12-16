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
//var digitRegexp = regexp.MustCompile("[0-9]+")
var total Total

type Total struct {
	lineNo  int
	left    string
	right   string
	joined  int
	running int
}

func FindDigits(filepath string) Total {

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		check(err)
	}
	defer file.Close()

	//fmt.Print(file)
	scanner := bufio.NewScanner(file)

	var digitRegexp = regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	counter := 0
	found := Total{}
	// Read each line and search for the first integer
	for scanner.Scan() {
		line := scanner.Text()
		//var wn = regexp.MustCompile(`(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)`)
		fmt.Print(line, " - ")
		replacer := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
		line = replacer.Replace(line)
		fmt.Println(line)
		all := digitRegexp.FindAllString(line, -1)
		right_i := strings.Split(all[len(all)-1], "")
		//found.right,_= strconv.Atoi(right_i[len(right_i)-1:][0])
		//found.left,_= strconv.Atoi(left_i)
		found.right = right_i[len(right_i)-1:][0]
		found.left = all[0][:1]
		found.joined, _ = strconv.Atoi(found.left + found.right)
		found.running += found.joined
		fmt.Printf("line: %d, left: %d, right:%d, joined:%d, total:%d\n", found.lineNo, found.left, found.right, found.joined, found.running)
		found.lineNo = counter
		counter++
	}
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return found
}

func main() {
	line := FindDigits("./day1_12123/day1_data.txt")
	//line:=FindDigits("./day1_12123/1line.txt")
	total = line
	//fmt.Println(total)
}

//>6pm - part 1 total:54916
//6:11pm - part 1 total:54702
