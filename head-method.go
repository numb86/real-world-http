package main

import (
  "io/ioutil"
  "fmt"
  "net/http"
  "reflect"
)

func main() {
  resp, err := http.Head("http://localhost:18888")
  // resp, err := http.Head("http://numb86.net")
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }
  fmt.Println(body) // []
  fmt.Println(reflect.TypeOf(body)) // []uint8
  fmt.Println("StatusCode", resp.StatusCode) // 200
  fmt.Println("Headers", resp.Header)
}
