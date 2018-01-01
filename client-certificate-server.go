package main

import (
  "crypto/tls"
  "fmt"
  "log"
  "net/http"
  "net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
  dump, err := httputil.DumpRequest(r, true)
  if err != nil {
    http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
    return
  }
  fmt.Println(string(dump))
  fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func main() {
  server := &http.Server{
    TLSConfig: &tls.Config{
      // ClientAuth: tls.NoClientCert, // クライアント証明書を要求しない（これがデフォルトの挙動）
      ClientAuth: tls.RequireAndVerifyClientCert, // クライアント証明書を要求し、検証する
      MinVersion: tls.VersionTLS12,
    },
    Addr: ":18443",
  }
  http.HandleFunc("/", handler)
  log.Println("start http listening :18443")
  err := server.ListenAndServeTLS("tls/server.crt", "tls/server.key")
  log.Println(err)
}
