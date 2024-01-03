package puzzle

import (
	"bufio"
	"fmt"
	//"io"
	"os"
)

var Counter=0

var Lines map[int]string

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filepath string) *bufio.Scanner {

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		Check(err)
	}

	return bufio.NewScanner(file)
}

