// 事前に以下のコマンドを実行しライブラリをダウンロードする必要がある
// $ go get golang.org/x/oauth2
// $ go get github.com/skratchdot/open-golang/open
// $ go get github.com/joho/godotenv

package main

import (
  "context"
  "encoding/json"
  "fmt"
  "github.com/skratchdot/open-golang/open"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/github"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  // "strings"
  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file")
  }

  var clientID = os.Getenv("CLIENT_ID")
  var clientSecret = os.Getenv("CLIENT_SECRET")
  var state = "your state"

  conf := &oauth2.Config{
    ClientID: clientID,
    ClientSecret: clientSecret,
    Scopes: []string{"user:email", "gist"},
    Endpoint: github.Endpoint,
  }
  var token *oauth2.Token

  file, err := os.Open("access_token.json")
  if os.IsNotExist(err) {
    url := conf.AuthCodeURL(state, oauth2.AccessTypeOnline)

    code := make(chan string)
    var server *http.Server
    server = &http.Server{
      Addr: ":18888",
      Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html")
        io.WriteString(w, "<html><script>window.open('about:blank', '_self').close()</script></html>")
        w.(http.Flusher).Flush()
        code <- r.URL.Query().Get("code")
        server.Shutdown(context.Background())
      }),
    }
    go server.ListenAndServe()

    open.Start(url)

    token, err := conf.Exchange(oauth2.NoContext, <-code)
    if err != nil {
      panic(err)
    }

    file, err := os.Create("access_token.json")
    if err != nil {
      panic(err)
    }
    json.NewEncoder(file).Encode(token)
  } else if err == nil {
    token = &oauth2.Token{}
    json.NewDecoder(file).Decode(token)
  } else {
    panic(err)
  }

  client := oauth2.NewClient(oauth2.NoContext, conf.TokenSource(oauth2.NoContext, token))

  // Emailの取得
  resp, err := client.Get("https://api.github.com/user/emails")
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  emails, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }
  fmt.Println(string(emails))
}
