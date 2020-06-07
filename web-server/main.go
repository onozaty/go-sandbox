package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {

	var dir string
	var port int

	flag.StringVar(&dir, "d", "./", "dir")
	flag.IntVar(&port, "p", 3000, "port")
	flag.Parse()

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	log.Printf("Listening...  dir=%s port=%d", dir, port)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
