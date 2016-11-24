package main

import "fmt"
import "net/http"

func main() {
  fmt.Println("Hello world")

  http.HandleFunc("/latency", defaultLatency)

  http.ListenAndServe(":6001", nil)
}

func defaultLatency(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Sleeping for 500ms")
}
