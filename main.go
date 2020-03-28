package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	url = "https://golang.org"
)
func main() {
	headerTitles, err := GetHeaderTitles(url, "h1")
	
	if err != nil {
		log.Println(err)
	}

	links, err := GetlinksOnSpecificURL(url)

	if err != nil {
		log.Println(err)
	}

	//print h1 titles
	fmt.Println("Header titles")
	fmt.Println(headerTitles)
	
	//print links
	fmt.Println("Links")
	fmt.Println(links)
}

//GetHeaderTitles from website url
func GetHeaderTitles(url string, tag string) (string, error){
//Get HTML from DOM
res, err := http.Get(url)

if err != nil {
	return "", err
}

//Convert HTML into goquery
doc, err := goquery.NewDocumentFromReader(res.Body)
if err != nil {
	return "", err
}

//Save each h1 as list
titles := ""
doc.Find(tag).Each(func (i int, s *goquery.Selection)  {
	titles += "- " + s.Text() + "\n"
})

return titles, nil

}

//GetlinksOnSpecificURL on specific url
func GetlinksOnSpecificURL(url string) (string, error){
res, err := http.Get(url)

if err != nil {
	return "", err
}

//Convert HTML into goquery
doc, err := goquery.NewDocumentFromReader(res.Body)
if err != nil {
	return "", err
}

//Save each link as list
links := ""
doc.Find("a[href]").Each(func (i int, s *goquery.Selection)  {
	links += "- " + s.Text() + "\n"
})

return links, nil

}