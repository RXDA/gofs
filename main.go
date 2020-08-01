package main

import (
	"log"
	"net/http"
	"os"
)

func main()  {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Current Dir is:", dir)
	log.Println("start file server at http://127.0.0.1:8080")
	fileserver := http.FileServer(http.Dir(dir))
	err = http.ListenAndServe("127.0.0.1:8080", fileserver)
	if err != nil {
		log.Fatal(err)
	}
}
