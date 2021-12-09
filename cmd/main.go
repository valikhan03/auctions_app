package main

import (
	"auctionservice/server"
	"log"
)

func main() {
	app := server.NewApp()
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
