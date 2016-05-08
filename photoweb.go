package main

import (
  "log"
  "net/http"
  "github.com/photoweb/handlers"
)


func main() {
  http.HandleFunc("/", handlers.ListHandler)
  http.HandleFunc("/view", handlers.ViewHandler)
  http.HandleFunc("/upload", handlers.UploadHandler)
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err.Error())
  }
}
