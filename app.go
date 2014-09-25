package main

import (
  "log"
  "net/http"
  "os"
)

// Runserver
func main() {
  port := ":" + os.Args[1]
  http.Handle(
    "/entities/",
    http.StripPrefix("/entities/", http.FileServer(http.Dir("entities"))),
  )
  log.Println("Server listening on", port)
  err := http.ListenAndServe(port, nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}