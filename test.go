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
	http.Handle("/", http.HandlerFunc(testpage))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		fmt.Println("Oh noes.")
	}
}

func testpage(w http.ResponseWriter, req *http.Request) {
  req.ParseForm()
  if value, ok := req.Form["callback"]; ok {
    fmt.Fprintf(w, "Callback is: %v\n", value)
  }
	fmt.Fprintf(w, "hello ")
}

