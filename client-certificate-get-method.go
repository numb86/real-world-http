package main

import (
  "crypto/tls"
  "fmt"
  "net/http"
  "net/http/httputil"
)

func main() {
  cert, err := tls.LoadX509KeyPair("tls/client.crt", "tls/client.key")
  if err != nil {
    panic(err)
  }
  client := &http.Client{
    Transport: &http.Transport{
      TLSClientConfig: &tls.Config{
        Certificates: []tls.Certificate{cert},
      },
    },
  }
  request, err := http.NewRequest("GET", "https://localhost:18443", nil)
  if err != nil {
    panic(err)
  }
  resp, err := client.Do(request)
  if err != nil {
    panic(err)
  }
  dump, err := httputil.DumpResponse(resp, true)
  if err != nil {
    panic(err)
  }
  fmt.Println(string(dump))
}
