package main

import (
  "bufio"
  "bytes"
  "io"
  "log"
  "net/http"
)

func main() {
  resp, err := http.Get("http://localhost:18888")
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  reader := bufio.NewReader(resp.Body)
  for {
    line, err := reader.ReadBytes('\n')
    if err == io.EOF {
      break
    }
    log.Println(string(bytes.TrimSpace(line)))
  }
}
