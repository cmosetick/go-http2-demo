// Usage:
// create certificate and key according to README.md notes
// set consoleLog to true if you want console output
// go run http1.1.go
// in browser, open https://localhost:8282
package main

import (
	"fmt"
  "html"
  "log"
	"net/http"
)
// set to false to disable log output to console
var consoleLog = true

func main() {

  log.Println("Starting server on localhost port 8282...")

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL: %q\n", html.EscapeString(r.URL.Path))
		ShowRequestInfoHandler(w, r)
	})
	log.Fatal(http.ListenAndServeTLS(":8282", "localhost.cert", "localhost.key", nil))
}

func ShowRequestInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "RequestURI: %q\n", r.RequestURI)
	fmt.Fprintf(w, "URL: %#v\n", r.URL)
	fmt.Fprintf(w, "Body.ContentLength: %d (-1 means unknown)\n", r.ContentLength)
	fmt.Fprintf(w, "Close: %v (relevant for HTTP/1 only)\n", r.Close)
	fmt.Fprintf(w, "TLS: %#v\n", r.TLS)
	fmt.Fprintf(w, "\nHeaders:\n")

  r.Header.Write(w)

  if consoleLog != false {
    log.Print("served a web page:\n", r,"\n\n")  // add two new lines at and of request to separate log lines.
  }
}
