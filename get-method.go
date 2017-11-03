package main

import (
  "io/ioutil"
  "fmt"
  "net/http"
)

func main() {
  resp, err := http.Get("http://localhost:18888")
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }
  fmt.Println("レスポンスボディ->", string(body))
}
