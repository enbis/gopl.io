package main

import (
	"bytes"
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
	fmt.Fprintf(out, "%s", printmap())
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		counts[n.Data]++
		//fmt.Fprintf(out, "%s", counts)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func printmap() string {

	b := new(bytes.Buffer)
	for key, value := range counts {
		fmt.Fprintf(b, "%s=%d&&", key, value)
	}

	return b.String()
}
