package main

import (
	"github.com/nobuf/cas"
	"log"
	"os"
)

func main() {
	api := cas.New(os.Getenv("TWITCASTING_API_CLIENT_ID"), os.Getenv("TWITCASTING_API_CLIENT_SECRET"))
	lives, err := api.RecommendedLives()
	//lives, err := api.SearchLivesByTags([]string{"顔出し", "弾き語り"})
	if err != nil {
		log.Fatal(err)
	}
	for _, m := range lives.Movies {
		log.Printf("%s (%s) - %s", m.Movie.Title, m.Broadcaster.Name, m.Movie.Link)
	}
}
