package main

import (
	"os"
	"reflect"
	"testing"
)

func TestLinkParserEx1(t *testing.T) {
	// Create candidClient with cookieJAR
	htF, err := os.Open("ex1.html")
	check(err)
	defer htF.Close()
	link := Link{Href: "/other-page",
		Text: "A link to another page"}
	var links []Link
	links, err = parseHTML(htF)

	if len(links) != 1 {
		t.Fatal("We should only see 1 link")
	}
	if link != links[0] {
		t.Errorf("link do not match got %v expected %v", links[0], link)
	}
}

func TestLinkParserEx2(t *testing.T) {
	// Create candidClient with cookieJAR
	htF, err := os.Open("ex2.html")
	check(err)
	defer htF.Close()
	l1 := Link{Href: "https://www.twitter.com/joncalhoun",
		Text: "Check me out on twitter"}
	l2 := Link{Href: "https://github.com/gophercises",
		Text: "Gophercises is on <strong>Github</strong>!"}
	exLinks := []Link{l1, l2}
	var links []Link
	links, err = parseHTML(htF)

	if len(links) != 2 {
		t.Fatal("We should only see 2 link")
	}
	if !reflect.DeepEqual(exLinks, links) {
		t.Errorf("link do not match got %v expected %v", exLinks, links)
	}
}
