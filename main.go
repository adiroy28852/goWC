package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func count(r io.Reader) (lines, words, bytes int) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		lines++
		words += len(splitWords(line))
		bytes += len(line) + 1
	}
	return
}

func splitWords(line string) []string {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)
	words := []string{}
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
}

func main() {
	var linesOnly, wordsOnly, charsOnly = false, false, false
	fileNames := []string{}

	if len(os.Args) > 1 {
		if os.Args[1] == "-l" {
			linesOnly = true
			if len(os.Args) > 2 {
				fileNames = os.Args[2:]
			}
		}
		if os.Args[1] == "-c" || os.Args[1] == "m" {
			charsOnly = true
			if len(os.Args) > 2 {
				fileNames = os.Args[2:]
			}
		}
		if os.Args[1] == "-w" {
			wordsOnly = true
			if len(os.Args) > 2 {
				fileNames = os.Args[2:]
			}
		} else {
			fileNames = os.Args[1:]
		}
	}

	if len(fileNames) > 0 {
		for _, fileName := range fileNames {
			file, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				continue
			}
			defer file.Close()
			l, w, b := count(file)
			if linesOnly {
				fmt.Printf("%d %s\n", l, fileName)
			} else if wordsOnly {
				fmt.Printf("%d %s\n", w, fileName)
			} else {
				fmt.Printf("lines %d words %d characters %d %s\n", l, w, b, fileName)
			}
		}
	} else {
		l, w, b := count(os.Stdin)
		fmt.Printf(" %d %d %d\n", l, w, b)
	}
}
