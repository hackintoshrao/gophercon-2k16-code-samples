# gophercon-2k16-code-samples
Code samples for My talk at Gophercon 2016 
Simple Map Add and Struct Add functions

       package gobench

       func GoMapAdd() {
       	m := map[int]int{0: 0, 1: 1}
       	_ = m[0] + m[1]
       }

       func GoStructAdd() {
       	m := struct{ a, b int }{0, 1}
       	_ = m.a + m.b
       }

benchmark for simple map add and struct add 

       package gobench

       import (
       	"testing"
       )

       func BenchmarkGoMapAdd(b *testing.B) {

       	for i := 0; i < b.N; i++ {
       		GoMapAdd()
       	}

       }
       func BenchmarkGoStructAdd(b *testing.B) {

       	for i := 0; i < b.N; i++ {
	       	GoStructAdd()
	       }

       }

Execute the benchmark 

   $go test -bench=.

    BenchmarkGoMapAdd-4   	 5000000	       286 ns/op
    BenchmarkGoStructAdd-4	2000000000	         0.56 ns/op
    
Simple Http server with map and struct response  

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
        w.Write([]byte(html.String()))
    }

Here is the benchmark for the http handlers 
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
    


