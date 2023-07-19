package main

import (
	"fmt"
	"log"
	"net/http"
)

var pathsToUrls = map[string]string{
	"/new":  "https://godoc.org/github.com/gophercises/urlshort",
	"/test": "https://godoc.org/gopkg.in/yaml.v2",
}

func handler(w http.ResponseWriter, r *http.Request) {
	value, exists := pathsToUrls[r.URL.Path]

	if exists {
		http.Redirect(w, r, value, http.StatusTemporaryRedirect)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "test/plain")
		_, err := w.Write([]byte("Success!!!!"))
		if err != nil {
			return
		}
	}
	fmt.Printf("path is :%v \n", r.URL.Path)
	return
}

func MapHandler(pathsMap map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		if value, exists := pathsMap[request.URL.Path]; exists {
			http.Redirect(writer, request, value, http.StatusTemporaryRedirect)
		} else {
			fallback.ServeHTTP(writer, request)
		}
	}
}

func main() {

	mux := http.DefaultServeMux

	port := ":8080" // Replace with the desired port number

	// Print the port number
	fmt.Println("Opening port:", port)

	mapHandler := MapHandler(pathsToUrls)

	// Start the server
	err := http.ListenAndServe(port, mapHandler)
	if err != nil {
		log.Fatal(err)
	}
}

//func run() {
//
//}
