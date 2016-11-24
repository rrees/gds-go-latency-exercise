package main

import "fmt"
import "net/http"
import "time"
import "log"

func main() {
  fmt.Println("Hello world")

  http.HandleFunc("/latency", defaultLatency)

  http.ListenAndServe(":6001", nil)
}

func defaultLatency(w http.ResponseWriter, r *http.Request) {
  duration := "500ms"

  durationParam := r.FormValue("duration")

  if durationParam != "" {
    duration = durationParam
  }

  fmt.Fprintf(w, "Sleeping for %s\n", duration)
  parsedDuration, error := time.ParseDuration(duration)

  if error != nil {
    log.Fatal(error)
  }

  time.Sleep(parsedDuration)
  fmt.Fprintf(w, "Sleep now complete!\n")
}
