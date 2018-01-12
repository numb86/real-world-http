package main

import (
  "fmt"
  "log"
  "io/ioutil"
  "net/http"
)

var image []byte

func init() {
  var err error
  image, err = ioutil.ReadFile("./image.png")
  if err != nil {
    panic(err)
  }
}

func handlerHtml(w http.ResponseWriter, r *http.Request) {
  pusher, ok := w.(http.Pusher)
  // サーバープッシュが可能な環境でのみ、プッシュする
  // $ GODEBUG=http2server=0 で実行すると、HTTP/1.1になるためプッシュしない
  if ok {
    pusher.Push("/image", nil)
  }
  w.Header().Add("Content-Type", "text/html")
  fmt.Fprintf(w, `<html><body><img src="/image"></body></html>`)
}

func handlerImage(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "image/png")
  w.Write(image)
}

func main() {
  http.HandleFunc("/", handlerHtml)
  http.HandleFunc("/image", handlerImage)
  log.Println("start http listening :18443")
  err := http.ListenAndServeTLS(":18443", "tls/server.crt", "tls/server.key", nil)
  log.Println(err)
}
