// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, text := range visit(nil, doc) {
		fmt.Println(text)
	}
}

// task 5-3
func visit(texts []string, n *html.Node) []string {
	if n == nil {
		return texts
	}
	if n.Type == html.TextNode && n.Data == "text" {
		for _, a := range n.Attr {
			texts = append(texts, a.Val)
		}
	}
	texts = visit(texts, n.FirstChild) // Depth first
	texts = visit(texts, n.NextSibling)

	return texts
}
