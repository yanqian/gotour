package main

import ("fmt"
		// "net/http"
		// "io/ioutil"
		"encoding/xml")
var washPostXML = []byte(`
<sitemapindex>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-lifestyle-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-entertainment-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-entertainment-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-goingoutguide-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-goingoutguide-sitemap.xml</loc>
   </sitemap>
</sitemapindex>`)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct{
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	// resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	// bytes, _ := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// var s SitemapIndex
	// xml.Unmarshal(bytes, &s)
	// fmt.Println(s.Locations)
	
	bytes := washPostXML
	var s SitemapIndex
	// fmt.Println(string(bytes))
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
}