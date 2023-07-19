package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

type Link struct {
	Href string
	Text string
}
type Result struct {
	Links []Link
}

func main() {
	//Vars
	result := Result{}
	//1- Read HTML
	resInByte, _ := os.Open("ex1.html")

	//2- Parse HTML to Tree

	htmlTree, err := html.Parse(resInByte)
	if err != nil {
		return
	}

	//3- Traverse Tree
	traverserAndMutate(htmlTree, &result)
	fmt.Printf("%v", result.Links)
}

func traverserAndMutate(node *html.Node, result *Result) {
	if node.Type == html.ElementNode && node.Data == "a" {
		link := Link{}
		link.Text = extractText(node)
		for _, attr := range node.Attr {
			if attr.Key == "href" {

				link.Href = attr.Val
			}
		}
		result.Links = append(result.Links, link)
	}
	for i := node.FirstChild; i != nil; i = i.NextSibling {
		traverserAndMutate(i, result)
	}
}
func extractText(n *html.Node) string {
	var text string
	if n.Type != html.ElementNode && n.Data != "a" && n.Type != html.CommentNode {
		text = n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c)
	}
	return strings.Trim(text, "\n")
}
