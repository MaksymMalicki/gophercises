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

func ReadExampleHTML(path string) io.Reader {
	f, err := os.ReadFile(path)
	if err != nil {
		panic("Could't open the file")
	}
	htmlString := strings.TrimSpace(string(f))
	return strings.NewReader(htmlString)
}

func ParseHTML(buffer io.Reader) []Link {
	doc, err := html.Parse(buffer)
	if err != nil {
		panic("Could't parse the HTML")
	}
	links := []Link{}
	findLinkNodes(doc, &links)
	return links
}

func findLinkNodes(n *html.Node, links *[]Link) []Link {
	if n.Type == html.ElementNode && n.Data == "a" {
		*links = append(*links, createLinkFromNode(n))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findLinkNodes(c, links)
	}
	return nil
}

func createLinkFromNode(n *html.Node) Link {
	var href string = ""
	var result string = ""
	for _, a := range n.Attr {
		if a.Key == "href" {
			href = a.Val
			break
		}
	}
	expandLinkNodeChildren(n, &result)
	return Link{text: strings.Join(strings.Fields(result), " "), href: href}
}

func expandLinkNodeChildren(n *html.Node, result *string) string {
	if n.Type == html.TextNode {
		*result += n.Data + " "
		return *result
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		expandLinkNodeChildren(c, result)
	}
	return *result
}
