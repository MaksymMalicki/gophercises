package htmllinkparser

import (
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	href string
	text string
}

func ReadExampleHTML() io.Reader {
	f, err := os.ReadFile("./examples/ex1.html")
	if err != nil {
		panic("Could't open the file")
	}
	htmlString := strings.TrimSpace(string(f))
	print(htmlString)
	return strings.NewReader(htmlString)
}

func ParseHTML(buffer io.Reader) []Link {
	doc, err := html.Parse(buffer)
	if err != nil {
		panic("Could't parse the HTML")
	}
	links := []Link{}
	traverseHTMLTree(doc, &links)
	for _, l := range links {
		print(l.href, " ", l.text)
	}
	return links
}

func traverseHTMLTree(n *html.Node, links *[]Link) []Link {
	if n.Type == html.ElementNode && n.Data == "a" {
		*links = append(*links, retrieveATagData(n))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseHTMLTree(c, links)
	}
	return nil
}

func retrieveATagData(n *html.Node) Link {
	var href string = ""
	var result string = ""
	for _, a := range n.Attr {
		if a.Key == "href" {
			href = a.Val
		}
	}
	expandATagChildren(n, &result)
	return Link{text: strings.TrimSpace(result), href: href}
}

func expandATagChildren(n *html.Node, result *string) string {
	print("type: ", n.Type, " data: ", n.Data, "\n")
	if n.Type == html.TextNode {
		*result += n.Data + " "
		return *result
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		expandATagChildren(c, result)
	}
	return *result
}
