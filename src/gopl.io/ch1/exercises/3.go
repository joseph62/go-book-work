package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	timeFunction(BasicEcho, "Basic Echo")
	timeFunction(RangeEcho, "Range Echo")
	timeFunction(JoinEcho, "Join Echo")
}

func timeFunction(f func(), name string) {
	start := time.Now()
	f()
	fmt.Println(name, "took", time.Since(start).Nanoseconds(), "nanoseconds to complete")
}

func BasicEcho() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}

func RangeEcho() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

func JoinEcho() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
