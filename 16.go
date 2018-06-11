package main

import ("fmt"
		"net/http"
		"html/template")

type NewsAggPage struct {
	Title string
	News string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey there!</h1>")
	fmt.Fprintf(w, "<p>go is fast </p>")
	fmt.Fprintf(w, "<p>and simple...... </p>")

	fmt.Fprintf(w, "you %s even add %s  ", "can", "<strong>variable</strong>")
}

func aggPageHandler(w http.ResponseWriter, r *http.Request) {
	n := NewsAggPage{Title: "Balabla", News: "news area and you should fill this"}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, n)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", aggPageHandler)
	http.ListenAndServe(":8000", nil)
}