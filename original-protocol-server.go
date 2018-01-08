package main

import (
  "io"
  "time"
  "fmt"
  "log"
  "net/http"
)

func handlerUpgrade(w http.ResponseWriter, r *http.Request) {
  if r.Header.Get("Connection") != "Upgrade" || r.Header.Get("Upgrade") != "OriginalProtocol" {
    w.WriteHeader(400)
    return
  }

  fmt.Println("Upgrade to OriginalProtocol")

  // http.ResponseWriter をハイジャックして低層のソケットを直接操作できるようにする
  hijacker := w.(http.Hijacker)
  conn, readWriter, err := hijacker.Hijack()
  if err != nil {
    panic(err)
    return
  }
  defer conn.Close()

  // プロトコルが変わるというレスポンスを送信
  response := http.Response{
    StatusCode: 101,
    Header: make(http.Header),
  }
  response.Header.Set("Upgrade", "OriginalProtocol")
  response.Header.Set("Connection", "Upgrade")
  response.Write(conn)

  // オリジナルの通信の開始
  for i := 1; i <= 10; i++ {
    fmt.Fprintf(readWriter, "%d\n", i)
    fmt.Println("->", i)
    readWriter.Flush()
    recv, err := readWriter.ReadBytes('\n')
    if err == io.EOF {
      break
    }
    fmt.Println("<- %s", string(recv))
    time.Sleep(500 * time.Millisecond)
  }
}

func main() {
  var httpServer http.Server
  http.HandleFunc("/", handlerUpgrade)
  log.Println("start http listening :18888")
  httpServer.Addr = ":18888"
  log.Println(httpServer.ListenAndServe())
}
