package main

import "fmt"
import "net/http"
import "time"
import "log"

func main() {
  fmt.Println("Hello world")

  http.HandleFunc("/latency", latency)

  http.ListenAndServe(":6001", nil)
}

func latency(w http.ResponseWriter, r *http.Request) {
  duration := "500ms"

  durationParam := r.FormValue("duration")

  if durationParam != "" {
    duration = durationParam
  }

  parsedDuration, error := time.ParseDuration(duration)

  if error != nil {
    http.Error(w, fmt.Sprintf("Invalid duration %s", duration), http.StatusBadRequest)
    log.Print(error)
    return
  }

  fmt.Fprintf(w, "Sleeping for %s\n", duration)
  time.Sleep(parsedDuration)
  fmt.Fprintf(w, "Sleep now complete!\n")
}
