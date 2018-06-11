package main

import ("fmt"
		"net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey there!</h1>")
	fmt.Fprintf(w, "<p>go is fast </p>")
	fmt.Fprintf(w, "<p>and simple...... </p>")

	fmt.Fprintf(w, "you %s even add %s  ", "can", "<strong>variable</strong>")


	fmt.Fprintf(w, `<h1>Hey there!</h1> 
		<p>go is fast </p>
		<p>and simple...... </p>`)
}


func main() {
	http.HandleFunc("/", index_handler)
	
	http.ListenAndServe(":8000", nil)
}