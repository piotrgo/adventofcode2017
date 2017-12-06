package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func main() {
	//open the input file
	f, err := os.Open("input")
	check(err)
	//close the file when we finish
	defer f.Close()

	//read lines one by one and put them in an array
	var lines []string
	scanner := bufio.NewScanner(f)
    	for scanner.Scan() {
		lines = append(lines, scanner.Text())
    	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	//decompose into individual numbers to then pick max/min
	for i := range lines{
		words := strings.Fields(lines[1])
	}

}
