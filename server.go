package main 
import (
           "log"
           "net/http"
        )

func main() {
	
	mux := http.NewServeMux()

	rh := http.RedirectHandler("http://www.google.com", 303)

	mux.Handle("/foo", rh)
	log.Println("listining...")
	http.ListenAndServe(":3000", mux)
}