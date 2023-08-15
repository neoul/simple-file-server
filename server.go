// https://gist.github.com/paulmach/7271283
package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

var (
	path = flag.String("path", ".", "path to the folder to serve. Defaults to the current folder")
	port = flag.String("port", "8080", "port to serve on. Defaults to 8080")
)

func main() {
	flag.Parse()

	dirname, err := filepath.Abs(*path)
	if err != nil {
		log.Fatalf("Could not get absolute path to directory: %s: %s", dirname, err.Error())
	}

	log.Printf("Serving %s on port %s", dirname, *port)

	err = Serve(dirname, *port)
	if err != nil {
		log.Fatalf("Could not serve directory: %s: %s", dirname, err.Error())
	}
}

func Serve(dirname string, port string) error {
	fs := http.FileServer(http.Dir(dirname))
	http.Handle("/", fs)

	return http.ListenAndServe(":"+port, nil)
}
