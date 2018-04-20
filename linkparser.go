package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseHTML(d io.Reader) ([]Link, error) {
	var links []Link
	doc, err := html.Parse(d)
	check(err)
	//ha := false
	//var href string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val, n.FirstChild.Data)
					links = append(links, Link{a.Val, n.FirstChild.Data})
				}
			}
		} else if n.Type == html.ElementNode && n.Data == "/a" {
			fmt.Println("here", n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			fmt.Println(c)
			f(c)
		}
	}
	f(doc)
	return links, nil
}
func main() {

	htmlFile := flag.String("html", "ex2.html", "html to parse")
	flag.Parse()
	htF, err := os.Open(*htmlFile)
	check(err)
	defer htF.Close()
	links, err := parseHTML(htF)
	fmt.Printf("%v", links)
}
