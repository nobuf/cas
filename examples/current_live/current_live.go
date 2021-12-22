package main

import (
	"github.com/nobuf/cas"
	"log"
	"os"
)

func main() {
	api := cas.New(os.Getenv("TWITCASTING_API_CLIENT_ID"), os.Getenv("TWITCASTING_API_CLIENT_SECRET"))
	m, err := api.UserCurrentLive(os.Args[1])
	if err != nil {
		if requestError, ok := err.(*cas.RequestError); ok && requestError.Content.Code == 404 {
			log.Printf("%+v\n", requestError)
		} else {
			log.Fatal(err)
		}
	}
	log.Printf("%+v\n", m)
}
