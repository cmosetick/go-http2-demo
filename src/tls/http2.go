// `go run http2.go`
// In your browser enter: https://localhost:8080
// using curl v7.33 or newer run: `curl --http2 --tlsv1 --insecure https://localhost:8080`

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	var server http.Server

	http2.VerboseLogs = false  //set true for verbose console output
	server.Addr = ":8080"

	log.Println("Starting server on localhost port 8080...")

	http2.ConfigureServer(&server, nil)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL: %q\n", html.EscapeString(r.URL.Path))
		ShowRequestInfoHandler(w, r)
	})

	log.Fatal(server.ListenAndServeTLS("localhost.cert", "localhost.key"))
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
}

// Reference: http://www.integralist.co.uk/posts/http2.html
