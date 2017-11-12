package main

import (
  "bytes"
  "io"
  "fmt"
  "mime/multipart"
  "net/http"
  "net/textproto"
  "os"
)

func main() {
  var buffer bytes.Buffer
  writer := multipart.NewWriter(&buffer)
  writer.WriteField("foo boo", "hoge fuga")

  part := make(textproto.MIMEHeader)
  part.Set("Content-type", "text/plain")
  part.Set("Content-Disposition", `form-data; name="my-file"; filename="text.txt"`)
  fileWriter, err := writer.CreatePart(part)
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
