package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

const AddForm = `<!doctype html><html><head></head><body style="background: #1CABE3;"><img src="https://dl.dropboxusercontent.com/u/804305/Caylent-Logo-White01.png" style="width: 300px;"></img><p style="padding-top: 100px;width: 100%;color: white;font-size: 80px;font-family: varela round, sans-serif;text-align: center;">Q & A</p><p style="text-align: center;color: white;font-size: 30px;position: absolute;bottom: 10px;width: 30%;left: 35%;font-family: varela round, sans-serif;">#DevOps is Now</p></body></html>`


func hello(w http.ResponseWriter, req *http.Request) {
	url := req.FormValue("url")
  if url == "" {
      fmt.Fprint(w, AddForm)
      return
  }
  key := "Placeholder"
  fmt.Fprintf(w, "http://localhost:80", key)
	// log.Printf("%s %s\n", req.Proto, req.URL)
	// fmt.Fprintf(w, AddForm)
	fmt.Fprintln(w, AddForm)
	// fmt.Fprintln(w, "Welcome to Deis!")
	// fmt.Fprintln(w, "See the documentation at http://docs.deis.io/ for more information.")
}
