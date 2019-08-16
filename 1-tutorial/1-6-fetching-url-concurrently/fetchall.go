// Fetching all URLs in parallel and reports their times and sizes
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // starts a goroutine
		// goroutine is a concurrent function execution
		// this loop starts one for each argument and runs them asynchronously
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // gets the body of the response and then throws it away (Discard)
	resp.Body.Close()                                 // don't leak resources
	if err != nil {
		// the <- is an expression
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
