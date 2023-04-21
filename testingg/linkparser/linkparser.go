package linkparser

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseHtml(r io.Reader) ([]Link, error) {
	doc, err2 := html.Parse(r)
	if err2 != nil {
		return nil, fmt.Errorf("failed to parse HTML: %s", err2)
	}
	return parseLinks(doc), nil
}

func parseLinks(n *html.Node) []Link {
	links := make([]Link, 0)
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				url := a.Val
				text := parseText(n)
				l := Link{
					Href: url,
					Text: text,
				}
				links = append(links, l)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, parseLinks(c)...)
	}
	return links
}

func parseText(n *html.Node) string {
	str := ""
	if n.Type == html.TextNode {
		str = strings.TrimSpace(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parsedStr := parseText(c)
		if str == "" || parsedStr == "" {
			str += parsedStr
		} else {
			str += " " + parsedStr
		}
	}
	return strings.ReplaceAll(str, "\n", "")
}
