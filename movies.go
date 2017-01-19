package cas

import (
	"strconv"
)

type MovieID string

func (id MovieID) ToInt() int {
	i, err := strconv.Atoi(string(id))
	if err != nil {
		return -1 // TODO better error handling?
	}
	return i
}

func (id MovieID) Equal(a MovieID) bool {
	return id.ToInt() == a.ToInt()
}

func (id MovieID) GreaterThan(a MovieID) bool {
	return id.ToInt() > a.ToInt()
}

func (id MovieID) String() string {
	return string(id)
}

type Movie struct {
	ID               MovieID `json:"id"`
	UserID           string  `json:"user_id"`
	Title            string  `json:"title"`
	Subtitle         string  `json:"subtitle"`
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

type MovieContainer struct {
	Movie       Movie    `json:"movie"`
	Broadcaster User     `json:"broadcaster"`
	Tags        []string `json:"tags"`
}

type Movies struct {
	Movies []MovieContainer `json:"movies"`
}

func (api *Client) Movie(id MovieID) (*MovieContainer, error) {
	m := &MovieContainer{}
	err := get(api, "/movies/"+id.String(), m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
