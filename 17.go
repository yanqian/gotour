package main

import ("fmt"
		"net/http"
		"html/template"
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
type NewsAggPage struct {
	Title string
	News map[string]NewsMap
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey there!</h1>")
	fmt.Fprintf(w, "<p>go is fast </p>")
	fmt.Fprintf(w, "<p>and simple...... </p>")
	fmt.Fprintf(w, "you %s even add %s  ", "can", "<strong>variable</strong>")
}

func aggPageHandler(w http.ResponseWriter, r *http.Request) {
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
	}
	p := NewsAggPage{Title: "Balabla", News: news_map}
	t, _ := template.ParseFiles("aggnewstemplating.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", aggPageHandler)
	http.ListenAndServe(":8000", nil)
}