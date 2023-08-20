package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func readHtmlFromFile(fileName string) (string, error) {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(bs), nil
}

type Link struct {
	Href string
	Text string
}

// grabText - grab the text from HTML node and return it
func grabText(n *html.Node) string {
	var sb strings.Builder

	if n.Type == html.TextNode {
		s := n.Data
		sb.WriteString(s)
	}

	return sb.String()
}

func main() {
	// check if there is a html file provided as args
	if len(os.Args) == 1 {
		log.Fatal("You need to provide a valid html file")
	}

	// read from html file
	text, err := readHtmlFromFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// parse the HTML
	doc, err := html.Parse(strings.NewReader(text))
	if err != nil {
		log.Fatal(err)
	}

	// go to all tags, check if it's there an a and create a Link struct with href and text within the tag and add it to links
	var links []Link
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for a := range n.Attr {
				if n.Attr[a].Key == "href" {
					links = append(links, Link{
						Href: n.Attr[a].Val,
						Text: strings.TrimSpace(grabText(n.FirstChild)),
					})
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	// print all the links
	for i := range links {
		fmt.Println(links[i])
	}
}
