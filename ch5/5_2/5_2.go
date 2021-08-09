// prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var counts = make(map[string]int)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
	for k, v := range counts {
		if v > 1 {
			fmt.Printf("<%s> : %d\n", k, v)
		}
	}
}

// exercise 5-2
func visit(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		counts[n.Data] += 1
	}
	visit(n.FirstChild) // Depth first
	visit(n.NextSibling)
}
