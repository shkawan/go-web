package main

import (
	"fmt"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// begin web page
	var hostPlatform = os.Getenv("HOST_PLATFORM")
	var backColor = os.Getenv("BACK_COLOR")

	var htmlHeader = "<!DOCTYPE html><html><h2>Simple web app on Azure</h2>"
	fmt.Fprintf(w, htmlHeader)
	fmt.Fprintf(w, "<body style=background-color:%s><p>Running on: %s</p></body></html>", backColor, hostPlatform)

}

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
  hostname, err := os.Hostname()

  if err != nil {
    panic(err)
  }

	var htmlHeader = "<!DOCTYPE html><html><h2>My Hostname</h2>"
	fmt.Fprintf(w, htmlHeader)
  fmt.Fprintf(w, "<h2>My name is %s</h2>", hostname)
}

func main() {
	http.HandleFunc("/", indexHandler)
  http.HandleFunc("/hostname", hostnameHandler)
	http.ListenAndServe(":8080", nil)
}

