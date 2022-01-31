package main

import (
	"encoding/csv"
	"fmt"

	// "bufio"
	// "io"

	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFromFile(filename string) {
	f, err := os.Open(filename + ".csv")
	check(err)

	lines, err := csv.NewReader(f).ReadAll()
	check(err)

	fmt.Println(lines)
}

func main() {
	readFromFile("test")
}
