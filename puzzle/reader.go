package puzzle

import (
	"bufio"
	"fmt"
	"regexp"
	//"io"
	"os"
)

var Counter = 0

var Lines map[int]string


func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filepath string) map[int]string {

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		Check(err)
	}

	scanner := bufio.NewScanner(file)
	Lines= make(map[int]string)
	//var digitRegexp = regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	//var digitRegexp = regexp.MustCompile(`[^\d.]`)
	// [&%$*#]Read each line and search for the first integer
	for scanner.Scan() {
		line := scanner.Text()
		Lines[Counter] = line
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
