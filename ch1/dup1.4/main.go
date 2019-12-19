// Dup1.4 prints the names of all files in which each duplicated line occ urs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	var filesWithDups []string
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			if hasDups(f) {
				filesWithDups = append(filesWithDups, arg)
			}
			f.Close()
		}
	}
	for _, fileName := range filesWithDups {
		fmt.Printf("%s has duplicated lines.\n", fileName)
	}
}

func hasDups(f *os.File) bool {
	lines := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		lines[input.Text()]++
		for _, count := range lines {
			if count > 1 {
				return true
			}
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
