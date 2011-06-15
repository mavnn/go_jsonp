package main

import (
	"flag"
	"fmt"
	"http"
)

var addr = flag.String("addr",  ":9999", "http service address")

func main() {
	flag.Parse()
	fmt.Println("Hello, world")
	http.Handle("/", handlerWrapper(http.HandlerFunc(testpage)))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		fmt.Println("Oh noes.")
	}
}

func handlerWrapper(baseHandler http.HandlerFunc) http.HandlerFunc {
	f := func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		callback, ok := req.Form["callback"]
		if ok {
			w.Header()["Content-type"] = []string{"text/javascript"}
			fmt.Fprintf(w, "%v(", callback[0])
		}
		else {
			w.Header()["Content-type"] = []string{"application/json"}
		}
		baseHandler(w, req)
		if ok {
			fmt.Fprintf(w, ");")
		}
	}
	return f
}

func testpage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, `{ some : "json" }`)
}

