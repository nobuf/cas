package main

import (
	"github.com/nobuf/cas"
	"log"
	"os"
)

func main() {
	api := cas.New(os.Getenv("TWITCASTING_API_CLIENT_ID"), os.Getenv("TWITCASTING_API_CLIENT_SECRET"))
	userMovies, err := api.UserMovies(os.Args[1], &cas.UserMoviesOption{
		Limit: 3,
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, m := range userMovies.Movies {
		log.Printf("%s - %s", m.Title, m.Link)
	}
}
