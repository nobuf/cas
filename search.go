package cas

import (
	"github.com/google/go-querystring/query"
	"strings"
)

// SearchLivesOptions are the parameters for the Search Live Movies API.
type SearchLivesOptions struct {
	Limit   int    `url:"limit"`
	Type    string `url:"type"`
	Context string `url:"context"`
	Lang    string `url:"lang"`
}

// TODO NewLives

// RecommendedLives retrieves recommended Moves on live.
// See http://apiv2-doc.twitcasting.tv/#search-live-movies
func (api *Client) RecommendedLives() (*Movies, error) {
	p := &SearchLivesOptions{
		Limit: 50,
		Type:  "recommend",
		Lang:  "ja",
	}
	return api.SearchLives(p)
}

// SearchLivesByTags retrieves Movies by tag(s).
func (api *Client) SearchLivesByTags(tags []string) (*Movies, error) {
	p := &SearchLivesOptions{
		Limit:   50,
		Type:    "tag",
		Context: strings.Join(tags, " "),
		Lang:    "ja",
	}
	return api.SearchLives(p)
}

// TODO SearchLivesByWords

// TODO SearchLivesByCategory

// SearchLives retrieves Movies based on the given parameters.
func (api *Client) SearchLives(p *SearchLivesOptions) (*Movies, error) {
	movies := &Movies{}
	v, _ := query.Values(p)
	err := get(api, "/search/lives?"+v.Encode(), &movies)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
