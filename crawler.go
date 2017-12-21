package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//going to act as a concurrent wiki webcrawler for specific info
//for now creates a file with a list of n wiki articles relating
//to your search term

// going to need a set of visited urls
var visited = make(map[string]bool)

//this function will start crawling from url
//makes the http request and kicks off the scraper
//to grab the data
//Will be concurrent
func crawl(url string) {
	//can only make a request 1/s so need to gate this call
	bytes := makeRequest(url)

	fmt.Println(string(bytes))
	//at this point we have made the request and just need to get the urls from the response body
	//go to town. glhf kid
	//urls := scrape(bytes)

}

func makeRequest(url string) []byte {
	response, err := http.Get(url)

	if err != nil {
		//do the thing here
	}

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		//do the thing here
	}

	return bytes
}

// need to also have a scraper that takes the data from the
// http request and finds all the links within, and some
// data about the page i.e. title

//data is a string for now but will be whatever the http response
//gives us
func scrape(data []byte) []string {
	strs := make([]string, 5)
	return strs
}
