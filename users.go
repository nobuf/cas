package cas

// User object
// See http://apiv2-doc.twitcasting.tv/#get-user-info
type User struct {
	ID          string  `json:"id"`
	ScreenID    string  `json:"screen_id"`
	Name        string  `json:"name"`
	Image       string  `json:"image"`
	Profile     string  `json:"profile"`
	LastMovieID MovieID `json:"last_movie_id"`
	Live        bool    `json:"is_live"`
	Created     int     `json:"created"`
}

// UserContainer holds a User object
type UserContainer struct {
	User User `json:"user"`
}

// User retrieves a User object by the given user id
func (api *Client) User(id string) (*UserContainer, error) {
	u := &UserContainer{}
	err := get(api, "/users/"+id, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
