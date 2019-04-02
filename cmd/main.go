package main

import (
	"github.com/DavidNix/boggle/server"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	log.Println("starting server on port", port, "...")
	err := server.ListenAndServe(port)
	if err != nil {
		log.Fatalln("server failed", err)
	}
}
