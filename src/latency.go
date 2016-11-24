package main

import "fmt"
import "net/http"
import "time"

func main() {
  fmt.Println("Hello world")

  http.HandleFunc("/latency", defaultLatency)

  http.ListenAndServe(":6001", nil)
}

func defaultLatency(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Sleeping for 500ms\n")
  time.Sleep(time.Duration(500) * time.Millisecond)
  fmt.Fprintf(w, "Sleep now complete!\n")
}
