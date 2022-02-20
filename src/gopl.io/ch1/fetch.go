package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	urls := os.Args[1:]

	for _, url := range urls {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintln(os.Stderr, "fetch:", err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(resp.Body)

		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: Error reading url '%s': %v\n", url, err)
			os.Exit(1)
		}

		fmt.Print(string(body))
	}
}