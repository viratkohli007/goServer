package main 
import (
         "net/http"
         "fmt"
         "strings"
         "log"
         )

func main() {
  
    http.HandleFunc("/", sayHello)
    http.HandleFunc("/a", sayHii)

    err := http.ListenAndServe(":8080", nil)
    if err != nil{
      log.Fatal(err)
    }

}

func sayHello(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello, %q", strings.Trim(r.URL.Path, "/"))
}

func sayHii(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, `<h1>Aniket</h1>`)
}