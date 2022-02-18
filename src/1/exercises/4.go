package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, count_by_files := range counts {
		var total int
		var files, sep string
		for file, count := range count_by_files {
			total += count
			files += sep + file
			sep = ", "
		}

		if total > 1 {
			fmt.Printf("%d references\t from file(s) %s\n%s\n\n", total, files, line)
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	name := f.Name()

	for input.Scan() {
		text := input.Text()

		count_by_file, exists := counts[text]

		if !exists {
			count_by_file = make(map[string]int)
			counts[text] = count_by_file
		}

		count_by_file[name]++
	}
}
