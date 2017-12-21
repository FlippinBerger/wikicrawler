package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	searchTerm := flag.String("term", "", "Term to search for on wiki. Double quoted if whitespace is involved. Required.")

	flag.Parse()

	if len(os.Args) != 3 {
		fmt.Println("Usage of ./wikicrawl:")
		flag.PrintDefaults()
		return
	}

	fmt.Println("The current search terms is", *searchTerm)

	crawl("http://www.google.com")

	//now that we have our search term, we need to sanitize it

	//then we need to seed a crawler with the wiki page url to start with
	//and it should take care of the rest from there for us
}
