package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println("vim-go")
}

func errorStrategies() {
	// multiple strategies for handling error

	// 1. Propagate error
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// 1.1 generate better error message
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		// whole idea for the error message is to be as readable as possible
		// and also to be as consistent as possible
		// throughout the application
		// message strings shouldn't be capitalized and newlines should be avoided
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	// 2. Retry the operation if fails
	// check function for more explanation
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}

	// 3. Basically the same approach as above
	// but more concise
	if err := WaitForServer(url); err != nil {
		log.Fatalf(os.Stderr, "Site is down: %v\n", err)
		// 2019/08/29 11:00:00 Site is down: no such domain: bad.gopl.io
	}

	// 4. Log the error and then continue
	// using log
	if err := Ping(); err != nil {
		log.Printf("ping failed: %v; network disabled", err)
	}
	// printing to standard error stream
	if err := Ping(); err != nil {
		// as this approach doesn't use log, it has newline at the end
		fmt.Fprintf(os.Stderr, "ping failed: %v; network disabled\n", err)
	}

	// 5. --BETTER AVOID--
	// in rare cases, error can be avoided entirely
	dir, err := ioutil.TempDir("", "scratch")
	if err != nil {
		return fmt.Errorf("failed to create temp dir: %v", err)
	}
	// ...using temporary dir...
	os.RemoveAll(dir) // ignore errors; $TMPDIR is creanend periodically
}

// WaitForServer attempts to contact the server of a URL
// It tries for one minute using exponential back-off
// Reports an error if all attemps fail
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
