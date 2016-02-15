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
    
Simple Http server with  struct response  

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

Here is the benchmark for the http handlers 
    package main

    import (
        "net/http"
        "net/http/httptest"
        "testing"
    )    
    func TestHandleStructAdd(t *testing.T) {
	r := request(t, "/?first=20&second=30")
	rw := httptest.NewRecorder()
       
    	handleStructAdd(rw, r)
    	if rw.Code == 500 {
    		t.Fatal("Internal server Error: " + rw.Body.String())
    	}
    	if rw.Body.String() != "<h2> Sum of Struct: 50</h2>" {
    		t.Fatal("Wrong response")
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

    $go test 

    $go test -bench=.  -cpuprofile=prof.cpu

    $go tool pprof bench.test cpu.out 


    (pprof) top --cum
    0.13s of 2.23s total ( 5.83%)
    Dropped 78 nodes (cum <= 0.01s)
    Showing top 10 nodes out of 154 (cum >= 0.68s)
          flat  flat%   sum%        cum   cum%
             0     0%     0%      1.33s 59.64%  runtime.goexit
             0     0%     0%      1.11s 49.78%  _/home/hackintosh/bench.BenchmarkHandleStructAdd
             0     0%     0%      1.11s 49.78%  testing.(*B).launch
             0     0%     0%      1.11s 49.78%  testing.(*B).runN
         0.01s  0.45%  0.45%      1.10s 49.33%  _/home/hackintosh/bench.handleStructAdd
            0     0%  0.45%      0.89s 39.91%  runtime.mcall
         0.01s  0.45%   0.9%      0.87s 39.01%  runtime.schedule
         0.11s  4.93%  5.83%      0.85s 38.12%  runtime.findrunnable
            0     0%  5.83%      0.84s 37.67%  runtime.goexit0
            0     0%  5.83%      0.68s 30.49%  html/template.(*Template).Execute
    
    (pprof) top20
    1340ms of 2230ms total (60.09%)
    Dropped 78 nodes (cum <= 11.15ms)
    Showing top 20 nodes out of 154 (cum >= 220ms)
      flat  flat%   sum%        cum   cum%
     320ms 14.35% 14.35%      320ms 14.35%  runtime/internal/atomic.Xchg
     120ms  5.38% 19.73%      120ms  5.38%  runtime/internal/atomic.Xadd
     110ms  4.93% 24.66%      850ms 38.12%  runtime.findrunnable
      70ms  3.14% 27.80%      350ms 15.70%  runtime.mallocgc
      70ms  3.14% 30.94%      360ms 16.14%  runtime.mapassign1
      70ms  3.14% 34.08%       70ms  3.14%  runtime.usleep
      50ms  2.24% 36.32%       50ms  2.24%  runtime.acquirep1
      50ms  2.24% 38.57%       50ms  2.24%  runtime.heapBitsSetType
      50ms  2.24% 40.81%       80ms  3.59%  runtime.heapBitsSweepSpan
      50ms  2.24% 43.05%       90ms  4.04%  runtime.scanobject
      50ms  2.24% 45.29%       50ms  2.24%  runtime.stringiter2
      50ms  2.24% 47.53%       60ms  2.69%  syscall.Syscall
      40ms  1.79% 49.33%       40ms  1.79%  runtime.memclr
      40ms  1.79% 51.12%       40ms  1.79%  runtime.memmove
      40ms  1.79% 52.91%      230ms 10.31%  runtime.newarray
      40ms  1.79% 54.71%      100ms  4.48%  text/template.goodName
      30ms  1.35% 56.05%       30ms  1.35%  runtime.(*mspan).sweep.func1
      30ms  1.35% 57.40%       70ms  3.14%  runtime.evacuate
      30ms  1.35% 58.74%       30ms  1.35%  runtime.releasep
      30ms  1.35% 60.09%      220ms  9.87%  runtime.unlock

    (pprof) list handleStructAdd
    Total: 2.23s
    ROUTINE ======================== _/home/hackintosh/bench.handleStructAdd in /home/hackintosh/bench/simple_add.go
          10ms      1.10s (flat, cum) 49.33% of Total
                .          .     26:		http.Error(w, err.Error(), 500)
             .          .     27:	}
             .          .     28:	m := struct{ a, b int }{one, two}
             .          .     29:	structSum := add{Sum: m.a + m.b}
             .          .     30:
             .      350ms     31:	t, err := template.ParseFiles("template.html")
             .          .     32:	if err != nil {
             .          .     33:		http.Error(w, err.Error(), 500)
             .          .     34:	}
          10ms      690ms     35:	err = t.Execute(&html, structSum)
             .          .     36:	
             .          .     37:	if err != nil {
             .          .     38:		http.Error(w, err.Error(), 500)
             .          .     39:	}
             .       40ms     40:	w.Header().Set("Content-Type", "text/html; charset=utf-8")
             .       20ms     41:	w.Write([]byte("<h2> Sum of Struct: " + fmt.Sprint(m.a+m.b) + "</h2>"))
             .          .     42:}
             .          .     43:
             .          .     44:func main() {
             .          .     45:	
             .          .     46:	http.HandleFunc("/struct", handleStructAdd)


