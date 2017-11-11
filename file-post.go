package main

import (
  "fmt"
  "net/http"
  "os"
)

func main() {
  file, err := os.Open("text.txt")
  if err != nil {
    panic(err)
  }
  resp, err := http.Post("http://localhost:18888", "text/pain", file)
  if err != nil {
    panic(err)
  }
  fmt.Println("Status:", resp.Status)
}
