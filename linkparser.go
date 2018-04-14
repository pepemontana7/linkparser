package linkparser

import (
	"flag"
)

type Link struct {
	Href string
	Text string
}

func main() {

	htmlFile := flag.String("html", "ex1.html", "html to parse")
	flag.Parse()

}
