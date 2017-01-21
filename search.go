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

// SearchUsersOptions are the parameters for searching Users.
type SearchUsersOptions struct {
	Words string `url:"words"`
	Limit int    `url:"limit"`
	Lang  string `url:"lang"`
}

// UsersContainer holds a list of Users
type UsersContainer struct {
	Users []User `json:"users"`
}

// SearchUsersByWords retrieves Users by the given word(s).
func (api *Client) SearchUsersByWords(words []string) (*UsersContainer, error) {
	p := &SearchUsersOptions{
		Words: strings.Join(words, " "),
		Limit: 50,
		Lang:  "ja",
	}
	return api.SearchUsers(p)
}

// SearchUsers retrieves Users by the given parameters.
func (api *Client) SearchUsers(p *SearchUsersOptions) (*UsersContainer, error) {
	users := &UsersContainer{}
	v, _ := query.Values(p)
	err := get(api, "/search/users?"+v.Encode(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
