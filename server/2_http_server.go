package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	if err != nil {
		log.Printf("can't print to connection: %s", err)
	}
}

func myJSONHandler(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Code    string
		Message string
	}{
		Code:    "OK",
		Message: fmt.Sprintf("Hi there, I love %s!", strings.TrimPrefix(r.URL.Path, "/json/")),
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		log.Printf("can't json marshal %+v: %s", resp, err)
		return
	}

	if _, err := w.Write(respBytes); err != nil {
		log.Printf("can't write to connection: %s", err)
	}
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	const tmpl = `<h1>{{.Title}}</h1>
  <a href="http://golang.org">GO</a>`
	t, err := template.New("tmpl").Parse(tmpl)
	if err != nil {
		log.Printf("can't create template: %s", err)
		return
	}
	ctx := struct {
		Title string
	}{
		Title: fmt.Sprintf("Hi there, I love %s!", strings.TrimPrefix(r.URL.Path, "/tmpl/")),
	}
	if err := t.Execute(w, ctx); err != nil {
		log.Printf("can't execute and print template: %s", err)
	}
}

func main() {
	http.HandleFunc("/json/", myJSONHandler)
	http.HandleFunc("/tmpl/", templateHandler)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
