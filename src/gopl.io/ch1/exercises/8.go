package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := os.Args[1:]

	for _, url := range urls {
		url = addProtocol(url)

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintln(os.Stderr, "fetch:", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)

		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: Error reading url '%s': %v\n", url, err)
			os.Exit(1)
		}
	}
}

func addProtocol(url string) string {
	if strings.HasPrefix(url, "https://") {
		return url
	} else if strings.HasPrefix(url, "http://") {
		return url
	} else {
		return "https://" + url
	}
}
