package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handleMapAdd(w http.ResponseWriter, r *http.Request) {

	first, second := r.FormValue("first"), r.FormValue("second")
	one, _ := strconv.Atoi(first)
	two, _ := strconv.Atoi(second)
	m := map[int]int{0: one, 1: two}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h2> Sum of Maps: " + fmt.Sprint(m[0]+m[1]) + "</h2>"))
}

func handleStructAdd(w http.ResponseWriter, r *http.Request) {
	first, second := r.FormValue("first"), r.FormValue("second")
	one, _ := strconv.Atoi(first)
	two, _ := strconv.Atoi(second)
	m := struct{ a, b int }{one, two}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h2> Sum of Struct: " + fmt.Sprint(m.a+m.b) + "</h2>"))
}

func main() {
	http.HandleFunc("/map", handleMapAdd)
	http.HandleFunc("/struct", handleStructAdd)
	log.Fatal(http.ListenAndServe("127.0.0.1:8081", nil))
}
