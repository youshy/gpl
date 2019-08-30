package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func main() {
	returnFetch()
}

// execution functions
func returnTitle() {
	for _, arg := range os.Args[1:] {
		if err := title(arg); err != nil {
			fmt.Fprintf(os.Stderr, "title: %v\n", err)
		}
	}
}

func returnTrace() {
	bigSlowOperation()
}

func returnFetch() {
	for _, url := range os.Args[1:] {
		local, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}
		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
	}
}

// /execution functions

//
// below block returns title of a webpage
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// defer literally blocks further function execution before the
	// whole resource is available - regardless of it's status
	// meaning - that it can return a value normally (requested or error)
	// or abnormally - by panicking.
	defer resp.Body.Close()
	// why using defer? Because in this case, the connection
	// won't be closed afterwards - it's basically a memory leak

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)

	return nil
}

//

//
// trace simulates how defer works with big operations
func bigSlowOperation() {
	// basically how it works:
	// in this example defer works "on entry" and "on exit"
	// trace will call defer at the beginning
	// and then, by the power of () (extra parentheses) it will call "on exit"
	defer trace("big slow operation")()
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	// returns an executable function
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

//

//
// Fetch downloads the URL and returns the
// name and lenght of the local file
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		// Close file, but prefer error from Copy, if any.
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)
	return local, n, err
}

//
