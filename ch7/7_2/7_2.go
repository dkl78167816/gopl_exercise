package main

import (
	"fmt"
	"io"
	"os"
)

// exercise 7.2.
type ByteWriter struct {
	w     io.Writer
	count int64
}

func (c *ByteWriter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	if err != nil {
		return n, err
	}
	c.count += int64(n)
	return n, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	writer := &ByteWriter{w, 0}
	return writer, &writer.count
}

func main() {
	w, c := CountingWriter(os.Stdout)
	fmt.Fprintf(w, "Hello, this world\n")
	fmt.Println(*c)
}
