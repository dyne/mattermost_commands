package main

import (
  "log"
  "net/http"
)

func main() {
  handler := http.HandlerFunc(MiniServer)
  log.Fatal(http.ListenAndServe(":4444", handler))
}
