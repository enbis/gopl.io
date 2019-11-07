package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

var counts = make(map[string]int)
var out io.Writer = os.Stdout

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Printf("outline: %s\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "script" {
		stack = append(stack, n.Data) // push tag
		counts[n.Data]++
		fmt.Fprintln(out, stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
