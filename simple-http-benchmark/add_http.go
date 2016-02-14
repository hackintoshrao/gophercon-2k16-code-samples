package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Add struct {
	Sum int
}

func handleMapAdd(w http.ResponseWriter, r *http.Request) {
	var html bytes.Buffer
	first, second := r.FormValue("first"), r.FormValue("second")
	one, _ := strconv.Atoi(first)
	two, _ := strconv.Atoi(second)

	m := map[int]int{0: one, 1: two}
	mapSum := Add{Sum: m[0] + m[1]}
	//parse the html file
	t, _ := template.ParseFiles("template.html")
	t.Execute(&html, mapSum)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html.String()))
}

func handleStructAdd(w http.ResponseWriter, r *http.Request) {

	var html bytes.Buffer
	first, second := r.FormValue("first"), r.FormValue("second")
	one, _ := strconv.Atoi(first)
	two, _ := strconv.Atoi(second)

	m := struct{ a, b int }{one, two}
	structSum := Add{Sum: m.a + m.b}

	t, _ := template.ParseFiles("template.html")
	t.Execute(&html, structSum)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h2> Sum of Struct: " + fmt.Sprint(m.a+m.b) + "</h2>"))
}

func main() {
	http.HandleFunc("/map", handleMapAdd)
	http.HandleFunc("/struct", handleStructAdd)
	log.Fatal(http.ListenAndServe("127.0.0.1:8081", nil))
}
