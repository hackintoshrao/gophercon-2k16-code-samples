package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type add struct {
	Sum int
}

func handleStructAdd(w http.ResponseWriter, r *http.Request) {

	var html bytes.Buffer
	first, second := r.FormValue("first"), r.FormValue("second")
	one, err := strconv.Atoi(first)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	two, err := strconv.Atoi(second)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	m := struct{ a, b int }{one, two}
	structSum := add{Sum: m.a + m.b}

	t, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	err = t.Execute(&html, structSum)
	
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h2> Sum of Struct: " + fmt.Sprint(m.a+m.b) + "</h2>"))
}

func main() {
	
	http.HandleFunc("/struct", handleStructAdd)
	log.Fatal(http.ListenAndServe("127.0.0.1:8081", nil))
}
