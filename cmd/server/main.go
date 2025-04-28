package main

import (
	"log"
	"github.com/Vkanhan/go-server/internal/server"

)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
