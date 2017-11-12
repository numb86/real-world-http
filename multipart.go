package main

import (
  "bytes"
  "io"
  "fmt"
  "mime/multipart"
  "net/http"
  "os"
)

func main() {
  var buffer bytes.Buffer
  writer := multipart.NewWriter(&buffer)
  writer.WriteField("foo boo", "hoge fuga")
  fileWriter, err := writer.CreateFormFile("my-file", "text.txt")
  if err != nil {
    panic(err)
  }
  readFile, err := os.Open("text.txt")
  if err != nil {
    panic(err)
  }
  defer readFile.Close()
  io.Copy(fileWriter, readFile)
  writer.Close()

  resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
  if err != nil {
    panic(err)
  }
  fmt.Println("Status:", resp.Status)
}
