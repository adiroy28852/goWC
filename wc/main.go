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

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]

		switch arg {
		case "-l":
			linesOnly = true
		case "-w":
			wordsOnly = true
		case "-c":
			charsOnly = true
		case "-m":
			charsOnly = true 
		default:
			fileNames = append(fileNames, arg)
		}
	}

	if !linesOnly && !wordsOnly && !charsOnly {
		linesOnly = true
		wordsOnly = true
		charsOnly = true
	}

	if len(fileNames) == 0 {
		l, w, b := count(os.Stdin)
		if linesOnly && wordsOnly && charsOnly {
			fmt.Printf(" %d %d %d\n", l, w, b)
		} else {
			if linesOnly {
				fmt.Printf("%d\n", l)
			}
			if wordsOnly {
				fmt.Printf("%d\n", w)
			}
			if charsOnly {
				fmt.Printf("%d\n", b)
			}
		}
		return
	}

	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}
		defer file.Close()

		l, w, b := count(file)

		var outputParts []string
		if linesOnly {
			outputParts = append(outputParts, fmt.Sprintf("%d", l))
		}
		if wordsOnly {
			outputParts = append(outputParts, fmt.Sprintf("%d", w))
		}
		if charsOnly {
			outputParts = append(outputParts, fmt.Sprintf("%d", b))
		}

		fmt.Printf("%s %s\n", strings.Join(outputParts, " "), fileName)
	}
}
