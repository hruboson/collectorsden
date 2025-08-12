package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: stripheader <input> <output>")
		os.Exit(1)
	}

	inFile := os.Args[1]
	outFile := os.Args[2]

	in, err := os.Open(inFile)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	scanner := bufio.NewScanner(in)

	firstLine := true
	for scanner.Scan() {
		if firstLine {
			firstLine = false
			continue // skip the header line
		}
		fmt.Fprintln(out, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
