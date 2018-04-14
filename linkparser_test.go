package linkparser

import (
	"os"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func TestLinkParserEx1(t *testing.T) {
	// Create candidClient with cookieJAR
	ht, err := os.Open("ex1.html")
	check(err)
	defer ht.Close()
	var links []Link
	links, err = parseHTML(ht)
	if len(links) != 1 {
		t.Error("We should only see 1 link")
	}

}
