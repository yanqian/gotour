package main

import ("fmt"
		"net/http"
		"html/template"
		"io/ioutil"
		"encoding/xml"
		"sync")

var wg sync.WaitGroup
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

func NewsRoutine(c chan News, Location string) {
	defer wg.Done()
	var n News
	resp_loc, _ := http.Get(Location)
	bytes_news, _ := ioutil.ReadAll(resp_loc.Body)
	resp_loc.Body.Close()
	xml.Unmarshal(bytes_news, &n)
	c <- n
}

func aggPageHandler(w http.ResponseWriter, r *http.Request) {
	var s SitemapIndex
	news_map := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)

	// use it as early as possible
	c := make(chan News, 30)
	for _, Location := range s.Locations{
		wg.Add(1)
		go NewsRoutine(c, Location)
	}
	wg.Wait()
	close(c)

	for ele := range c {
		for index, _ := range ele.Titles{
			news_map[ele.Titles[index]] = NewsMap{ele.Keywords[index], ele.Locations[index]}
		}
	}
	
	p := NewsAggPage{Title: "New App with go routine!", News: news_map}
	t, _ := template.ParseFiles("aggnewstemplating.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", aggPageHandler)
	http.ListenAndServe(":8000", nil)
}