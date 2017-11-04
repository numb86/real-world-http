package main

import (
  "io/ioutil"
  "fmt"
  "net/http"
  "net/url"
  "reflect"
)

func main() {
  values := url.Values{
    "query": {"hello world"},
  }
  resp, err := http.Get("http://localhost:18888" + "?" + values.Encode())
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }
  fmt.Println("レスポンスボディ:", string(body))
  fmt.Println("Status:", resp.Status)
  fmt.Println("StatusCode", resp.StatusCode)
  fmt.Println("Headers", resp.Header)
  fmt.Println("Content_Length", resp.Header.Get("Content-Length"))
  result := resp.Header.Get("Hoge")
  fmt.Println(reflect.TypeOf(result)) // string
  fmt.Println(result == "") // true
}
