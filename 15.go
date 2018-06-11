package main

import ("fmt"
		"net/http"
		"io/ioutil"
		"encoding/xml")

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct{
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct{
	Keyword string 
	Location string 
}

func main() {
	var s SitemapIndex
	var n News

	news_map := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	
	xml.Unmarshal(bytes, &s)
	for _, Location := range s.Locations{
		resp_loc, _ := http.Get(Location)
		bytes_news, _ := ioutil.ReadAll(resp_loc.Body)
		resp_loc.Body.Close()
		xml.Unmarshal(bytes_news, &n)
		for index, _ := range n.Titles{
			news_map[n.Titles[index]] = NewsMap{n.Keywords[index], n.Locations[index]}
		}

		for k,v := range news_map{
			fmt.Print("\n\n\n------ Title --------", k)
			fmt.Print("\n------ Keyword --------", v.Keyword)
			fmt.Print("\n------ Location --------", v.Location)
		}
		// fmt.Printf("Index: %d, Value:  %s \n",index,Location)
	}
	
	
}