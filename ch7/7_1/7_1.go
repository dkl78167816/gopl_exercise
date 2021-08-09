// Discard the written bytes and add the length into ByteCounter
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// exercise 7.1, counter to words and lines
type WordCounter int
type LineCounter int

// manual code
func (c *WordCounter) Write(s []byte) (int, error) {
	isWord := false // state of previous token
	var count = 0
	for _, word := range s {
		if word != ' ' && !isWord {
			isWord = true
			count += 1
		} else if word == ' ' {
			isWord = false
		}
	}
	*c += WordCounter(count)
	return count, nil
}

// using bufio.ScanWords
func (c *WordCounter) Scan(s []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(s))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += WordCounter(count)
	return count, nil
}

// using bufio.ScanWords
func (c *LineCounter) Write(s []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(s))
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += LineCounter(count)
	return count, nil
}

func main() {
	fmt.Println("----------count words----------")
	var c WordCounter
	var name = "Dolly"
	fmt.Fprintf(&c, "   hello, %s", name)
	c.Write([]byte("hello    this world   "))
	fmt.Println(c) // 5
	c = 0          // reset the counter
	c.Scan([]byte(fmt.Sprintf("   hello, %s", name)))
	c.Scan([]byte("hello    this world   "))
	fmt.Println(c)

	fmt.Println("----------count lines----------")
	var l LineCounter
	var data = ` Hello

    this dummy but
    beautiful world`

	fmt.Fprintf(&l, "%s", data)
	fmt.Println(l)
}
