package main

import (
	"log"
	"os"

	"github.com/DavidNix/boggle/server"
)

func main() {
	port := os.Getenv("PORT")
	log.Println("starting server on port", port, "...")
	err := server.ListenAndServe(port)
	if err != nil {
		log.Fatalln("server failed", err)
	}
}
