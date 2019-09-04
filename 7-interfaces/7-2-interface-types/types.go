package main

import (
	"fmt"
	"io"
)

func main() {
	fmt.Println("vim-go")
}

// ex7.4
// simple version of NewReader
type stringReader struct {
	s string
}

func (r *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}

func NewReader(s string) io.Reader {
	return &stringReader{s}
}

// ex7.5
// LimitReader implements io.Reader and number of bytes
// and returns another reater
type limitReader struct {
	r        io.Reader
	n, limit int
}

func (r *limitReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p[:r.limit])
	r.n += n
	if r.n >= r.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, limit int) io.Reader {
	return &limitReader{r: r, limit: limit}
}
