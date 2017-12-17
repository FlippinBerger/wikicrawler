package main

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

}

// need to also have a scraper that takes the data from the
// http request and finds all the links within, and some
// data about the page i.e. title

//data is a string for now but will be whatever the http response
//gives us
func scrape(data string) {

}
