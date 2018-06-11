package main

import ("fmt"
		"net/http"
		"io/ioutil"
		"encoding/xml")

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct{
	Title []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Location []string `xml:"url>loc"`
}

func main() {
	var s SitemapIndex
	var n News
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	
	xml.Unmarshal(bytes, &s)
	for _, Location := range s.Locations{
		resp_loc, _ := http.Get(Location)
		bytes_news, _ := ioutil.ReadAll(resp_loc.Body)
		resp_loc.Body.Close()
		xml.Unmarshal(bytes_news, &n)
		// fmt.Printf("Index: %d, Value:  %s \n",index,Location)
	}
	fmt.Println("Nothing????")
	for index, title := range n.Title{
		fmt.Printf("Index: %d, Value:  %s \n",index,title)
	}
	// fmt.Println(n.Title)
	
}