package main

import (
	"github.com/nobuf/cas"
	"log"
	"os"
)

func main() {
	api := cas.New(os.Getenv("TWITCASTING_API_CLIENT_ID"), os.Getenv("TWITCASTING_API_CLIENT_SECRET"))
	m, err := api.Categories()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", m)
}
