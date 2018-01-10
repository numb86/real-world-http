// $ go run confirm-protocol.go -> Protocol Version: HTTP/2.0
// $ GODEBUG=http2client=0 go run confirm-protocol.go -> Protocol Version: HTTP/1.1
// $ GODEBUG=http2server=0 go run confirm-protocol.go -> Protocol Version: HTTP/2.0

package main

import (
  "fmt"
  "net/http"
)

func main() {
  resp, err := http.Get("https://google.com/")  
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  fmt.Printf("Protocol Version: %s\n", resp.Proto)
}
