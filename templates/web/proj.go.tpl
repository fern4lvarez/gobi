package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
      http.ServeFile(w, r, r.URL.Path[1:])
  })
	fmt.Println("Listening on", os.Getenv("PORT"), "...")
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.ExecuteTemplate(res, "index.html", nil)
}
