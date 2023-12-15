package main

import (
	//"bufio"
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func readFile(path string){
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
}

func main() {
	//fmt.Println("Hello, world.")
	//file, err := os.ReadFile("./day1_12123/120123_puzzle.txt")
	//fmt.Print(string(file))
	file, err := os.Open("./day1_12123/120123_puzzle.txt")
	//b1 := make([]byte,5)
	//n1, err := file.Read(b1)
	//check(err)
	//fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2,err := file.Seek(6,0) // reads the index 5 char in first line
	check(err)
	b2 := make([]byte,4)
	//n2,err := file.Read(b2)
	n3, err := io.ReadAtLeast(file, b2, 2)
	check(err)
	//fmt.Printf("%d bytes @ %d: ", n2, o2)
	//fmt.Printf("%v\n", string(b2[:n2])) // returns 4 chars from index 5-9 char in first line: \ans
	fmt.Printf("%d bytes @ %d: %s\n", n3,o2,string(b2))

	_, err = file.Seek(0, 0)
    check(err)

	r := bufio.NewReader(file)
	b4,err := r.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	file.Close()

	file.
}
