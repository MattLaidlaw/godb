package main

import (
	"github.com/MattLaidlaw/godb/pkg/godb"
	"log"
)

func main() {
	storage := godb.NewBasicMap()
	srv, err := godb.NewServer(storage)
	if err != nil {
		log.Fatalln(err)
	}
	srv.Listen()
}
