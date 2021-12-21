package cas

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"strconv"
)

// MovieID is identifier of a Movie object.
// In the TwitCasting world, it's sequential digits.
type MovieID string

// Int converts MovieID to int.
func (id MovieID) Int() int {
	i, err := strconv.Atoi(string(id))
	if err != nil {
		return -1 // TODO better error handling?
	}
	return i
}

// Equal returns true if two IDs are identical.
func (id MovieID) Equal(a MovieID) bool {
	return id.Int() == a.Int()
}

// GreaterThan return true if the source MovieID is greater than the given ID,
// that means it is newer than the given ID.
func (id MovieID) GreaterThan(a MovieID) bool {
	return id.Int() > a.Int()
}

// String returns MovieID in string.
func (id MovieID) String() string {
	return string(id)
}

// Movie object.
// See http://apiv2-doc.twitcasting.tv/#get-movie-info
type Movie struct {
	ID               MovieID `json:"id"`
	UserID           string  `json:"user_id"`
	Title            string  `json:"title"`
	Subtitle         string  `json:"subtitle"`
	LastOwnerComment string  `json:"last_owner_comment"`
	Link             string  `json:"link"`
	Live             bool    `json:"is_live"`
	Recorded         bool    `json:"is_recorded"`
	CommentCount     int     `json:"comment_count"`
	LargeThumbnail   string  `json:"large_thumbnail"`
	SmallThumbnail   string  `json:"small_thumbnail"`
	Country          string  `json:"country"`
	Duration         int     `json:"duration"`
	Created          int     `json:"created"`
	Collabo          bool    `json:"is_collabo"`
	Protected        bool    `json:"is_protected"`
	MaxViewCount     int     `json:"max_view_count"`
	CurrentViewCount int     `json:"current_view_count"`
	TotalViewCount   int     `json:"total_view_count"`
	HLS              string  `json:"hls_url"`
}

// MovieContainer holds Movie and its related data.
type MovieContainer struct {
	Movie       Movie    `json:"movie"`
	Broadcaster User     `json:"broadcaster"`
	Tags        []string `json:"tags"`
}

// Movies is a slice of MovieContainers.
type Movies struct {
	Movies []MovieContainer `json:"movies"`
}

// UserMovies is a slice of Movie objects.
type UserMovies struct {
	Movies []Movie `json:"movies"`
}

// Movie retrieves a Movie object by the given MovieID.
func (api *Client) Movie(id MovieID) (*MovieContainer, error) {
	m := &MovieContainer{}
	err := get(api, "/movies/"+id.String(), m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// CommentsOption are the parameters for getting Comments.
type CommentsOption struct {
	Offset int    `url:"offset"`
	Limit  int    `url:"limit"`
	Since  string `url:"since_id"`
}

// DefaultCommentsOption returns a CommentsOption with default values.
func (api *Client) DefaultCommentsOption() *CommentsOption {
	return &CommentsOption{
		Offset: 0,
		Limit:  10,
		Since:  "",
	}
}

// UserCurrentLive get the current live of the user by user id, if no, it will throw 404
func (api *Client) UserCurrentLive(id string) (*MovieContainer, error) {
	m := &MovieContainer{}
	err := get(api, fmt.Sprintf("/users/%s/current_live", id), m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Comments retrieves Comments by the given MovieID.
func (api *Client) Comments(id MovieID, p *CommentsOption) (*CommentsContainer, error) {
	c := &CommentsContainer{}
	q, _ := query.Values(p)
	err := get(api, "/movies/"+id.String()+"/comments?"+q.Encode(), c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// UserMoviesOption is a collection of parameters for getting UserMovies.
type UserMoviesOption struct {
	Offset int `url:"offset"`
	Limit  int `url:"limit"`
}

// DefaultUserMoviesOptions returns a UserMoviesOption with default values.
func (api *Client) DefaultUserMoviesOption() *UserMoviesOption {
	return &UserMoviesOption{
		Offset: 0,
		Limit:  10,
	}
}

// UserMovies retrieves Movies by the given user id.
func (api *Client) UserMovies(id string, p *UserMoviesOption) (*UserMovies, error) {
	movies := &UserMovies{}
	q, _ := query.Values(p)
	err := get(api, "/users/"+id+"/movies?"+q.Encode(), movies)
	if err != nil {
		return nil, err
	}
	return movies, nil
}
