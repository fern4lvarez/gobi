package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var port = "5555"

func main() {
	if os.Getenv("PORT") != "" {
		 port = os.Getenv("PORT")
	}
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
      http.ServeFile(w, r, r.URL.Path[1:])
  })
	fmt.Println("Listening on", port, "...")
	http.ListenAndServe(":"+port, nil)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.ExecuteTemplate(res, "index.html", nil)
}
