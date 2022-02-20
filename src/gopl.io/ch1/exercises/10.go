/*
	I tested Amazon.com's book search and found roughly the same response time give or take a few milliseconds.
	The number of bytes did fluctuate a bit (a few hundred difference between runs for a 75k sized page).
	The differences were down to some class name differences, whitespace, and errant javascript blurbs.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	url = addProtocol(url)
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	bodyData, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	nbytes := len(bodyData)
	body := string(bodyData)

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n%s\n\n", secs, nbytes, url, body)
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
