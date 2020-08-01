package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main()  {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Current Dir is:", dir)
	log.Println("start server at localhost:8080")
	err = http.ListenAndServe(":8080", http.FileServer(http.Dir(dir)))
	if err != nil {
		log.Fatal(err)
	}
}
