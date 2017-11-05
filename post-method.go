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
    "foo": {"hoge fuga"},
  }
  resp, err := http.PostForm("http://localhost:18888", values)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }
  fmt.Println(string(body)) // []
  fmt.Println(reflect.TypeOf(body)) // []uint8
  fmt.Println("StatusCode", resp.StatusCode) // 200
  fmt.Println("Headers", resp.Header)
}
