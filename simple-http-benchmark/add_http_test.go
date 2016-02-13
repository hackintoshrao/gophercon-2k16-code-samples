package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkHandleMapAdd(b *testing.B) {
	r := request(b, "/?first=20&second=30")
	for i := 0; i < b.N; i++ {
		rw := httptest.NewRecorder()
		handleMapAdd(rw, r)
	}

}

func BenchmarkHandleStructAdd(b *testing.B) {
	r := request(b, "/?first=20&second=30")
	for i := 0; i < b.N; i++ {
		rw := httptest.NewRecorder()
		handleStructAdd(rw, r)
	}

}
func request(t testing.TB, url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}
