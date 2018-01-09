package main

import (
  "fmt"
  "log"
  "net/http"
  "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
  flusher, ok := w.(http.Flusher)
  if !ok {
    panic("expected http.ResponseWriter to be an http.Flusher")
  }
  for i := 1; i <= 10; i++ {
    fmt.Fprintf(w, "Chunk #%d\n", i)

    // これがあることで、500ミリ秒毎に送信するようになる
    // これをコメントアウトすると、(500*5)ミリ秒待ってからまとめて送信する
    flusher.Flush()

    time.Sleep(500 * time.Millisecond)
  }
  flusher.Flush()
}

func main() {
  var httpServer http.Server
  http.HandleFunc("/", handler)
  log.Println("start http listening :18888")
  httpServer.Addr = ":18888"
  log.Println(httpServer.ListenAndServe())
}
