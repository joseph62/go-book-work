package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, file := range os.Args[1:] {
		data, err := ioutil.ReadFile(file)

		if err != nil {
			fmt.Fprintln(os.Stderr, "dup3:", err)
			continue
		}

		lines := strings.Split(string(data), "\n")

		for _, line := range lines {
			counts[line]++
		}
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
